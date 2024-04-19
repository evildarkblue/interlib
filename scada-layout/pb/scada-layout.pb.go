// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: scada-layout/proto/scada-layout.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *FieldList) Reset() {
	*x = FieldList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scada_layout_proto_scada_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldList) ProtoMessage() {}

func (x *FieldList) ProtoReflect() protoreflect.Message {
	mi := &file_scada_layout_proto_scada_layout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldList.ProtoReflect.Descriptor instead.
func (*FieldList) Descriptor() ([]byte, []int) {
	return file_scada_layout_proto_scada_layout_proto_rawDescGZIP(), []int{0}
}

func (x *FieldList) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type GetFieldsTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldTags []*GetFieldsTagsResponse_FieldTag `protobuf:"bytes,1,rep,name=fieldTags,proto3" json:"fieldTags,omitempty"`
}

func (x *GetFieldsTagsResponse) Reset() {
	*x = GetFieldsTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scada_layout_proto_scada_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFieldsTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFieldsTagsResponse) ProtoMessage() {}

func (x *GetFieldsTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scada_layout_proto_scada_layout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFieldsTagsResponse.ProtoReflect.Descriptor instead.
func (*GetFieldsTagsResponse) Descriptor() ([]byte, []int) {
	return file_scada_layout_proto_scada_layout_proto_rawDescGZIP(), []int{1}
}

func (x *GetFieldsTagsResponse) GetFieldTags() []*GetFieldsTagsResponse_FieldTag {
	if x != nil {
		return x.FieldTags
	}
	return nil
}

type GetFieldsTagsResponse_FieldTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field      string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Equipments []string `protobuf:"bytes,2,rep,name=equipments,proto3" json:"equipments,omitempty"`
	Groups     []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	Floors     []string `protobuf:"bytes,4,rep,name=floors,proto3" json:"floors,omitempty"`
	Unit       string   `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	SensorType string   `protobuf:"bytes,6,opt,name=sensorType,proto3" json:"sensorType,omitempty"`
}

func (x *GetFieldsTagsResponse_FieldTag) Reset() {
	*x = GetFieldsTagsResponse_FieldTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scada_layout_proto_scada_layout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFieldsTagsResponse_FieldTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFieldsTagsResponse_FieldTag) ProtoMessage() {}

func (x *GetFieldsTagsResponse_FieldTag) ProtoReflect() protoreflect.Message {
	mi := &file_scada_layout_proto_scada_layout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFieldsTagsResponse_FieldTag.ProtoReflect.Descriptor instead.
func (*GetFieldsTagsResponse_FieldTag) Descriptor() ([]byte, []int) {
	return file_scada_layout_proto_scada_layout_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetFieldsTagsResponse_FieldTag) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *GetFieldsTagsResponse_FieldTag) GetEquipments() []string {
	if x != nil {
		return x.Equipments
	}
	return nil
}

func (x *GetFieldsTagsResponse_FieldTag) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GetFieldsTagsResponse_FieldTag) GetFloors() []string {
	if x != nil {
		return x.Floors
	}
	return nil
}

func (x *GetFieldsTagsResponse_FieldTag) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GetFieldsTagsResponse_FieldTag) GetSensorType() string {
	if x != nil {
		return x.SensorType
	}
	return ""
}

var File_scada_layout_proto_scada_layout_proto protoreflect.FileDescriptor

var file_scada_layout_proto_scada_layout_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x63, 0x61, 0x64, 0x61, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x61, 0x64, 0x61, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x63, 0x61, 0x64, 0x61, 0x5f, 0x6c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x22, 0x23, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x8a, 0x02, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x63, 0x61, 0x64, 0x61, 0x5f,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x61, 0x67, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x73,
	0x1a, 0xa4, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x6f,
	0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x32, 0x65, 0x0a, 0x12, 0x53, 0x63, 0x61, 0x64, 0x61,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x0d, 0x67, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x61, 0x67, 0x73, 0x12, 0x17,
	0x2e, 0x73, 0x63, 0x61, 0x64, 0x61, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x64, 0x61, 0x5f,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11,
	0x5a, 0x0f, 0x73, 0x63, 0x61, 0x64, 0x61, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scada_layout_proto_scada_layout_proto_rawDescOnce sync.Once
	file_scada_layout_proto_scada_layout_proto_rawDescData = file_scada_layout_proto_scada_layout_proto_rawDesc
)

func file_scada_layout_proto_scada_layout_proto_rawDescGZIP() []byte {
	file_scada_layout_proto_scada_layout_proto_rawDescOnce.Do(func() {
		file_scada_layout_proto_scada_layout_proto_rawDescData = protoimpl.X.CompressGZIP(file_scada_layout_proto_scada_layout_proto_rawDescData)
	})
	return file_scada_layout_proto_scada_layout_proto_rawDescData
}

var file_scada_layout_proto_scada_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scada_layout_proto_scada_layout_proto_goTypes = []interface{}{
	(*FieldList)(nil),                      // 0: scada_layout.FieldList
	(*GetFieldsTagsResponse)(nil),          // 1: scada_layout.GetFieldsTagsResponse
	(*GetFieldsTagsResponse_FieldTag)(nil), // 2: scada_layout.GetFieldsTagsResponse.FieldTag
}
var file_scada_layout_proto_scada_layout_proto_depIdxs = []int32{
	2, // 0: scada_layout.GetFieldsTagsResponse.fieldTags:type_name -> scada_layout.GetFieldsTagsResponse.FieldTag
	0, // 1: scada_layout.ScadaLayoutService.getFieldsTags:input_type -> scada_layout.FieldList
	1, // 2: scada_layout.ScadaLayoutService.getFieldsTags:output_type -> scada_layout.GetFieldsTagsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scada_layout_proto_scada_layout_proto_init() }
func file_scada_layout_proto_scada_layout_proto_init() {
	if File_scada_layout_proto_scada_layout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scada_layout_proto_scada_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldList); i {
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
		file_scada_layout_proto_scada_layout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFieldsTagsResponse); i {
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
		file_scada_layout_proto_scada_layout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFieldsTagsResponse_FieldTag); i {
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
			RawDescriptor: file_scada_layout_proto_scada_layout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scada_layout_proto_scada_layout_proto_goTypes,
		DependencyIndexes: file_scada_layout_proto_scada_layout_proto_depIdxs,
		MessageInfos:      file_scada_layout_proto_scada_layout_proto_msgTypes,
	}.Build()
	File_scada_layout_proto_scada_layout_proto = out.File
	file_scada_layout_proto_scada_layout_proto_rawDesc = nil
	file_scada_layout_proto_scada_layout_proto_goTypes = nil
	file_scada_layout_proto_scada_layout_proto_depIdxs = nil
}
