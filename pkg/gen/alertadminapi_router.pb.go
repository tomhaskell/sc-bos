// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// AlertAdminApiRouter is a gen.AlertAdminApiServer that allows routing named requests to specific gen.AlertAdminApiClient
type AlertAdminApiRouter struct {
	UnimplementedAlertAdminApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ AlertAdminApiServer = (*AlertAdminApiRouter)(nil)

func NewAlertAdminApiRouter(opts ...router.Option) *AlertAdminApiRouter {
	return &AlertAdminApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithAlertAdminApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithAlertAdminApiClientFactory(f func(name string) (AlertAdminApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *AlertAdminApiRouter) Register(server *grpc.Server) {
	RegisterAlertAdminApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type gen.AlertAdminApiClient.
func (r *AlertAdminApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a gen.AlertAdminApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *AlertAdminApiRouter) HoldsType(client any) bool {
	_, ok := client.(AlertAdminApiClient)
	return ok
}

func (r *AlertAdminApiRouter) AddAlertAdminApiClient(name string, client AlertAdminApiClient) AlertAdminApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(AlertAdminApiClient)
}

func (r *AlertAdminApiRouter) RemoveAlertAdminApiClient(name string) AlertAdminApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(AlertAdminApiClient)
}

func (r *AlertAdminApiRouter) GetAlertAdminApiClient(name string) (AlertAdminApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AlertAdminApiClient), nil
}

func (r *AlertAdminApiRouter) CreateAlert(ctx context.Context, request *CreateAlertRequest) (*Alert, error) {
	child, err := r.GetAlertAdminApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.CreateAlert(ctx, request)
}

func (r *AlertAdminApiRouter) UpdateAlert(ctx context.Context, request *UpdateAlertRequest) (*Alert, error) {
	child, err := r.GetAlertAdminApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateAlert(ctx, request)
}

func (r *AlertAdminApiRouter) DeleteAlert(ctx context.Context, request *DeleteAlertRequest) (*DeleteAlertResponse, error) {
	child, err := r.GetAlertAdminApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DeleteAlert(ctx, request)
}
