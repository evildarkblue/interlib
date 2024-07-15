// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: scada-layout/proto/scada-layout.proto

package pb

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

// ScadaLayoutServiceClient is the client API for ScadaLayoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScadaLayoutServiceClient interface {
	GetFieldsTags(ctx context.Context, in *FieldList, opts ...grpc.CallOption) (*GetFieldsTagsResponse, error)
	GetReportFields(ctx context.Context, in *FieldList, opts ...grpc.CallOption) (*GetReportFieldsResponse, error)
	GetFieldsWithId(ctx context.Context, in *GetFieldListRequest, opts ...grpc.CallOption) (*GetFieldListResponse, error)
	GetSmartDefrost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSmartDefrostResponse, error)
	GetReport(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetReportResponse, error)
	GetScenario(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetScenarioResponse, error)
	GetAlarmFields(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAlarmFieldsResponse, error)
}

type scadaLayoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScadaLayoutServiceClient(cc grpc.ClientConnInterface) ScadaLayoutServiceClient {
	return &scadaLayoutServiceClient{cc}
}

func (c *scadaLayoutServiceClient) GetFieldsTags(ctx context.Context, in *FieldList, opts ...grpc.CallOption) (*GetFieldsTagsResponse, error) {
	out := new(GetFieldsTagsResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/getFieldsTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scadaLayoutServiceClient) GetReportFields(ctx context.Context, in *FieldList, opts ...grpc.CallOption) (*GetReportFieldsResponse, error) {
	out := new(GetReportFieldsResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/getReportFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scadaLayoutServiceClient) GetFieldsWithId(ctx context.Context, in *GetFieldListRequest, opts ...grpc.CallOption) (*GetFieldListResponse, error) {
	out := new(GetFieldListResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/getFieldsWithId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scadaLayoutServiceClient) GetSmartDefrost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSmartDefrostResponse, error) {
	out := new(GetSmartDefrostResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/getSmartDefrost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scadaLayoutServiceClient) GetReport(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetReportResponse, error) {
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/getReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scadaLayoutServiceClient) GetScenario(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetScenarioResponse, error) {
	out := new(GetScenarioResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/GetScenario", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scadaLayoutServiceClient) GetAlarmFields(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAlarmFieldsResponse, error) {
	out := new(GetAlarmFieldsResponse)
	err := c.cc.Invoke(ctx, "/scada_layout.ScadaLayoutService/GetAlarmFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScadaLayoutServiceServer is the server API for ScadaLayoutService service.
// All implementations must embed UnimplementedScadaLayoutServiceServer
// for forward compatibility
type ScadaLayoutServiceServer interface {
	GetFieldsTags(context.Context, *FieldList) (*GetFieldsTagsResponse, error)
	GetReportFields(context.Context, *FieldList) (*GetReportFieldsResponse, error)
	GetFieldsWithId(context.Context, *GetFieldListRequest) (*GetFieldListResponse, error)
	GetSmartDefrost(context.Context, *emptypb.Empty) (*GetSmartDefrostResponse, error)
	GetReport(context.Context, *emptypb.Empty) (*GetReportResponse, error)
	GetScenario(context.Context, *emptypb.Empty) (*GetScenarioResponse, error)
	GetAlarmFields(context.Context, *emptypb.Empty) (*GetAlarmFieldsResponse, error)
	mustEmbedUnimplementedScadaLayoutServiceServer()
}

// UnimplementedScadaLayoutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScadaLayoutServiceServer struct {
}

func (UnimplementedScadaLayoutServiceServer) GetFieldsTags(context.Context, *FieldList) (*GetFieldsTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFieldsTags not implemented")
}
func (UnimplementedScadaLayoutServiceServer) GetReportFields(context.Context, *FieldList) (*GetReportFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportFields not implemented")
}
func (UnimplementedScadaLayoutServiceServer) GetFieldsWithId(context.Context, *GetFieldListRequest) (*GetFieldListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFieldsWithId not implemented")
}
func (UnimplementedScadaLayoutServiceServer) GetSmartDefrost(context.Context, *emptypb.Empty) (*GetSmartDefrostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmartDefrost not implemented")
}
func (UnimplementedScadaLayoutServiceServer) GetReport(context.Context, *emptypb.Empty) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedScadaLayoutServiceServer) GetScenario(context.Context, *emptypb.Empty) (*GetScenarioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScenario not implemented")
}
func (UnimplementedScadaLayoutServiceServer) GetAlarmFields(context.Context, *emptypb.Empty) (*GetAlarmFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlarmFields not implemented")
}
func (UnimplementedScadaLayoutServiceServer) mustEmbedUnimplementedScadaLayoutServiceServer() {}

// UnsafeScadaLayoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScadaLayoutServiceServer will
// result in compilation errors.
type UnsafeScadaLayoutServiceServer interface {
	mustEmbedUnimplementedScadaLayoutServiceServer()
}

func RegisterScadaLayoutServiceServer(s grpc.ServiceRegistrar, srv ScadaLayoutServiceServer) {
	s.RegisterService(&ScadaLayoutService_ServiceDesc, srv)
}

func _ScadaLayoutService_GetFieldsTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FieldList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetFieldsTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/getFieldsTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetFieldsTags(ctx, req.(*FieldList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScadaLayoutService_GetReportFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FieldList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetReportFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/getReportFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetReportFields(ctx, req.(*FieldList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScadaLayoutService_GetFieldsWithId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFieldListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetFieldsWithId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/getFieldsWithId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetFieldsWithId(ctx, req.(*GetFieldListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScadaLayoutService_GetSmartDefrost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetSmartDefrost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/getSmartDefrost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetSmartDefrost(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScadaLayoutService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/getReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetReport(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScadaLayoutService_GetScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/GetScenario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetScenario(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScadaLayoutService_GetAlarmFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScadaLayoutServiceServer).GetAlarmFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scada_layout.ScadaLayoutService/GetAlarmFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScadaLayoutServiceServer).GetAlarmFields(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ScadaLayoutService_ServiceDesc is the grpc.ServiceDesc for ScadaLayoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScadaLayoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scada_layout.ScadaLayoutService",
	HandlerType: (*ScadaLayoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getFieldsTags",
			Handler:    _ScadaLayoutService_GetFieldsTags_Handler,
		},
		{
			MethodName: "getReportFields",
			Handler:    _ScadaLayoutService_GetReportFields_Handler,
		},
		{
			MethodName: "getFieldsWithId",
			Handler:    _ScadaLayoutService_GetFieldsWithId_Handler,
		},
		{
			MethodName: "getSmartDefrost",
			Handler:    _ScadaLayoutService_GetSmartDefrost_Handler,
		},
		{
			MethodName: "getReport",
			Handler:    _ScadaLayoutService_GetReport_Handler,
		},
		{
			MethodName: "GetScenario",
			Handler:    _ScadaLayoutService_GetScenario_Handler,
		},
		{
			MethodName: "GetAlarmFields",
			Handler:    _ScadaLayoutService_GetAlarmFields_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scada-layout/proto/scada-layout.proto",
}
