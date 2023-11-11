// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: device/proto/device_v2.proto

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

// DeviceV2ServiceClient is the client API for DeviceV2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceV2ServiceClient interface {
	// 查看裝置連線狀態
	CheckState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateV2Response, error)
}

type deviceV2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceV2ServiceClient(cc grpc.ClientConnInterface) DeviceV2ServiceClient {
	return &deviceV2ServiceClient{cc}
}

func (c *deviceV2ServiceClient) CheckState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateV2Response, error) {
	out := new(GetStateV2Response)
	err := c.cc.Invoke(ctx, "/device.DeviceV2Service/CheckState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceV2ServiceServer is the server API for DeviceV2Service service.
// All implementations must embed UnimplementedDeviceV2ServiceServer
// for forward compatibility
type DeviceV2ServiceServer interface {
	// 查看裝置連線狀態
	CheckState(context.Context, *GetStateRequest) (*GetStateV2Response, error)
	mustEmbedUnimplementedDeviceV2ServiceServer()
}

// UnimplementedDeviceV2ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceV2ServiceServer struct {
}

func (UnimplementedDeviceV2ServiceServer) CheckState(context.Context, *GetStateRequest) (*GetStateV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckState not implemented")
}
func (UnimplementedDeviceV2ServiceServer) mustEmbedUnimplementedDeviceV2ServiceServer() {}

// UnsafeDeviceV2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceV2ServiceServer will
// result in compilation errors.
type UnsafeDeviceV2ServiceServer interface {
	mustEmbedUnimplementedDeviceV2ServiceServer()
}

func RegisterDeviceV2ServiceServer(s grpc.ServiceRegistrar, srv DeviceV2ServiceServer) {
	s.RegisterService(&DeviceV2Service_ServiceDesc, srv)
}

func _DeviceV2Service_CheckState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceV2ServiceServer).CheckState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.DeviceV2Service/CheckState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceV2ServiceServer).CheckState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceV2Service_ServiceDesc is the grpc.ServiceDesc for DeviceV2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceV2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.DeviceV2Service",
	HandlerType: (*DeviceV2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckState",
			Handler:    _DeviceV2Service_CheckState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device/proto/device_v2.proto",
}
