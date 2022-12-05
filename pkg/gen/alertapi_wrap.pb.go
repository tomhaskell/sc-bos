// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	context "context"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapAlertApi	adapts a gen.AlertApiServer	and presents it as a gen.AlertApiClient
func WrapAlertApi(server AlertApiServer) AlertApiClient {
	return &alertApiWrapper{server}
}

type alertApiWrapper struct {
	server AlertApiServer
}

// compile time check that we implement the interface we need
var _ AlertApiClient = (*alertApiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *alertApiWrapper) UnwrapServer() AlertApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *alertApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *alertApiWrapper) ListAlerts(ctx context.Context, req *ListAlertsRequest, _ ...grpc.CallOption) (*ListAlertsResponse, error) {
	return w.server.ListAlerts(ctx, req)
}

func (w *alertApiWrapper) PullAlerts(ctx context.Context, in *PullAlertsRequest, opts ...grpc.CallOption) (AlertApi_PullAlertsClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullAlertsAlertApiServerWrapper{stream.Server()}
	client := &pullAlertsAlertApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullAlerts(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullAlertsAlertApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullAlertsAlertApiClientWrapper) Recv() (*PullAlertsResponse, error) {
	m := new(PullAlertsResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullAlertsAlertApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullAlertsAlertApiServerWrapper) Send(response *PullAlertsResponse) error {
	return s.ServerStream.SendMsg(response)
}

func (w *alertApiWrapper) AcknowledgeAlert(ctx context.Context, req *AcknowledgeAlertRequest, _ ...grpc.CallOption) (*Alert, error) {
	return w.server.AcknowledgeAlert(ctx, req)
}

func (w *alertApiWrapper) UnacknowledgeAlert(ctx context.Context, req *AcknowledgeAlertRequest, _ ...grpc.CallOption) (*Alert, error) {
	return w.server.UnacknowledgeAlert(ctx, req)
}
