// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: report/proto/service.proto

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

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	// Query Fields Value
	QueryFieldsValue(ctx context.Context, in *QueryFieldsReq, opts ...grpc.CallOption) (*QueryFieldsResp, error)
	QueryFieldsTimeValue(ctx context.Context, in *QueryFieldsReq, opts ...grpc.CallOption) (*QueryFieldsTimeValueResp, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) QueryFieldsValue(ctx context.Context, in *QueryFieldsReq, opts ...grpc.CallOption) (*QueryFieldsResp, error) {
	out := new(QueryFieldsResp)
	err := c.cc.Invoke(ctx, "/report.ReportService/QueryFieldsValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) QueryFieldsTimeValue(ctx context.Context, in *QueryFieldsReq, opts ...grpc.CallOption) (*QueryFieldsTimeValueResp, error) {
	out := new(QueryFieldsTimeValueResp)
	err := c.cc.Invoke(ctx, "/report.ReportService/QueryFieldsTimeValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations must embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	// Query Fields Value
	QueryFieldsValue(context.Context, *QueryFieldsReq) (*QueryFieldsResp, error)
	QueryFieldsTimeValue(context.Context, *QueryFieldsReq) (*QueryFieldsTimeValueResp, error)
	mustEmbedUnimplementedReportServiceServer()
}

// UnimplementedReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) QueryFieldsValue(context.Context, *QueryFieldsReq) (*QueryFieldsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFieldsValue not implemented")
}
func (UnimplementedReportServiceServer) QueryFieldsTimeValue(context.Context, *QueryFieldsReq) (*QueryFieldsTimeValueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFieldsTimeValue not implemented")
}
func (UnimplementedReportServiceServer) mustEmbedUnimplementedReportServiceServer() {}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_QueryFieldsValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFieldsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).QueryFieldsValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.ReportService/QueryFieldsValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).QueryFieldsValue(ctx, req.(*QueryFieldsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_QueryFieldsTimeValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFieldsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).QueryFieldsTimeValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.ReportService/QueryFieldsTimeValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).QueryFieldsTimeValue(ctx, req.(*QueryFieldsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "report.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryFieldsValue",
			Handler:    _ReportService_QueryFieldsValue_Handler,
		},
		{
			MethodName: "QueryFieldsTimeValue",
			Handler:    _ReportService_QueryFieldsTimeValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report/proto/service.proto",
}
