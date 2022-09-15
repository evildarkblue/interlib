// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: device/core/proto/device.proto

package service

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

// CoreDeviceServiceClient is the client API for CoreDeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreDeviceServiceClient interface {
	// 更新裝置上傳數據
	UpdateRawdata(ctx context.Context, opts ...grpc.CallOption) (CoreDeviceService_UpdateRawdataClient, error)
	// 取得所有裝置連線狀態
	GetStateMap(ctx context.Context, in *GetStateMapRequest, opts ...grpc.CallOption) (*GetStateMapResponse, error)
	// 取得裝置中所有欄位數值
	GetValueMap(ctx context.Context, in *GetValueMapRequest, opts ...grpc.CallOption) (CoreDeviceService_GetValueMapClient, error)
	// 遠端指令至裝置
	Remote(ctx context.Context, in *RemoteRequest, opts ...grpc.CallOption) (*RemoteResponse, error)
	// 更新裝置狀態
	UpdateDeviceState(ctx context.Context, in *UpdateDeviceStateRequest, opts ...grpc.CallOption) (CoreDeviceService_UpdateDeviceStateClient, error)
}

type coreDeviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreDeviceServiceClient(cc grpc.ClientConnInterface) CoreDeviceServiceClient {
	return &coreDeviceServiceClient{cc}
}

func (c *coreDeviceServiceClient) UpdateRawdata(ctx context.Context, opts ...grpc.CallOption) (CoreDeviceService_UpdateRawdataClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoreDeviceService_ServiceDesc.Streams[0], "/service.CoreDeviceService/UpdateRawdata", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreDeviceServiceUpdateRawdataClient{stream}
	return x, nil
}

type CoreDeviceService_UpdateRawdataClient interface {
	Send(*UpdateRawdataRequest) error
	Recv() (*UpdateRawdataResponse, error)
	grpc.ClientStream
}

type coreDeviceServiceUpdateRawdataClient struct {
	grpc.ClientStream
}

func (x *coreDeviceServiceUpdateRawdataClient) Send(m *UpdateRawdataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coreDeviceServiceUpdateRawdataClient) Recv() (*UpdateRawdataResponse, error) {
	m := new(UpdateRawdataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreDeviceServiceClient) GetStateMap(ctx context.Context, in *GetStateMapRequest, opts ...grpc.CallOption) (*GetStateMapResponse, error) {
	out := new(GetStateMapResponse)
	err := c.cc.Invoke(ctx, "/service.CoreDeviceService/GetStateMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreDeviceServiceClient) GetValueMap(ctx context.Context, in *GetValueMapRequest, opts ...grpc.CallOption) (CoreDeviceService_GetValueMapClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoreDeviceService_ServiceDesc.Streams[1], "/service.CoreDeviceService/GetValueMap", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreDeviceServiceGetValueMapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoreDeviceService_GetValueMapClient interface {
	Recv() (*GetValueMapResponse, error)
	grpc.ClientStream
}

type coreDeviceServiceGetValueMapClient struct {
	grpc.ClientStream
}

func (x *coreDeviceServiceGetValueMapClient) Recv() (*GetValueMapResponse, error) {
	m := new(GetValueMapResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreDeviceServiceClient) Remote(ctx context.Context, in *RemoteRequest, opts ...grpc.CallOption) (*RemoteResponse, error) {
	out := new(RemoteResponse)
	err := c.cc.Invoke(ctx, "/service.CoreDeviceService/Remote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreDeviceServiceClient) UpdateDeviceState(ctx context.Context, in *UpdateDeviceStateRequest, opts ...grpc.CallOption) (CoreDeviceService_UpdateDeviceStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoreDeviceService_ServiceDesc.Streams[2], "/service.CoreDeviceService/UpdateDeviceState", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreDeviceServiceUpdateDeviceStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoreDeviceService_UpdateDeviceStateClient interface {
	Recv() (*UpdateDeviceStateResponse, error)
	grpc.ClientStream
}

type coreDeviceServiceUpdateDeviceStateClient struct {
	grpc.ClientStream
}

func (x *coreDeviceServiceUpdateDeviceStateClient) Recv() (*UpdateDeviceStateResponse, error) {
	m := new(UpdateDeviceStateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoreDeviceServiceServer is the server API for CoreDeviceService service.
// All implementations must embed UnimplementedCoreDeviceServiceServer
// for forward compatibility
type CoreDeviceServiceServer interface {
	// 更新裝置上傳數據
	UpdateRawdata(CoreDeviceService_UpdateRawdataServer) error
	// 取得所有裝置連線狀態
	GetStateMap(context.Context, *GetStateMapRequest) (*GetStateMapResponse, error)
	// 取得裝置中所有欄位數值
	GetValueMap(*GetValueMapRequest, CoreDeviceService_GetValueMapServer) error
	// 遠端指令至裝置
	Remote(context.Context, *RemoteRequest) (*RemoteResponse, error)
	// 更新裝置狀態
	UpdateDeviceState(*UpdateDeviceStateRequest, CoreDeviceService_UpdateDeviceStateServer) error
	mustEmbedUnimplementedCoreDeviceServiceServer()
}

// UnimplementedCoreDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoreDeviceServiceServer struct {
}

func (UnimplementedCoreDeviceServiceServer) UpdateRawdata(CoreDeviceService_UpdateRawdataServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRawdata not implemented")
}
func (UnimplementedCoreDeviceServiceServer) GetStateMap(context.Context, *GetStateMapRequest) (*GetStateMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateMap not implemented")
}
func (UnimplementedCoreDeviceServiceServer) GetValueMap(*GetValueMapRequest, CoreDeviceService_GetValueMapServer) error {
	return status.Errorf(codes.Unimplemented, "method GetValueMap not implemented")
}
func (UnimplementedCoreDeviceServiceServer) Remote(context.Context, *RemoteRequest) (*RemoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remote not implemented")
}
func (UnimplementedCoreDeviceServiceServer) UpdateDeviceState(*UpdateDeviceStateRequest, CoreDeviceService_UpdateDeviceStateServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateDeviceState not implemented")
}
func (UnimplementedCoreDeviceServiceServer) mustEmbedUnimplementedCoreDeviceServiceServer() {}

// UnsafeCoreDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreDeviceServiceServer will
// result in compilation errors.
type UnsafeCoreDeviceServiceServer interface {
	mustEmbedUnimplementedCoreDeviceServiceServer()
}

func RegisterCoreDeviceServiceServer(s grpc.ServiceRegistrar, srv CoreDeviceServiceServer) {
	s.RegisterService(&CoreDeviceService_ServiceDesc, srv)
}

func _CoreDeviceService_UpdateRawdata_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoreDeviceServiceServer).UpdateRawdata(&coreDeviceServiceUpdateRawdataServer{stream})
}

type CoreDeviceService_UpdateRawdataServer interface {
	Send(*UpdateRawdataResponse) error
	Recv() (*UpdateRawdataRequest, error)
	grpc.ServerStream
}

type coreDeviceServiceUpdateRawdataServer struct {
	grpc.ServerStream
}

func (x *coreDeviceServiceUpdateRawdataServer) Send(m *UpdateRawdataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coreDeviceServiceUpdateRawdataServer) Recv() (*UpdateRawdataRequest, error) {
	m := new(UpdateRawdataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CoreDeviceService_GetStateMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreDeviceServiceServer).GetStateMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CoreDeviceService/GetStateMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreDeviceServiceServer).GetStateMap(ctx, req.(*GetStateMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreDeviceService_GetValueMap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetValueMapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreDeviceServiceServer).GetValueMap(m, &coreDeviceServiceGetValueMapServer{stream})
}

type CoreDeviceService_GetValueMapServer interface {
	Send(*GetValueMapResponse) error
	grpc.ServerStream
}

type coreDeviceServiceGetValueMapServer struct {
	grpc.ServerStream
}

func (x *coreDeviceServiceGetValueMapServer) Send(m *GetValueMapResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CoreDeviceService_Remote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreDeviceServiceServer).Remote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CoreDeviceService/Remote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreDeviceServiceServer).Remote(ctx, req.(*RemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreDeviceService_UpdateDeviceState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateDeviceStateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreDeviceServiceServer).UpdateDeviceState(m, &coreDeviceServiceUpdateDeviceStateServer{stream})
}

type CoreDeviceService_UpdateDeviceStateServer interface {
	Send(*UpdateDeviceStateResponse) error
	grpc.ServerStream
}

type coreDeviceServiceUpdateDeviceStateServer struct {
	grpc.ServerStream
}

func (x *coreDeviceServiceUpdateDeviceStateServer) Send(m *UpdateDeviceStateResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CoreDeviceService_ServiceDesc is the grpc.ServiceDesc for CoreDeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreDeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.CoreDeviceService",
	HandlerType: (*CoreDeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStateMap",
			Handler:    _CoreDeviceService_GetStateMap_Handler,
		},
		{
			MethodName: "Remote",
			Handler:    _CoreDeviceService_Remote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateRawdata",
			Handler:       _CoreDeviceService_UpdateRawdata_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetValueMap",
			Handler:       _CoreDeviceService_GetValueMap_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateDeviceState",
			Handler:       _CoreDeviceService_UpdateDeviceState_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "device/core/proto/device.proto",
}
