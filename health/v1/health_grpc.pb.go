/*
 * Licensed to the TerrAPI under one or more contributor
 * license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The TerrAPI licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: health/v1/health.proto

package health

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
	HealthService_Check_FullMethodName    = "/health.v1.HealthService/Check"
	HealthService_CheckAll_FullMethodName = "/health.v1.HealthService/CheckAll"
)

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	// Check gets the health of the specified service.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// CheckAll gets the health of all services.
	CheckAll(ctx context.Context, in *CheckAllRequest, opts ...grpc.CallOption) (*CheckAllResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, HealthService_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) CheckAll(ctx context.Context, in *CheckAllRequest, opts ...grpc.CallOption) (*CheckAllResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAllResponse)
	err := c.cc.Invoke(ctx, HealthService_CheckAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility.
type HealthServiceServer interface {
	// Check gets the health of the specified service.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// CheckAll gets the health of all services.
	CheckAll(context.Context, *CheckAllRequest) (*CheckAllResponse, error)
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHealthServiceServer struct{}

func (UnimplementedHealthServiceServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthServiceServer) CheckAll(context.Context, *CheckAllRequest) (*CheckAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAll not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}
func (UnimplementedHealthServiceServer) testEmbeddedByValue()                       {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	// If the following call pancis, it indicates UnimplementedHealthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthService_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_CheckAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).CheckAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthService_CheckAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).CheckAll(ctx, req.(*CheckAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health.v1.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HealthService_Check_Handler,
		},
		{
			MethodName: "CheckAll",
			Handler:    _HealthService_CheckAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health/v1/health.proto",
}
