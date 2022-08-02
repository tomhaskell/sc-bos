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

// EnrollmentApiClient is the client API for EnrollmentApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnrollmentApiClient interface {
	GetEnrollment(ctx context.Context, in *GetEnrollmentRequest, opts ...grpc.CallOption) (*Enrollment, error)
	CreateEnrollment(ctx context.Context, in *CreateEnrollmentRequest, opts ...grpc.CallOption) (*Enrollment, error)
	DeleteEnrollment(ctx context.Context, in *DeleteEnrollmentRequest, opts ...grpc.CallOption) (*Enrollment, error)
}

type enrollmentApiClient struct {
	cc grpc.ClientConnInterface
}

func NewEnrollmentApiClient(cc grpc.ClientConnInterface) EnrollmentApiClient {
	return &enrollmentApiClient{cc}
}

func (c *enrollmentApiClient) GetEnrollment(ctx context.Context, in *GetEnrollmentRequest, opts ...grpc.CallOption) (*Enrollment, error) {
	out := new(Enrollment)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.EnrollmentApi/GetEnrollment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enrollmentApiClient) CreateEnrollment(ctx context.Context, in *CreateEnrollmentRequest, opts ...grpc.CallOption) (*Enrollment, error) {
	out := new(Enrollment)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.EnrollmentApi/CreateEnrollment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enrollmentApiClient) DeleteEnrollment(ctx context.Context, in *DeleteEnrollmentRequest, opts ...grpc.CallOption) (*Enrollment, error) {
	out := new(Enrollment)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.EnrollmentApi/DeleteEnrollment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnrollmentApiServer is the server API for EnrollmentApi service.
// All implementations must embed UnimplementedEnrollmentApiServer
// for forward compatibility
type EnrollmentApiServer interface {
	GetEnrollment(context.Context, *GetEnrollmentRequest) (*Enrollment, error)
	CreateEnrollment(context.Context, *CreateEnrollmentRequest) (*Enrollment, error)
	DeleteEnrollment(context.Context, *DeleteEnrollmentRequest) (*Enrollment, error)
	mustEmbedUnimplementedEnrollmentApiServer()
}

// UnimplementedEnrollmentApiServer must be embedded to have forward compatible implementations.
type UnimplementedEnrollmentApiServer struct {
}

func (UnimplementedEnrollmentApiServer) GetEnrollment(context.Context, *GetEnrollmentRequest) (*Enrollment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnrollment not implemented")
}
func (UnimplementedEnrollmentApiServer) CreateEnrollment(context.Context, *CreateEnrollmentRequest) (*Enrollment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEnrollment not implemented")
}
func (UnimplementedEnrollmentApiServer) DeleteEnrollment(context.Context, *DeleteEnrollmentRequest) (*Enrollment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEnrollment not implemented")
}
func (UnimplementedEnrollmentApiServer) mustEmbedUnimplementedEnrollmentApiServer() {}

// UnsafeEnrollmentApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnrollmentApiServer will
// result in compilation errors.
type UnsafeEnrollmentApiServer interface {
	mustEmbedUnimplementedEnrollmentApiServer()
}

func RegisterEnrollmentApiServer(s grpc.ServiceRegistrar, srv EnrollmentApiServer) {
	s.RegisterService(&EnrollmentApi_ServiceDesc, srv)
}

func _EnrollmentApi_GetEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrollmentApiServer).GetEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.EnrollmentApi/GetEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrollmentApiServer).GetEnrollment(ctx, req.(*GetEnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnrollmentApi_CreateEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrollmentApiServer).CreateEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.EnrollmentApi/CreateEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrollmentApiServer).CreateEnrollment(ctx, req.(*CreateEnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnrollmentApi_DeleteEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrollmentApiServer).DeleteEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.EnrollmentApi/DeleteEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrollmentApiServer).DeleteEnrollment(ctx, req.(*DeleteEnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EnrollmentApi_ServiceDesc is the grpc.ServiceDesc for EnrollmentApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnrollmentApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vanti.bsp.ew.EnrollmentApi",
	HandlerType: (*EnrollmentApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEnrollment",
			Handler:    _EnrollmentApi_GetEnrollment_Handler,
		},
		{
			MethodName: "CreateEnrollment",
			Handler:    _EnrollmentApi_CreateEnrollment_Handler,
		},
		{
			MethodName: "DeleteEnrollment",
			Handler:    _EnrollmentApi_DeleteEnrollment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "enrollment.proto",
}