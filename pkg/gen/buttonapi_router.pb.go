// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
	io "io"
)

// ButtonApiRouter is a ButtonApiServer that allows routing named requests to specific ButtonApiClient
type ButtonApiRouter struct {
	UnimplementedButtonApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ ButtonApiServer = (*ButtonApiRouter)(nil)

func NewButtonApiRouter(opts ...router.Option) *ButtonApiRouter {
	return &ButtonApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithButtonApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithButtonApiClientFactory(f func(name string) (ButtonApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ButtonApiRouter) Register(server grpc.ServiceRegistrar) {
	RegisterButtonApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type ButtonApiClient.
func (r *ButtonApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a ButtonApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ButtonApiRouter) HoldsType(client any) bool {
	_, ok := client.(ButtonApiClient)
	return ok
}

func (r *ButtonApiRouter) AddButtonApiClient(name string, client ButtonApiClient) ButtonApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(ButtonApiClient)
}

func (r *ButtonApiRouter) RemoveButtonApiClient(name string) ButtonApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(ButtonApiClient)
}

func (r *ButtonApiRouter) GetButtonApiClient(name string) (ButtonApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ButtonApiClient), nil
}

func (r *ButtonApiRouter) GetButtonState(ctx context.Context, request *GetButtonStateRequest) (*ButtonState, error) {
	child, err := r.GetButtonApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetButtonState(ctx, request)
}

func (r *ButtonApiRouter) PullButtonState(request *PullButtonStateRequest, server ButtonApi_PullButtonStateServer) error {
	child, err := r.GetButtonApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullButtonState(reqCtx, request)
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

		var msg *PullButtonStateResponse
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

func (r *ButtonApiRouter) UpdateButtonState(ctx context.Context, request *UpdateButtonStateRequest) (*ButtonState, error) {
	child, err := r.GetButtonApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateButtonState(ctx, request)
}
