// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: channel/proto/channel_report.proto

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

type GetSensorsReportInfoResponse_TypeOfDisplay int32

const (
	GetSensorsReportInfoResponse_REALTIME     GetSensorsReportInfoResponse_TypeOfDisplay = 0
	GetSensorsReportInfoResponse_ACCUMULATION GetSensorsReportInfoResponse_TypeOfDisplay = 1
	GetSensorsReportInfoResponse_STATE        GetSensorsReportInfoResponse_TypeOfDisplay = 2
)

// Enum value maps for GetSensorsReportInfoResponse_TypeOfDisplay.
var (
	GetSensorsReportInfoResponse_TypeOfDisplay_name = map[int32]string{
		0: "REALTIME",
		1: "ACCUMULATION",
		2: "STATE",
	}
	GetSensorsReportInfoResponse_TypeOfDisplay_value = map[string]int32{
		"REALTIME":     0,
		"ACCUMULATION": 1,
		"STATE":        2,
	}
)

func (x GetSensorsReportInfoResponse_TypeOfDisplay) Enum() *GetSensorsReportInfoResponse_TypeOfDisplay {
	p := new(GetSensorsReportInfoResponse_TypeOfDisplay)
	*p = x
	return p
}

func (x GetSensorsReportInfoResponse_TypeOfDisplay) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSensorsReportInfoResponse_TypeOfDisplay) Descriptor() protoreflect.EnumDescriptor {
	return file_channel_proto_channel_report_proto_enumTypes[0].Descriptor()
}

func (GetSensorsReportInfoResponse_TypeOfDisplay) Type() protoreflect.EnumType {
	return &file_channel_proto_channel_report_proto_enumTypes[0]
}

func (x GetSensorsReportInfoResponse_TypeOfDisplay) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSensorsReportInfoResponse_TypeOfDisplay.Descriptor instead.
func (GetSensorsReportInfoResponse_TypeOfDisplay) EnumDescriptor() ([]byte, []int) {
	return file_channel_proto_channel_report_proto_rawDescGZIP(), []int{2, 0}
}

type SensorIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorIds []string `protobuf:"bytes,1,rep,name=sensorIds,proto3" json:"sensorIds,omitempty"`
}

func (x *SensorIdsRequest) Reset() {
	*x = SensorIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorIdsRequest) ProtoMessage() {}

func (x *SensorIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorIdsRequest.ProtoReflect.Descriptor instead.
func (*SensorIdsRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_report_proto_rawDescGZIP(), []int{0}
}

func (x *SensorIdsRequest) GetSensorIds() []string {
	if x != nil {
		return x.SensorIds
	}
	return nil
}

type CountSensorWarningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result map[string]int64 `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *CountSensorWarningResponse) Reset() {
	*x = CountSensorWarningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountSensorWarningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountSensorWarningResponse) ProtoMessage() {}

func (x *CountSensorWarningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountSensorWarningResponse.ProtoReflect.Descriptor instead.
func (*CountSensorWarningResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_report_proto_rawDescGZIP(), []int{1}
}

func (x *CountSensorWarningResponse) GetResult() map[string]int64 {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetSensorsReportInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string                                     `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Sensors []*GetSensorsReportInfoResponse_SensorInfo `protobuf:"bytes,2,rep,name=sensors,proto3" json:"sensors,omitempty"`
}

func (x *GetSensorsReportInfoResponse) Reset() {
	*x = GetSensorsReportInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorsReportInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorsReportInfoResponse) ProtoMessage() {}

func (x *GetSensorsReportInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorsReportInfoResponse.ProtoReflect.Descriptor instead.
func (*GetSensorsReportInfoResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetSensorsReportInfoResponse) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GetSensorsReportInfoResponse) GetSensors() []*GetSensorsReportInfoResponse_SensorInfo {
	if x != nil {
		return x.Sensors
	}
	return nil
}

type GetSensorsReportInfoResponse_Warning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable  bool    `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Upper   float64 `protobuf:"fixed64,2,opt,name=upper,proto3" json:"upper,omitempty"`
	Lower   float64 `protobuf:"fixed64,3,opt,name=lower,proto3" json:"lower,omitempty"`
	Minutes int32   `protobuf:"varint,4,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Times   int32   `protobuf:"varint,5,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *GetSensorsReportInfoResponse_Warning) Reset() {
	*x = GetSensorsReportInfoResponse_Warning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorsReportInfoResponse_Warning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorsReportInfoResponse_Warning) ProtoMessage() {}

func (x *GetSensorsReportInfoResponse_Warning) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorsReportInfoResponse_Warning.ProtoReflect.Descriptor instead.
func (*GetSensorsReportInfoResponse_Warning) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_report_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetSensorsReportInfoResponse_Warning) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *GetSensorsReportInfoResponse_Warning) GetUpper() float64 {
	if x != nil {
		return x.Upper
	}
	return 0
}

func (x *GetSensorsReportInfoResponse_Warning) GetLower() float64 {
	if x != nil {
		return x.Lower
	}
	return 0
}

func (x *GetSensorsReportInfoResponse_Warning) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *GetSensorsReportInfoResponse_Warning) GetTimes() int32 {
	if x != nil {
		return x.Times
	}
	return 0
}

type GetSensorsReportInfoResponse_SensorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Unit        string                                     `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Type        string                                     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Equip       string                                     `protobuf:"bytes,5,opt,name=equip,proto3" json:"equip,omitempty"`
	DisplayType GetSensorsReportInfoResponse_TypeOfDisplay `protobuf:"varint,6,opt,name=displayType,proto3,enum=channel.GetSensorsReportInfoResponse_TypeOfDisplay" json:"displayType,omitempty"`
	Dp          uint32                                     `protobuf:"varint,7,opt,name=dp,proto3" json:"dp,omitempty"`
	TransMap    map[string]string                          `protobuf:"bytes,8,rep,name=transMap,proto3" json:"transMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WarningConf *GetSensorsReportInfoResponse_Warning      `protobuf:"bytes,9,opt,name=warningConf,proto3" json:"warningConf,omitempty"`
}

func (x *GetSensorsReportInfoResponse_SensorInfo) Reset() {
	*x = GetSensorsReportInfoResponse_SensorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorsReportInfoResponse_SensorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorsReportInfoResponse_SensorInfo) ProtoMessage() {}

func (x *GetSensorsReportInfoResponse_SensorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorsReportInfoResponse_SensorInfo.ProtoReflect.Descriptor instead.
func (*GetSensorsReportInfoResponse_SensorInfo) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_report_proto_rawDescGZIP(), []int{2, 1}
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetEquip() string {
	if x != nil {
		return x.Equip
	}
	return ""
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetDisplayType() GetSensorsReportInfoResponse_TypeOfDisplay {
	if x != nil {
		return x.DisplayType
	}
	return GetSensorsReportInfoResponse_REALTIME
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetDp() uint32 {
	if x != nil {
		return x.Dp
	}
	return 0
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetTransMap() map[string]string {
	if x != nil {
		return x.TransMap
	}
	return nil
}

func (x *GetSensorsReportInfoResponse_SensorInfo) GetWarningConf() *GetSensorsReportInfoResponse_Warning {
	if x != nil {
		return x.WarningConf
	}
	return nil
}

var File_channel_proto_channel_report_proto protoreflect.FileDescriptor

var file_channel_proto_channel_report_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x30, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x73, 0x22,
	0xa0, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x57,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x81, 0x06, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4a, 0x0a,
	0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x1a, 0x7d, 0x0a, 0x07, 0x57, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x1a, 0xbf, 0x03, 0x0a, 0x0a, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x71, 0x75, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x71, 0x75, 0x69, 0x70, 0x12, 0x55, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x64, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x64, 0x70,
	0x12, 0x5a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x4d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x4f, 0x0a, 0x0b,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x0b, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x1a, 0x3b, 0x0a,
	0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x0d, 0x54, 0x79,
	0x70, 0x65, 0x4f, 0x66, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x41, 0x4c, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43,
	0x55, 0x4d, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0xc3, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x13, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x19, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_channel_proto_channel_report_proto_rawDescOnce sync.Once
	file_channel_proto_channel_report_proto_rawDescData = file_channel_proto_channel_report_proto_rawDesc
)

func file_channel_proto_channel_report_proto_rawDescGZIP() []byte {
	file_channel_proto_channel_report_proto_rawDescOnce.Do(func() {
		file_channel_proto_channel_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_channel_report_proto_rawDescData)
	})
	return file_channel_proto_channel_report_proto_rawDescData
}

var file_channel_proto_channel_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_channel_proto_channel_report_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_channel_proto_channel_report_proto_goTypes = []interface{}{
	(GetSensorsReportInfoResponse_TypeOfDisplay)(0), // 0: channel.GetSensorsReportInfoResponse.TypeOfDisplay
	(*SensorIdsRequest)(nil),                        // 1: channel.SensorIdsRequest
	(*CountSensorWarningResponse)(nil),              // 2: channel.CountSensorWarningResponse
	(*GetSensorsReportInfoResponse)(nil),            // 3: channel.GetSensorsReportInfoResponse
	nil,                                             // 4: channel.CountSensorWarningResponse.ResultEntry
	(*GetSensorsReportInfoResponse_Warning)(nil),    // 5: channel.GetSensorsReportInfoResponse.Warning
	(*GetSensorsReportInfoResponse_SensorInfo)(nil), // 6: channel.GetSensorsReportInfoResponse.SensorInfo
	nil, // 7: channel.GetSensorsReportInfoResponse.SensorInfo.TransMapEntry
}
var file_channel_proto_channel_report_proto_depIdxs = []int32{
	4, // 0: channel.CountSensorWarningResponse.result:type_name -> channel.CountSensorWarningResponse.ResultEntry
	6, // 1: channel.GetSensorsReportInfoResponse.sensors:type_name -> channel.GetSensorsReportInfoResponse.SensorInfo
	0, // 2: channel.GetSensorsReportInfoResponse.SensorInfo.displayType:type_name -> channel.GetSensorsReportInfoResponse.TypeOfDisplay
	7, // 3: channel.GetSensorsReportInfoResponse.SensorInfo.transMap:type_name -> channel.GetSensorsReportInfoResponse.SensorInfo.TransMapEntry
	5, // 4: channel.GetSensorsReportInfoResponse.SensorInfo.warningConf:type_name -> channel.GetSensorsReportInfoResponse.Warning
	1, // 5: channel.ReportService.CountSensorsWarning:input_type -> channel.SensorIdsRequest
	1, // 6: channel.ReportService.GetSensorReportInfo:input_type -> channel.SensorIdsRequest
	2, // 7: channel.ReportService.CountSensorsWarning:output_type -> channel.CountSensorWarningResponse
	3, // 8: channel.ReportService.GetSensorReportInfo:output_type -> channel.GetSensorsReportInfoResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_channel_proto_channel_report_proto_init() }
func file_channel_proto_channel_report_proto_init() {
	if File_channel_proto_channel_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_channel_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorIdsRequest); i {
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
		file_channel_proto_channel_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountSensorWarningResponse); i {
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
		file_channel_proto_channel_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorsReportInfoResponse); i {
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
		file_channel_proto_channel_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorsReportInfoResponse_Warning); i {
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
		file_channel_proto_channel_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorsReportInfoResponse_SensorInfo); i {
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
			RawDescriptor: file_channel_proto_channel_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_channel_proto_channel_report_proto_goTypes,
		DependencyIndexes: file_channel_proto_channel_report_proto_depIdxs,
		EnumInfos:         file_channel_proto_channel_report_proto_enumTypes,
		MessageInfos:      file_channel_proto_channel_report_proto_msgTypes,
	}.Build()
	File_channel_proto_channel_report_proto = out.File
	file_channel_proto_channel_report_proto_rawDesc = nil
	file_channel_proto_channel_report_proto_goTypes = nil
	file_channel_proto_channel_report_proto_depIdxs = nil
}
