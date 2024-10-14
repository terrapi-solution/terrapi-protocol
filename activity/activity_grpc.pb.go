// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: activity/activity.proto

package activity

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ActivityService_InsertActivity_FullMethodName = "/protocol.ActivityService/InsertActivity"
	ActivityService_ListActivity_FullMethodName   = "/protocol.ActivityService/ListActivity"
)

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityServiceClient interface {
	InsertActivity(ctx context.Context, in *InsertActivityRequest, opts ...grpc.CallOption) (*InsertActivityResponse, error)
	ListActivity(ctx context.Context, in *ListActivityRequest, opts ...grpc.CallOption) (*Activities, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) InsertActivity(ctx context.Context, in *InsertActivityRequest, opts ...grpc.CallOption) (*InsertActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertActivityResponse)
	err := c.cc.Invoke(ctx, ActivityService_InsertActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) ListActivity(ctx context.Context, in *ListActivityRequest, opts ...grpc.CallOption) (*Activities, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Activities)
	err := c.cc.Invoke(ctx, ActivityService_ListActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
// All implementations must embed UnimplementedActivityServiceServer
// for forward compatibility.
type ActivityServiceServer interface {
	InsertActivity(context.Context, *InsertActivityRequest) (*InsertActivityResponse, error)
	ListActivity(context.Context, *ListActivityRequest) (*Activities, error)
	mustEmbedUnimplementedActivityServiceServer()
}

// UnimplementedActivityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityServiceServer struct{}

func (UnimplementedActivityServiceServer) InsertActivity(context.Context, *InsertActivityRequest) (*InsertActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertActivity not implemented")
}
func (UnimplementedActivityServiceServer) ListActivity(context.Context, *ListActivityRequest) (*Activities, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivity not implemented")
}
func (UnimplementedActivityServiceServer) mustEmbedUnimplementedActivityServiceServer() {}
func (UnimplementedActivityServiceServer) testEmbeddedByValue()                         {}

// UnsafeActivityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServiceServer will
// result in compilation errors.
type UnsafeActivityServiceServer interface {
	mustEmbedUnimplementedActivityServiceServer()
}

func RegisterActivityServiceServer(s grpc.ServiceRegistrar, srv ActivityServiceServer) {
	// If the following call pancis, it indicates UnimplementedActivityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActivityService_ServiceDesc, srv)
}

func _ActivityService_InsertActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).InsertActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_InsertActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).InsertActivity(ctx, req.(*InsertActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_ListActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).ListActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_ListActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).ListActivity(ctx, req.(*ListActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityService_ServiceDesc is the grpc.ServiceDesc for ActivityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertActivity",
			Handler:    _ActivityService_InsertActivity_Handler,
		},
		{
			MethodName: "ListActivity",
			Handler:    _ActivityService_ListActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity/activity.proto",
}
