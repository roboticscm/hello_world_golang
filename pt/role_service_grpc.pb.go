// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pt

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	FindHandler(ctx context.Context, in *FindRoleRequest, opts ...grpc.CallOption) (*FindRoleResponse, error)
	UpsertHandler(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	DeleteHandler(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetHandler(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) FindHandler(ctx context.Context, in *FindRoleRequest, opts ...grpc.CallOption) (*FindRoleResponse, error) {
	out := new(FindRoleResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/FindHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpsertHandler(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/role.RoleService/UpsertHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteHandler(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/role.RoleService/DeleteHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) GetHandler(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/role.RoleService/GetHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations should embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	FindHandler(context.Context, *FindRoleRequest) (*FindRoleResponse, error)
	UpsertHandler(context.Context, *Role) (*Role, error)
	DeleteHandler(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error)
	GetHandler(context.Context, *GetRoleRequest) (*Role, error)
}

// UnimplementedRoleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) FindHandler(context.Context, *FindRoleRequest) (*FindRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHandler not implemented")
}
func (UnimplementedRoleServiceServer) UpsertHandler(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertHandler not implemented")
}
func (UnimplementedRoleServiceServer) DeleteHandler(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHandler not implemented")
}
func (UnimplementedRoleServiceServer) GetHandler(context.Context, *GetRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHandler not implemented")
}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_FindHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/FindHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindHandler(ctx, req.(*FindRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpsertHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpsertHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/UpsertHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpsertHandler(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/DeleteHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteHandler(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_GetHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/GetHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetHandler(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindHandler",
			Handler:    _RoleService_FindHandler_Handler,
		},
		{
			MethodName: "UpsertHandler",
			Handler:    _RoleService_UpsertHandler_Handler,
		},
		{
			MethodName: "DeleteHandler",
			Handler:    _RoleService_DeleteHandler_Handler,
		},
		{
			MethodName: "GetHandler",
			Handler:    _RoleService_GetHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/role/role_service.proto",
}