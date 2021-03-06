// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ShogiServiceClient is the client API for ShogiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShogiServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Move(ctx context.Context, opts ...grpc.CallOption) (ShogiService_MoveClient, error)
}

type shogiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShogiServiceClient(cc grpc.ClientConnInterface) ShogiServiceClient {
	return &shogiServiceClient{cc}
}

func (c *shogiServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/shogi.v1.ShogiService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shogiServiceClient) Move(ctx context.Context, opts ...grpc.CallOption) (ShogiService_MoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShogiService_serviceDesc.Streams[0], "/shogi.v1.ShogiService/Move", opts...)
	if err != nil {
		return nil, err
	}
	x := &shogiServiceMoveClient{stream}
	return x, nil
}

type ShogiService_MoveClient interface {
	Send(*MoveRequest) error
	Recv() (*MoveResponse, error)
	grpc.ClientStream
}

type shogiServiceMoveClient struct {
	grpc.ClientStream
}

func (x *shogiServiceMoveClient) Send(m *MoveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shogiServiceMoveClient) Recv() (*MoveResponse, error) {
	m := new(MoveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShogiServiceServer is the server API for ShogiService service.
// All implementations must embed UnimplementedShogiServiceServer
// for forward compatibility
type ShogiServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Move(ShogiService_MoveServer) error
	mustEmbedUnimplementedShogiServiceServer()
}

// UnimplementedShogiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShogiServiceServer struct {
}

func (UnimplementedShogiServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedShogiServiceServer) Move(ShogiService_MoveServer) error {
	return status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedShogiServiceServer) mustEmbedUnimplementedShogiServiceServer() {}

// UnsafeShogiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShogiServiceServer will
// result in compilation errors.
type UnsafeShogiServiceServer interface {
	mustEmbedUnimplementedShogiServiceServer()
}

func RegisterShogiServiceServer(s grpc.ServiceRegistrar, srv ShogiServiceServer) {
	s.RegisterService(&_ShogiService_serviceDesc, srv)
}

func _ShogiService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShogiServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shogi.v1.ShogiService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShogiServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShogiService_Move_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShogiServiceServer).Move(&shogiServiceMoveServer{stream})
}

type ShogiService_MoveServer interface {
	Send(*MoveResponse) error
	Recv() (*MoveRequest, error)
	grpc.ServerStream
}

type shogiServiceMoveServer struct {
	grpc.ServerStream
}

func (x *shogiServiceMoveServer) Send(m *MoveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shogiServiceMoveServer) Recv() (*MoveRequest, error) {
	m := new(MoveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ShogiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shogi.v1.ShogiService",
	HandlerType: (*ShogiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ShogiService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Move",
			Handler:       _ShogiService_Move_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "shogi/v1/shogi.proto",
}
