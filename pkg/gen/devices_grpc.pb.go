// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DevicesApiClient is the client API for DevicesApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicesApiClient interface {
	ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	PullDevices(ctx context.Context, in *PullDevicesRequest, opts ...grpc.CallOption) (DevicesApi_PullDevicesClient, error)
}

type devicesApiClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesApiClient(cc grpc.ClientConnInterface) DevicesApiClient {
	return &devicesApiClient{cc}
}

func (c *devicesApiClient) ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	out := new(ListDevicesResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.DevicesApi/ListDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesApiClient) PullDevices(ctx context.Context, in *PullDevicesRequest, opts ...grpc.CallOption) (DevicesApi_PullDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &DevicesApi_ServiceDesc.Streams[0], "/smartcore.bos.DevicesApi/PullDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &devicesApiPullDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DevicesApi_PullDevicesClient interface {
	Recv() (*PullDevicesResponse, error)
	grpc.ClientStream
}

type devicesApiPullDevicesClient struct {
	grpc.ClientStream
}

func (x *devicesApiPullDevicesClient) Recv() (*PullDevicesResponse, error) {
	m := new(PullDevicesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DevicesApiServer is the server API for DevicesApi service.
// All implementations must embed UnimplementedDevicesApiServer
// for forward compatibility
type DevicesApiServer interface {
	ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	PullDevices(*PullDevicesRequest, DevicesApi_PullDevicesServer) error
	mustEmbedUnimplementedDevicesApiServer()
}

// UnimplementedDevicesApiServer must be embedded to have forward compatible implementations.
type UnimplementedDevicesApiServer struct {
}

func (UnimplementedDevicesApiServer) ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (UnimplementedDevicesApiServer) PullDevices(*PullDevicesRequest, DevicesApi_PullDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullDevices not implemented")
}
func (UnimplementedDevicesApiServer) mustEmbedUnimplementedDevicesApiServer() {}

// UnsafeDevicesApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicesApiServer will
// result in compilation errors.
type UnsafeDevicesApiServer interface {
	mustEmbedUnimplementedDevicesApiServer()
}

func RegisterDevicesApiServer(s grpc.ServiceRegistrar, srv DevicesApiServer) {
	s.RegisterService(&DevicesApi_ServiceDesc, srv)
}

func _DevicesApi_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesApiServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.DevicesApi/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesApiServer).ListDevices(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesApi_PullDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DevicesApiServer).PullDevices(m, &devicesApiPullDevicesServer{stream})
}

type DevicesApi_PullDevicesServer interface {
	Send(*PullDevicesResponse) error
	grpc.ServerStream
}

type devicesApiPullDevicesServer struct {
	grpc.ServerStream
}

func (x *devicesApiPullDevicesServer) Send(m *PullDevicesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DevicesApi_ServiceDesc is the grpc.ServiceDesc for DevicesApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevicesApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.DevicesApi",
	HandlerType: (*DevicesApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDevices",
			Handler:    _DevicesApi_ListDevices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullDevices",
			Handler:       _DevicesApi_PullDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "devices.proto",
}
