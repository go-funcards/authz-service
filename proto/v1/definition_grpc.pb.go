// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: v1/definition.proto

package v1

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

// DefinitionClient is the client API for Definition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefinitionClient interface {
	SaveDefs(ctx context.Context, in *SaveDefsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDefs(ctx context.Context, in *DeleteDefsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetDefs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DefsResponse, error)
}

type definitionClient struct {
	cc grpc.ClientConnInterface
}

func NewDefinitionClient(cc grpc.ClientConnInterface) DefinitionClient {
	return &definitionClient{cc}
}

func (c *definitionClient) SaveDefs(ctx context.Context, in *SaveDefsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.v1.Definition/SaveDefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionClient) DeleteDefs(ctx context.Context, in *DeleteDefsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.v1.Definition/DeleteDefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionClient) GetDefs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DefsResponse, error) {
	out := new(DefsResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.Definition/GetDefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefinitionServer is the server API for Definition service.
// All implementations must embed UnimplementedDefinitionServer
// for forward compatibility
type DefinitionServer interface {
	SaveDefs(context.Context, *SaveDefsRequest) (*emptypb.Empty, error)
	DeleteDefs(context.Context, *DeleteDefsRequest) (*emptypb.Empty, error)
	GetDefs(context.Context, *emptypb.Empty) (*DefsResponse, error)
	mustEmbedUnimplementedDefinitionServer()
}

// UnimplementedDefinitionServer must be embedded to have forward compatible implementations.
type UnimplementedDefinitionServer struct {
}

func (UnimplementedDefinitionServer) SaveDefs(context.Context, *SaveDefsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveDefs not implemented")
}
func (UnimplementedDefinitionServer) DeleteDefs(context.Context, *DeleteDefsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDefs not implemented")
}
func (UnimplementedDefinitionServer) GetDefs(context.Context, *emptypb.Empty) (*DefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefs not implemented")
}
func (UnimplementedDefinitionServer) mustEmbedUnimplementedDefinitionServer() {}

// UnsafeDefinitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DefinitionServer will
// result in compilation errors.
type UnsafeDefinitionServer interface {
	mustEmbedUnimplementedDefinitionServer()
}

func RegisterDefinitionServer(s grpc.ServiceRegistrar, srv DefinitionServer) {
	s.RegisterService(&Definition_ServiceDesc, srv)
}

func _Definition_SaveDefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).SaveDefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.Definition/SaveDefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).SaveDefs(ctx, req.(*SaveDefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Definition_DeleteDefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).DeleteDefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.Definition/DeleteDefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).DeleteDefs(ctx, req.(*DeleteDefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Definition_GetDefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefinitionServer).GetDefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.Definition/GetDefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefinitionServer).GetDefs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Definition_ServiceDesc is the grpc.ServiceDesc for Definition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Definition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.Definition",
	HandlerType: (*DefinitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveDefs",
			Handler:    _Definition_SaveDefs_Handler,
		},
		{
			MethodName: "DeleteDefs",
			Handler:    _Definition_DeleteDefs_Handler,
		},
		{
			MethodName: "GetDefs",
			Handler:    _Definition_GetDefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/definition.proto",
}
