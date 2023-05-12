// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: graph/v1/graph.proto

package v1

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

const (
	LanternService_Illuminate_FullMethodName   = "/graph.v1.LanternService/Illuminate"
	LanternService_GetVertex_FullMethodName    = "/graph.v1.LanternService/GetVertex"
	LanternService_PutVertex_FullMethodName    = "/graph.v1.LanternService/PutVertex"
	LanternService_DeleteVertex_FullMethodName = "/graph.v1.LanternService/DeleteVertex"
	LanternService_GetEdge_FullMethodName      = "/graph.v1.LanternService/GetEdge"
	LanternService_AddEdge_FullMethodName      = "/graph.v1.LanternService/AddEdge"
	LanternService_PutEdge_FullMethodName      = "/graph.v1.LanternService/PutEdge"
	LanternService_DeleteEdge_FullMethodName   = "/graph.v1.LanternService/DeleteEdge"
)

// LanternServiceClient is the client API for LanternService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LanternServiceClient interface {
	Illuminate(ctx context.Context, in *IlluminateRequest, opts ...grpc.CallOption) (*IlluminateResponse, error)
	GetVertex(ctx context.Context, in *GetVertexRequest, opts ...grpc.CallOption) (*GetVertexResponse, error)
	PutVertex(ctx context.Context, in *PutVertexRequest, opts ...grpc.CallOption) (*PutVertexResponse, error)
	DeleteVertex(ctx context.Context, in *DeleteVertexRequest, opts ...grpc.CallOption) (*DeleteVertexResponse, error)
	GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...grpc.CallOption) (*GetEdgeResponse, error)
	AddEdge(ctx context.Context, in *AddEdgeRequest, opts ...grpc.CallOption) (*AddEdgeResponse, error)
	PutEdge(ctx context.Context, in *PutEdgeRequest, opts ...grpc.CallOption) (*PutEdgeResponse, error)
	DeleteEdge(ctx context.Context, in *DeleteEdgeRequest, opts ...grpc.CallOption) (*DeleteEdgeResponse, error)
}

type lanternServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLanternServiceClient(cc grpc.ClientConnInterface) LanternServiceClient {
	return &lanternServiceClient{cc}
}

func (c *lanternServiceClient) Illuminate(ctx context.Context, in *IlluminateRequest, opts ...grpc.CallOption) (*IlluminateResponse, error) {
	out := new(IlluminateResponse)
	err := c.cc.Invoke(ctx, LanternService_Illuminate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) GetVertex(ctx context.Context, in *GetVertexRequest, opts ...grpc.CallOption) (*GetVertexResponse, error) {
	out := new(GetVertexResponse)
	err := c.cc.Invoke(ctx, LanternService_GetVertex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) PutVertex(ctx context.Context, in *PutVertexRequest, opts ...grpc.CallOption) (*PutVertexResponse, error) {
	out := new(PutVertexResponse)
	err := c.cc.Invoke(ctx, LanternService_PutVertex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) DeleteVertex(ctx context.Context, in *DeleteVertexRequest, opts ...grpc.CallOption) (*DeleteVertexResponse, error) {
	out := new(DeleteVertexResponse)
	err := c.cc.Invoke(ctx, LanternService_DeleteVertex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...grpc.CallOption) (*GetEdgeResponse, error) {
	out := new(GetEdgeResponse)
	err := c.cc.Invoke(ctx, LanternService_GetEdge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) AddEdge(ctx context.Context, in *AddEdgeRequest, opts ...grpc.CallOption) (*AddEdgeResponse, error) {
	out := new(AddEdgeResponse)
	err := c.cc.Invoke(ctx, LanternService_AddEdge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) PutEdge(ctx context.Context, in *PutEdgeRequest, opts ...grpc.CallOption) (*PutEdgeResponse, error) {
	out := new(PutEdgeResponse)
	err := c.cc.Invoke(ctx, LanternService_PutEdge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lanternServiceClient) DeleteEdge(ctx context.Context, in *DeleteEdgeRequest, opts ...grpc.CallOption) (*DeleteEdgeResponse, error) {
	out := new(DeleteEdgeResponse)
	err := c.cc.Invoke(ctx, LanternService_DeleteEdge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanternServiceServer is the server API for LanternService service.
// All implementations must embed UnimplementedLanternServiceServer
// for forward compatibility
type LanternServiceServer interface {
	Illuminate(context.Context, *IlluminateRequest) (*IlluminateResponse, error)
	GetVertex(context.Context, *GetVertexRequest) (*GetVertexResponse, error)
	PutVertex(context.Context, *PutVertexRequest) (*PutVertexResponse, error)
	DeleteVertex(context.Context, *DeleteVertexRequest) (*DeleteVertexResponse, error)
	GetEdge(context.Context, *GetEdgeRequest) (*GetEdgeResponse, error)
	AddEdge(context.Context, *AddEdgeRequest) (*AddEdgeResponse, error)
	PutEdge(context.Context, *PutEdgeRequest) (*PutEdgeResponse, error)
	DeleteEdge(context.Context, *DeleteEdgeRequest) (*DeleteEdgeResponse, error)
	mustEmbedUnimplementedLanternServiceServer()
}

// UnimplementedLanternServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLanternServiceServer struct {
}

func (UnimplementedLanternServiceServer) Illuminate(context.Context, *IlluminateRequest) (*IlluminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Illuminate not implemented")
}
func (UnimplementedLanternServiceServer) GetVertex(context.Context, *GetVertexRequest) (*GetVertexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVertex not implemented")
}
func (UnimplementedLanternServiceServer) PutVertex(context.Context, *PutVertexRequest) (*PutVertexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutVertex not implemented")
}
func (UnimplementedLanternServiceServer) DeleteVertex(context.Context, *DeleteVertexRequest) (*DeleteVertexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVertex not implemented")
}
func (UnimplementedLanternServiceServer) GetEdge(context.Context, *GetEdgeRequest) (*GetEdgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEdge not implemented")
}
func (UnimplementedLanternServiceServer) AddEdge(context.Context, *AddEdgeRequest) (*AddEdgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEdge not implemented")
}
func (UnimplementedLanternServiceServer) PutEdge(context.Context, *PutEdgeRequest) (*PutEdgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutEdge not implemented")
}
func (UnimplementedLanternServiceServer) DeleteEdge(context.Context, *DeleteEdgeRequest) (*DeleteEdgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEdge not implemented")
}
func (UnimplementedLanternServiceServer) mustEmbedUnimplementedLanternServiceServer() {}

// UnsafeLanternServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LanternServiceServer will
// result in compilation errors.
type UnsafeLanternServiceServer interface {
	mustEmbedUnimplementedLanternServiceServer()
}

func RegisterLanternServiceServer(s grpc.ServiceRegistrar, srv LanternServiceServer) {
	s.RegisterService(&LanternService_ServiceDesc, srv)
}

func _LanternService_Illuminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IlluminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).Illuminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_Illuminate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).Illuminate(ctx, req.(*IlluminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_GetVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVertexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).GetVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_GetVertex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).GetVertex(ctx, req.(*GetVertexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_PutVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutVertexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).PutVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_PutVertex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).PutVertex(ctx, req.(*PutVertexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_DeleteVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVertexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).DeleteVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_DeleteVertex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).DeleteVertex(ctx, req.(*DeleteVertexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_GetEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).GetEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_GetEdge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).GetEdge(ctx, req.(*GetEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_AddEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).AddEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_AddEdge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).AddEdge(ctx, req.(*AddEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_PutEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).PutEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_PutEdge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).PutEdge(ctx, req.(*PutEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanternService_DeleteEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanternServiceServer).DeleteEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanternService_DeleteEdge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanternServiceServer).DeleteEdge(ctx, req.(*DeleteEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LanternService_ServiceDesc is the grpc.ServiceDesc for LanternService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LanternService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "graph.v1.LanternService",
	HandlerType: (*LanternServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Illuminate",
			Handler:    _LanternService_Illuminate_Handler,
		},
		{
			MethodName: "GetVertex",
			Handler:    _LanternService_GetVertex_Handler,
		},
		{
			MethodName: "PutVertex",
			Handler:    _LanternService_PutVertex_Handler,
		},
		{
			MethodName: "DeleteVertex",
			Handler:    _LanternService_DeleteVertex_Handler,
		},
		{
			MethodName: "GetEdge",
			Handler:    _LanternService_GetEdge_Handler,
		},
		{
			MethodName: "AddEdge",
			Handler:    _LanternService_AddEdge_Handler,
		},
		{
			MethodName: "PutEdge",
			Handler:    _LanternService_PutEdge_Handler,
		},
		{
			MethodName: "DeleteEdge",
			Handler:    _LanternService_DeleteEdge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "graph/v1/graph.proto",
}
