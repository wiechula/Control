// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/flp.proto

package bkpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FlpCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HostName  string `protobuf:"bytes,2,opt,name=hostName,proto3" json:"hostName,omitempty"`
	RunNumber int32  `protobuf:"varint,3,opt,name=runNumber,proto3" json:"runNumber,omitempty"`
}

func (x *FlpCreationRequest) Reset() {
	*x = FlpCreationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_flp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlpCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlpCreationRequest) ProtoMessage() {}

func (x *FlpCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_flp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlpCreationRequest.ProtoReflect.Descriptor instead.
func (*FlpCreationRequest) Descriptor() ([]byte, []int) {
	return file_protos_flp_proto_rawDescGZIP(), []int{0}
}

func (x *FlpCreationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlpCreationRequest) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *FlpCreationRequest) GetRunNumber() int32 {
	if x != nil {
		return x.RunNumber
	}
	return 0
}

type ManyFlpsCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flps []*FlpCreationRequest `protobuf:"bytes,1,rep,name=flps,proto3" json:"flps,omitempty"`
}

func (x *ManyFlpsCreationRequest) Reset() {
	*x = ManyFlpsCreationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_flp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyFlpsCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyFlpsCreationRequest) ProtoMessage() {}

func (x *ManyFlpsCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_flp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyFlpsCreationRequest.ProtoReflect.Descriptor instead.
func (*ManyFlpsCreationRequest) Descriptor() ([]byte, []int) {
	return file_protos_flp_proto_rawDescGZIP(), []int{1}
}

func (x *ManyFlpsCreationRequest) GetFlps() []*FlpCreationRequest {
	if x != nil {
		return x.Flps
	}
	return nil
}

type FlpList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flps []*Flp `protobuf:"bytes,1,rep,name=flps,proto3" json:"flps,omitempty"`
}

func (x *FlpList) Reset() {
	*x = FlpList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_flp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlpList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlpList) ProtoMessage() {}

func (x *FlpList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_flp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlpList.ProtoReflect.Descriptor instead.
func (*FlpList) Descriptor() ([]byte, []int) {
	return file_protos_flp_proto_rawDescGZIP(), []int{2}
}

func (x *FlpList) GetFlps() []*Flp {
	if x != nil {
		return x.Flps
	}
	return nil
}

type Flp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Unix timestamp when this entity was created.
	CreatedAt int64 `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// Unix timestamp when this entity was last updated.
	UpdatedAt             int64 `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	BytesEquipmentReadOut int32 `protobuf:"varint,6,opt,name=bytesEquipmentReadOut,proto3" json:"bytesEquipmentReadOut,omitempty"`
	BytesFairMQReadOut    int32 `protobuf:"varint,7,opt,name=bytesFairMQReadOut,proto3" json:"bytesFairMQReadOut,omitempty"`
	BytesProcessed        int32 `protobuf:"varint,8,opt,name=bytesProcessed,proto3" json:"bytesProcessed,omitempty"`
	BytesRecordingReadOut int32 `protobuf:"varint,9,opt,name=bytesRecordingReadOut,proto3" json:"bytesRecordingReadOut,omitempty"`
	NTimeframes           int32 `protobuf:"varint,10,opt,name=nTimeframes,proto3" json:"nTimeframes,omitempty"`
}

func (x *Flp) Reset() {
	*x = Flp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_flp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flp) ProtoMessage() {}

func (x *Flp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_flp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flp.ProtoReflect.Descriptor instead.
func (*Flp) Descriptor() ([]byte, []int) {
	return file_protos_flp_proto_rawDescGZIP(), []int{3}
}

func (x *Flp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Flp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flp) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Flp) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Flp) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Flp) GetBytesEquipmentReadOut() int32 {
	if x != nil {
		return x.BytesEquipmentReadOut
	}
	return 0
}

func (x *Flp) GetBytesFairMQReadOut() int32 {
	if x != nil {
		return x.BytesFairMQReadOut
	}
	return 0
}

func (x *Flp) GetBytesProcessed() int32 {
	if x != nil {
		return x.BytesProcessed
	}
	return 0
}

func (x *Flp) GetBytesRecordingReadOut() int32 {
	if x != nil {
		return x.BytesRecordingReadOut
	}
	return 0
}

func (x *Flp) GetNTimeframes() int32 {
	if x != nil {
		return x.NTimeframes
	}
	return 0
}

var File_protos_flp_proto protoreflect.FileDescriptor

var file_protos_flp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x6c, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x62, 0x0a, 0x12, 0x46, 0x6c, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x75, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x17, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x6c,
	0x70, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x04, 0x66, 0x6c, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x46, 0x6c, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x66, 0x6c, 0x70, 0x73, 0x22, 0x32, 0x0a, 0x07, 0x46, 0x6c, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x66, 0x6c, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6c, 0x70, 0x52, 0x04, 0x66, 0x6c, 0x70, 0x73, 0x22, 0xe7, 0x02,
	0x0a, 0x03, 0x46, 0x6c, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x34, 0x0a, 0x15, 0x62, 0x79, 0x74, 0x65, 0x73, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x62, 0x79, 0x74, 0x65, 0x73, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x46, 0x61, 0x69, 0x72, 0x4d, 0x51, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x12, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x61, 0x69, 0x72, 0x4d, 0x51,
	0x52, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x15, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x61, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x5c, 0x0a, 0x0a, 0x46, 0x6c, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x6e, 0x79, 0x12, 0x27, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x6c, 0x70, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f,
	0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6c,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x4f, 0x32, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65,
	0x65, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x62, 0x6b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_flp_proto_rawDescOnce sync.Once
	file_protos_flp_proto_rawDescData = file_protos_flp_proto_rawDesc
)

func file_protos_flp_proto_rawDescGZIP() []byte {
	file_protos_flp_proto_rawDescOnce.Do(func() {
		file_protos_flp_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_flp_proto_rawDescData)
	})
	return file_protos_flp_proto_rawDescData
}

var file_protos_flp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_flp_proto_goTypes = []interface{}{
	(*FlpCreationRequest)(nil),      // 0: o2.bookkeeping.FlpCreationRequest
	(*ManyFlpsCreationRequest)(nil), // 1: o2.bookkeeping.ManyFlpsCreationRequest
	(*FlpList)(nil),                 // 2: o2.bookkeeping.FlpList
	(*Flp)(nil),                     // 3: o2.bookkeeping.Flp
}
var file_protos_flp_proto_depIdxs = []int32{
	0, // 0: o2.bookkeeping.ManyFlpsCreationRequest.flps:type_name -> o2.bookkeeping.FlpCreationRequest
	3, // 1: o2.bookkeeping.FlpList.flps:type_name -> o2.bookkeeping.Flp
	1, // 2: o2.bookkeeping.FlpService.CreateMany:input_type -> o2.bookkeeping.ManyFlpsCreationRequest
	2, // 3: o2.bookkeeping.FlpService.CreateMany:output_type -> o2.bookkeeping.FlpList
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_flp_proto_init() }
func file_protos_flp_proto_init() {
	if File_protos_flp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_flp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlpCreationRequest); i {
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
		file_protos_flp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyFlpsCreationRequest); i {
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
		file_protos_flp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlpList); i {
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
		file_protos_flp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flp); i {
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
			RawDescriptor: file_protos_flp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_flp_proto_goTypes,
		DependencyIndexes: file_protos_flp_proto_depIdxs,
		MessageInfos:      file_protos_flp_proto_msgTypes,
	}.Build()
	File_protos_flp_proto = out.File
	file_protos_flp_proto_rawDesc = nil
	file_protos_flp_proto_goTypes = nil
	file_protos_flp_proto_depIdxs = nil
}
