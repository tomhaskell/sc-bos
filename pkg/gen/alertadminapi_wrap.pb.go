// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// WrapAlertAdminApi	adapts a AlertAdminApiServer	and presents it as a AlertAdminApiClient
func WrapAlertAdminApi(server AlertAdminApiServer) AlertAdminApiClient {
	return &alertAdminApiWrapper{server}
}

type alertAdminApiWrapper struct {
	server AlertAdminApiServer
}

// compile time check that we implement the interface we need
var _ AlertAdminApiClient = (*alertAdminApiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *alertAdminApiWrapper) UnwrapServer() AlertAdminApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *alertAdminApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *alertAdminApiWrapper) CreateAlert(ctx context.Context, req *CreateAlertRequest, _ ...grpc.CallOption) (*Alert, error) {
	return w.server.CreateAlert(ctx, req)
}

func (w *alertAdminApiWrapper) UpdateAlert(ctx context.Context, req *UpdateAlertRequest, _ ...grpc.CallOption) (*Alert, error) {
	return w.server.UpdateAlert(ctx, req)
}

func (w *alertAdminApiWrapper) ResolveAlert(ctx context.Context, req *ResolveAlertRequest, _ ...grpc.CallOption) (*Alert, error) {
	return w.server.ResolveAlert(ctx, req)
}

func (w *alertAdminApiWrapper) DeleteAlert(ctx context.Context, req *DeleteAlertRequest, _ ...grpc.CallOption) (*DeleteAlertResponse, error) {
	return w.server.DeleteAlert(ctx, req)
}
