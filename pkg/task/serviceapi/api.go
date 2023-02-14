// Package serviceapi implements gen.ServiceApi backed by a service.Map.
package serviceapi

import (
	"context"
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/smart-core-os/sc-api/go/types"
	"github.com/smart-core-os/sc-golang/pkg/masks"
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"github.com/vanti-dev/sc-bos/pkg/task/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Api implements a gen.ServicesApiServer backed by service.Map.
type Api struct {
	gen.UnimplementedServicesApiServer
	m   *service.Map
	now func() time.Time

	knownTypes []string
}

func NewApi(m *service.Map, opts ...Option) *Api {
	a := &Api{m: m, now: time.Now}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a *Api) GetService(_ context.Context, request *gen.GetServiceRequest) (*gen.Service, error) {
	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id missing")
	}

	r := a.m.Get(request.Id)
	if r == nil {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	p := recordToProto(r)
	masks.NewResponseFilter(masks.WithFieldMask(request.ReadMask)).Filter(p)
	return p, nil
}

func (a *Api) PullService(request *gen.PullServiceRequest, server gen.ServicesApi_PullServiceServer) error {
	if request.Id == "" {
		return status.Error(codes.InvalidArgument, "id missing")
	}

	r := a.m.Get(request.Id)
	if r == nil {
		return status.Error(codes.NotFound, "id not found")
	}

	ctx, stop := context.WithCancel(server.Context())
	defer stop()

	// we watch the map for changes because we want to stop listening to the service if it's not in the map anymore
	mapChanges := a.m.Listen(ctx)

	var serviceChanges <-chan service.State
	if request.UpdatesOnly {
		serviceChanges = r.Service.StateChanges(ctx)
	} else {
		var state service.State
		state, serviceChanges = r.Service.StateAndChanges(ctx)
		change := stateToPullServiceResponse(request.Name, r.Id, r.Kind, state)
		if err := server.Send(change); err != nil {
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case c := <-mapChanges:
			if c.NewValue == nil && c.OldValue != nil && c.OldValue.Id == request.Id {
				// the service was removed
				return nil
			}
		case c := <-serviceChanges:
			change := stateToPullServiceResponse(request.Name, r.Id, r.Kind, c)
			if err := server.Send(change); err != nil {
				return err
			}
		}
	}
}

func (a *Api) CreateService(_ context.Context, request *gen.CreateServiceRequest) (*gen.Service, error) {
	id, kind, state := protoToState(request.Service)
	id, state, err := a.m.Create(id, kind, state)
	if err != nil {
		return nil, err
	}
	return stateToProto(id, kind, state), nil
}

func (a *Api) DeleteService(ctx context.Context, request *gen.DeleteServiceRequest) (*gen.Service, error) {
	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id missing")
	}

	state, err := a.m.Delete(request.Id)
	if errors.Is(err, service.ErrNotFound) && request.AllowMissing {
		err = nil
	}
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "id not found")
		}
		return nil, err
	}
	return stateToProto(request.Id, "", state), nil
}

func (a *Api) ListServices(_ context.Context, request *gen.ListServicesRequest) (*gen.ListServicesResponse, error) {
	var page PageToken
	if request.PageToken != "" {
		var err error
		page, err = DecodePageToken(request.PageToken)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid page token")
		}
	}

	values := a.m.Values()
	totalSize := len(values)
	sort.Slice(values, func(i, j int) bool {
		return values[i].Id < values[j].Id
	})

	if page.NextId != "" {
		i, _ := sort.Find(len(values), func(i int) int {
			return strings.Compare(page.NextId, values[i].Id)
		})
		values = values[i:]
	}

	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 50
	}
	if pageSize > 1000 {
		pageSize = 1000
	}
	var nextId string
	if len(values) > int(pageSize) {
		if len(values) >= int(pageSize) {
			nextId = values[pageSize].Id
		}
		values = values[:pageSize]
	}

	res := &gen.ListServicesResponse{
		Services:  make([]*gen.Service, len(values)),
		TotalSize: int32(totalSize),
	}
	filter := masks.NewResponseFilter(masks.WithFieldMask(request.ReadMask))
	for i, value := range values {
		res.Services[i] = filter.FilterClone(recordToProto(value)).(*gen.Service)
	}
	if nextId != "" {
		nextPage := PageToken{NextId: nextId}
		nextToken, err := nextPage.Encode()
		if err != nil {
			return nil, err
		}
		res.NextPageToken = nextToken
	}

	return res, nil
}

func (a *Api) PullServices(request *gen.PullServicesRequest, server gen.ServicesApi_PullServicesServer) error {
	ctx, stop := context.WithCancel(server.Context())
	defer stop()

	for c := range a.pullServices(ctx, request) {
		change := &gen.PullServicesResponse{Changes: []*gen.PullServicesResponse_Change{c}}
		if err := server.Send(change); err != nil {
			return err
		}
	}

	return nil
}

func (a *Api) pullServices(ctx context.Context, request *gen.PullServicesRequest) <-chan *gen.PullServicesResponse_Change {
	out := make(chan *gen.PullServicesResponse_Change)

	ctx, stop := context.WithCancel(ctx)

	// publish sends change to out unless ctx is done.
	// Returns true if sending to out won, false if ctx is done.
	publish := func(change *gen.PullServicesResponse_Change) bool {
		select {
		case <-ctx.Done():
			return false
		case out <- change:
			return true
		}
	}
	// watchRecord listens to changes in records service and publishes to out until ctx is done.
	// Calling stop should cancel ctx.
	watchRecord := func(ctx context.Context, stop context.CancelFunc, record *service.Record, updateOnly bool) {
		defer stop() // we shouldn't need this, ctx cancellation is the only way to exit this func anyway

		var last *gen.Service // used for updates as OldValue

		var serviceChanges <-chan service.State
		if updateOnly {
			serviceChanges = record.Service.StateChanges(ctx)
		} else {
			var state service.State
			state, serviceChanges = record.Service.StateAndChanges(ctx)
			last = stateToProto(record.Id, record.Kind, state)
			change := &gen.PullServicesResponse_Change{
				Name:       request.Name,
				ChangeTime: timestamppb.New(a.now()),
				Type:       types.ChangeType_ADD,
				NewValue:   last,
			}
			if !publish(change) {
				return
			}
		}

		for state := range serviceChanges {
			old := last
			last = stateToProto(record.Id, record.Kind, state)
			change := &gen.PullServicesResponse_Change{
				Name:       request.Name,
				ChangeTime: timestamppb.New(a.now()),
				Type:       types.ChangeType_UPDATE,
				OldValue:   old,
				NewValue:   last,
			}
			if !publish(change) {
				return
			}
		}
	}

	listeners := make(map[string]context.CancelFunc)

	changes := a.m.Listen(ctx) // do this before getting the map values

	for _, record := range a.m.Values() {
		ctx, stop := context.WithCancel(ctx)
		listeners[record.Id] = stop
		go watchRecord(ctx, stop, record, request.UpdatesOnly)
	}

	go func() {
		defer stop()
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case c, ok := <-changes:
				if !ok {
					return // changes closed
				}
				if c.OldValue != nil && c.NewValue == nil {
					// remove
					if stop, ok := listeners[c.OldValue.Id]; ok {
						delete(listeners, c.OldValue.Id)
						stop()
						change := &gen.PullServicesResponse_Change{
							Name:       request.Name,
							ChangeTime: timestamppb.New(c.ChangeTime),
							Type:       types.ChangeType_REMOVE,
							OldValue:   recordToProto(c.OldValue),
						}
						if !publish(change) {
							return
						}
					}
				} else if c.OldValue == nil && c.NewValue != nil {
					// add
					ctx, stop := context.WithCancel(ctx)
					listeners[c.NewValue.Id] = stop
					// false here forces watchRecord to publish the ADD event
					go watchRecord(ctx, stop, c.NewValue, false)
				}
			}
		}
	}()

	return out
}

func (a *Api) StartService(_ context.Context, request *gen.StartServiceRequest) (*gen.Service, error) {
	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id missing")
	}

	r := a.m.Get(request.Id)
	if r == nil {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	state, err := r.Service.Start()
	if errors.Is(err, service.ErrAlreadyStarted) {
		if !request.AllowActive {
			return nil, status.Error(codes.FailedPrecondition, "already started")
		}
		err = nil // clear the error
	}
	if err != nil {
		return nil, err
	}
	return stateToProto(r.Id, r.Kind, state), nil
}

func (a *Api) ConfigureService(_ context.Context, request *gen.ConfigureServiceRequest) (*gen.Service, error) {
	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id missing")
	}
	// todo: should we warn if attempting to write config that is empty?

	r := a.m.Get(request.Id)
	if r == nil {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	state, err := r.Service.Configure([]byte(request.ConfigRaw))
	if err != nil {
		return nil, err
	}
	return stateToProto(r.Id, r.Kind, state), nil
}

func (a *Api) StopService(_ context.Context, request *gen.StopServiceRequest) (*gen.Service, error) {
	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id missing")
	}

	r := a.m.Get(request.Id)
	if r == nil {
		return nil, status.Error(codes.NotFound, "id not found")
	}

	state, err := r.Service.Stop()
	if errors.Is(err, service.ErrAlreadyStopped) {
		if !request.AllowInactive {
			return nil, status.Error(codes.FailedPrecondition, "already stopped")
		}
		err = nil // clear the error
	}
	if err != nil {
		return nil, err
	}
	return stateToProto(r.Id, r.Kind, state), nil
}

func (a *Api) GetServiceMetadata(_ context.Context, request *gen.GetServiceMetadataRequest) (*gen.ServiceMetadata, error) {
	md := a.newMetadata()
	a.seedMetadata(md)

	filter := masks.NewResponseFilter(masks.WithFieldMask(request.ReadMask))
	filter.Filter(md)

	return md, nil
}

func (a *Api) PullServiceMetadata(request *gen.PullServiceMetadataRequest, server gen.ServicesApi_PullServiceMetadataServer) error {
	md := a.newMetadata()

	ctx, stop := context.WithCancel(server.Context())
	defer stop()

	changes := a.m.Listen(ctx) // do this before getting the map values
	a.seedMetadata(md)

	filter := masks.NewResponseFilter(masks.WithFieldMask(request.ReadMask))

	var lastSent *gen.ServiceMetadata
	send := func(md *gen.ServiceMetadata) error {
		md = filter.FilterClone(md).(*gen.ServiceMetadata)
		if proto.Equal(md, lastSent) {
			return nil // no change, don't send
		}
		lastSent = md
		change := &gen.PullServiceMetadataResponse_Change{Metadata: md}
		response := &gen.PullServiceMetadataResponse{Changes: []*gen.PullServiceMetadataResponse_Change{change}}
		return server.Send(response)
	}

	if !request.UpdatesOnly {
		if err := send(md); err != nil {
			return err
		}
	}

	for change := range changes {
		if change.OldValue == nil && change.NewValue != nil {
			// add
			md.TotalCount++
			md.TypeCounts[change.NewValue.Kind]++
		} else if change.OldValue != nil && change.NewValue == nil {
			// delete
			md.TotalCount--
			if md.TypeCounts[change.OldValue.Kind] > 0 { // just in case
				md.TypeCounts[change.OldValue.Kind]--
			}
		} else {
			continue // no change worth sending
		}

		if err := send(md); err != nil {
			return err
		}
	}
	return nil
}

func (a *Api) newMetadata() *gen.ServiceMetadata {
	md := &gen.ServiceMetadata{
		TypeCounts: make(map[string]uint32),
	}
	for _, knownType := range a.knownTypes {
		md.TypeCounts[knownType] = 0
	}
	return md
}

func (a *Api) seedMetadata(md *gen.ServiceMetadata) {
	for _, record := range a.m.Values() {
		md.TotalCount++
		md.TypeCounts[record.Kind]++
	}
}