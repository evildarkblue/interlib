// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: report/proto/report_consumer.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenerateReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	IsCustmized bool     `protobuf:"varint,2,opt,name=isCustmized,proto3" json:"isCustmized,omitempty"`
	EndDate     string   `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	StartDate   string   `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	Interval    string   `protobuf:"bytes,5,opt,name=interval,proto3" json:"interval,omitempty"`
	ShowDetail  bool     `protobuf:"varint,6,opt,name=showDetail,proto3" json:"showDetail,omitempty"`
	ShowCompany bool     `protobuf:"varint,7,opt,name=showCompany,proto3" json:"showCompany,omitempty"`
	Sensors     []string `protobuf:"bytes,8,rep,name=sensors,proto3" json:"sensors,omitempty"`
}

func (x *GenerateReportReq) Reset() {
	*x = GenerateReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_report_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateReportReq) ProtoMessage() {}

func (x *GenerateReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_report_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateReportReq.ProtoReflect.Descriptor instead.
func (*GenerateReportReq) Descriptor() ([]byte, []int) {
	return file_report_proto_report_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateReportReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *GenerateReportReq) GetIsCustmized() bool {
	if x != nil {
		return x.IsCustmized
	}
	return false
}

func (x *GenerateReportReq) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *GenerateReportReq) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GenerateReportReq) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *GenerateReportReq) GetShowDetail() bool {
	if x != nil {
		return x.ShowDetail
	}
	return false
}

func (x *GenerateReportReq) GetShowCompany() bool {
	if x != nil {
		return x.ShowCompany
	}
	return false
}

func (x *GenerateReportReq) GetSensors() []string {
	if x != nil {
		return x.Sensors
	}
	return nil
}

var File_report_proto_report_consumer_proto protoreflect.FileDescriptor

var file_report_proto_report_consumer_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73,
	0x43, 0x75, 0x73, 0x74, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x68, 0x6f, 0x77, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x68, 0x6f, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x32, 0x58, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_proto_report_consumer_proto_rawDescOnce sync.Once
	file_report_proto_report_consumer_proto_rawDescData = file_report_proto_report_consumer_proto_rawDesc
)

func file_report_proto_report_consumer_proto_rawDescGZIP() []byte {
	file_report_proto_report_consumer_proto_rawDescOnce.Do(func() {
		file_report_proto_report_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_proto_report_consumer_proto_rawDescData)
	})
	return file_report_proto_report_consumer_proto_rawDescData
}

var file_report_proto_report_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_report_proto_report_consumer_proto_goTypes = []interface{}{
	(*GenerateReportReq)(nil), // 0: report.GenerateReportReq
	(*emptypb.Empty)(nil),     // 1: google.protobuf.Empty
}
var file_report_proto_report_consumer_proto_depIdxs = []int32{
	0, // 0: report.ComsumerService.GenerateReport:input_type -> report.GenerateReportReq
	1, // 1: report.ComsumerService.GenerateReport:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_report_proto_report_consumer_proto_init() }
func file_report_proto_report_consumer_proto_init() {
	if File_report_proto_report_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_proto_report_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateReportReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_report_proto_report_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_proto_report_consumer_proto_goTypes,
		DependencyIndexes: file_report_proto_report_consumer_proto_depIdxs,
		MessageInfos:      file_report_proto_report_consumer_proto_msgTypes,
	}.Build()
	File_report_proto_report_consumer_proto = out.File
	file_report_proto_report_consumer_proto_rawDesc = nil
	file_report_proto_report_consumer_proto_goTypes = nil
	file_report_proto_report_consumer_proto_depIdxs = nil
}
