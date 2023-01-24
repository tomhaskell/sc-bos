// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"io"

	"github.com/smart-core-os/sc-golang/pkg/router"
	"google.golang.org/grpc"
)

// ServicesApiRouter is a gen.ServicesApiServer that allows routing named requests to specific gen.ServicesApiClient
type ServicesApiRouter struct {
	UnimplementedServicesApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ ServicesApiServer = (*ServicesApiRouter)(nil)

func NewServicesApiRouter(opts ...router.Option) *ServicesApiRouter {
	return &ServicesApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithServicesApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithServicesApiClientFactory(f func(name string) (ServicesApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ServicesApiRouter) Register(server *grpc.Server) {
	RegisterServicesApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type gen.ServicesApiClient.
func (r *ServicesApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a gen.ServicesApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ServicesApiRouter) HoldsType(client any) bool {
	_, ok := client.(ServicesApiClient)
	return ok
}

func (r *ServicesApiRouter) AddServicesApiClient(name string, client ServicesApiClient) ServicesApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(ServicesApiClient)
}

func (r *ServicesApiRouter) RemoveServicesApiClient(name string) ServicesApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(ServicesApiClient)
}

func (r *ServicesApiRouter) GetServicesApiClient(name string) (ServicesApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ServicesApiClient), nil
}

func (r *ServicesApiRouter) GetService(ctx context.Context, request *GetServiceRequest) (*Service, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetService(ctx, request)
}

func (r *ServicesApiRouter) PullService(request *PullServiceRequest, server ServicesApi_PullServiceServer) error {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullService(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *PullServiceResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}

func (r *ServicesApiRouter) CreateService(ctx context.Context, request *CreateServiceRequest) (*Service, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.CreateService(ctx, request)
}

func (r *ServicesApiRouter) DeleteService(ctx context.Context, request *DeleteServiceRequest) (*Service, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DeleteService(ctx, request)
}

func (r *ServicesApiRouter) ListServices(ctx context.Context, request *ListServicesRequest) (*ListServicesResponse, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.ListServices(ctx, request)
}

func (r *ServicesApiRouter) PullServices(request *PullServicesRequest, server ServicesApi_PullServicesServer) error {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullServices(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *PullServicesResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}

func (r *ServicesApiRouter) StartService(ctx context.Context, request *StartServiceRequest) (*Service, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.StartService(ctx, request)
}

func (r *ServicesApiRouter) ConfigureService(ctx context.Context, request *ConfigureServiceRequest) (*Service, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.ConfigureService(ctx, request)
}

func (r *ServicesApiRouter) StopService(ctx context.Context, request *StopServiceRequest) (*Service, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.StopService(ctx, request)
}

func (r *ServicesApiRouter) GetServiceMetadata(ctx context.Context, request *GetServiceMetadataRequest) (*ServiceMetadata, error) {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetServiceMetadata(ctx, request)
}

func (r *ServicesApiRouter) PullServiceMetadata(request *PullServiceMetadataRequest, server ServicesApi_PullServiceMetadataServer) error {
	child, err := r.GetServicesApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullServiceMetadata(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *PullServiceMetadataResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}