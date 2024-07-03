// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: com2scada/proto/com2scada.proto

package pb

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

// Com2ScadaServiceClient is the client API for Com2ScadaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Com2ScadaServiceClient interface {
	ChangeDataStream(ctx context.Context, in *FieldList, opts ...grpc.CallOption) (Com2ScadaService_ChangeDataStreamClient, error)
}

type com2ScadaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCom2ScadaServiceClient(cc grpc.ClientConnInterface) Com2ScadaServiceClient {
	return &com2ScadaServiceClient{cc}
}

func (c *com2ScadaServiceClient) ChangeDataStream(ctx context.Context, in *FieldList, opts ...grpc.CallOption) (Com2ScadaService_ChangeDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Com2ScadaService_ServiceDesc.Streams[0], "/com2scada.Com2ScadaService/changeDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &com2ScadaServiceChangeDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Com2ScadaService_ChangeDataStreamClient interface {
	Recv() (*ChangeDataResponse, error)
	grpc.ClientStream
}

type com2ScadaServiceChangeDataStreamClient struct {
	grpc.ClientStream
}

func (x *com2ScadaServiceChangeDataStreamClient) Recv() (*ChangeDataResponse, error) {
	m := new(ChangeDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Com2ScadaServiceServer is the server API for Com2ScadaService service.
// All implementations must embed UnimplementedCom2ScadaServiceServer
// for forward compatibility
type Com2ScadaServiceServer interface {
	ChangeDataStream(*FieldList, Com2ScadaService_ChangeDataStreamServer) error
	mustEmbedUnimplementedCom2ScadaServiceServer()
}

// UnimplementedCom2ScadaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCom2ScadaServiceServer struct {
}

func (UnimplementedCom2ScadaServiceServer) ChangeDataStream(*FieldList, Com2ScadaService_ChangeDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChangeDataStream not implemented")
}
func (UnimplementedCom2ScadaServiceServer) mustEmbedUnimplementedCom2ScadaServiceServer() {}

// UnsafeCom2ScadaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Com2ScadaServiceServer will
// result in compilation errors.
type UnsafeCom2ScadaServiceServer interface {
	mustEmbedUnimplementedCom2ScadaServiceServer()
}

func RegisterCom2ScadaServiceServer(s grpc.ServiceRegistrar, srv Com2ScadaServiceServer) {
	s.RegisterService(&Com2ScadaService_ServiceDesc, srv)
}

func _Com2ScadaService_ChangeDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FieldList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Com2ScadaServiceServer).ChangeDataStream(m, &com2ScadaServiceChangeDataStreamServer{stream})
}

type Com2ScadaService_ChangeDataStreamServer interface {
	Send(*ChangeDataResponse) error
	grpc.ServerStream
}

type com2ScadaServiceChangeDataStreamServer struct {
	grpc.ServerStream
}

func (x *com2ScadaServiceChangeDataStreamServer) Send(m *ChangeDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Com2ScadaService_ServiceDesc is the grpc.ServiceDesc for Com2ScadaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Com2ScadaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com2scada.Com2ScadaService",
	HandlerType: (*Com2ScadaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "changeDataStream",
			Handler:       _Com2ScadaService_ChangeDataStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com2scada/proto/com2scada.proto",
}
