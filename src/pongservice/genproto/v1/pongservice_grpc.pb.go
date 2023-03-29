// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: pongservice.proto

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
	PongService_Pong_FullMethodName = "/proto.v1.PongService/Pong"
)

// PongServiceClient is the client API for PongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PongServiceClient interface {
	Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type pongServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPongServiceClient(cc grpc.ClientConnInterface) PongServiceClient {
	return &pongServiceClient{cc}
}

func (c *pongServiceClient) Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, PongService_Pong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PongServiceServer is the server API for PongService service.
// All implementations must embed UnimplementedPongServiceServer
// for forward compatibility
type PongServiceServer interface {
	Pong(context.Context, *PongRequest) (*PongResponse, error)
	mustEmbedUnimplementedPongServiceServer()
}

// UnimplementedPongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPongServiceServer struct {
}

func (UnimplementedPongServiceServer) Pong(context.Context, *PongRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedPongServiceServer) mustEmbedUnimplementedPongServiceServer() {}

// UnsafePongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PongServiceServer will
// result in compilation errors.
type UnsafePongServiceServer interface {
	mustEmbedUnimplementedPongServiceServer()
}

func RegisterPongServiceServer(s grpc.ServiceRegistrar, srv PongServiceServer) {
	s.RegisterService(&PongService_ServiceDesc, srv)
}

func _PongService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PongService_Pong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongServiceServer).Pong(ctx, req.(*PongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PongService_ServiceDesc is the grpc.ServiceDesc for PongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.PongService",
	HandlerType: (*PongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _PongService_Pong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pongservice.proto",
}
