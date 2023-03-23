// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// DaliApiRouter is a DaliApiServer that allows routing named requests to specific DaliApiClient
type DaliApiRouter struct {
	UnimplementedDaliApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ DaliApiServer = (*DaliApiRouter)(nil)

func NewDaliApiRouter(opts ...router.Option) *DaliApiRouter {
	return &DaliApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithDaliApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithDaliApiClientFactory(f func(name string) (DaliApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *DaliApiRouter) Register(server *grpc.Server) {
	RegisterDaliApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type DaliApiClient.
func (r *DaliApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a DaliApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *DaliApiRouter) HoldsType(client any) bool {
	_, ok := client.(DaliApiClient)
	return ok
}

func (r *DaliApiRouter) AddDaliApiClient(name string, client DaliApiClient) DaliApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(DaliApiClient)
}

func (r *DaliApiRouter) RemoveDaliApiClient(name string) DaliApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(DaliApiClient)
}

func (r *DaliApiRouter) GetDaliApiClient(name string) (DaliApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(DaliApiClient), nil
}

func (r *DaliApiRouter) AddToGroup(ctx context.Context, request *AddToGroupRequest) (*AddToGroupResponse, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.AddToGroup(ctx, request)
}

func (r *DaliApiRouter) RemoveFromGroup(ctx context.Context, request *RemoveFromGroupRequest) (*RemoveFromGroupResponse, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.RemoveFromGroup(ctx, request)
}

func (r *DaliApiRouter) GetGroupMembership(ctx context.Context, request *GetGroupMembershipRequest) (*GetGroupMembershipResponse, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetGroupMembership(ctx, request)
}

func (r *DaliApiRouter) GetControlGearStatus(ctx context.Context, request *GetControlGearStatusRequest) (*ControlGearStatus, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetControlGearStatus(ctx, request)
}

func (r *DaliApiRouter) GetEmergencyStatus(ctx context.Context, request *GetEmergencyStatusRequest) (*EmergencyStatus, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetEmergencyStatus(ctx, request)
}

func (r *DaliApiRouter) Identify(ctx context.Context, request *IdentifyRequest) (*IdentifyResponse, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.Identify(ctx, request)
}

func (r *DaliApiRouter) StartTest(ctx context.Context, request *StartTestRequest) (*StartTestResponse, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.StartTest(ctx, request)
}

func (r *DaliApiRouter) StopTest(ctx context.Context, request *StopTestRequest) (*StopTestResponse, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.StopTest(ctx, request)
}

func (r *DaliApiRouter) GetTestResult(ctx context.Context, request *GetTestResultRequest) (*TestResult, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetTestResult(ctx, request)
}

func (r *DaliApiRouter) DeleteTestResult(ctx context.Context, request *DeleteTestResultRequest) (*TestResult, error) {
	child, err := r.GetDaliApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DeleteTestResult(ctx, request)
}