// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.2
// source: pb/common/common.proto

package common

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

type Code int32

const (
	Code_Err Code = 0
	Code_Ok  Code = 200
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:   "Err",
		200: "Ok",
	}
	Code_value = map[string]int32{
		"Err": 0,
		"Ok":  200,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_common_common_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_pb_common_common_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_pb_common_common_proto_rawDescGZIP(), []int{0}
}

type EventType int32

const (
	EventType_EventTypeNone EventType = 0
	EventType_UserRegister  EventType = 1
	EventType_Mail          EventType = 2
	EventType_Like          EventType = 3
	EventType_DisLike       EventType = 4
	EventType_Follow        EventType = 5
	EventType_UnFollow      EventType = 6
	EventType_CommentVideo  EventType = 7
	EventType_CalculateHot  EventType = 8
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EventTypeNone",
		1: "UserRegister",
		2: "Mail",
		3: "Like",
		4: "DisLike",
		5: "Follow",
		6: "UnFollow",
		7: "CommentVideo",
		8: "CalculateHot",
	}
	EventType_value = map[string]int32{
		"EventTypeNone": 0,
		"UserRegister":  1,
		"Mail":          2,
		"Like":          3,
		"DisLike":       4,
		"Follow":        5,
		"UnFollow":      6,
		"CommentVideo":  7,
		"CalculateHot":  8,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_common_common_proto_enumTypes[1].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_pb_common_common_proto_enumTypes[1]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_pb_common_common_proto_rawDescGZIP(), []int{1}
}

type TargetType int32

const (
	TargetType_TargetTypeNone TargetType = 0
	TargetType_Video          TargetType = 1
	TargetType_Comment        TargetType = 2
	TargetType_SubComment     TargetType = 3
)

// Enum value maps for TargetType.
var (
	TargetType_name = map[int32]string{
		0: "TargetTypeNone",
		1: "Video",
		2: "Comment",
		3: "SubComment",
	}
	TargetType_value = map[string]int32{
		"TargetTypeNone": 0,
		"Video":          1,
		"Comment":        2,
		"SubComment":     3,
	}
)

func (x TargetType) Enum() *TargetType {
	p := new(TargetType)
	*p = x
	return p
}

func (x TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_common_common_proto_enumTypes[2].Descriptor()
}

func (TargetType) Type() protoreflect.EnumType {
	return &file_pb_common_common_proto_enumTypes[2]
}

func (x TargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetType.Descriptor instead.
func (TargetType) EnumDescriptor() ([]byte, []int) {
	return file_pb_common_common_proto_rawDescGZIP(), []int{2}
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason   string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_pb_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_pb_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Res) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Res) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Res) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Res) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_pb_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType     EventType  `protobuf:"varint,1,opt,name=eventType,proto3,enum=EventType" json:"eventType,omitempty"`
	TargetId      int64      `protobuf:"varint,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType    TargetType `protobuf:"varint,3,opt,name=targetType,proto3,enum=TargetType" json:"targetType,omitempty"`
	ActorId       int64      `protobuf:"varint,4,opt,name=ActorId,proto3" json:"ActorId,omitempty"`
	TargetOwnerId int64      `protobuf:"varint,5,opt,name=targetOwnerId,proto3" json:"targetOwnerId,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pb_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pb_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EventTypeNone
}

func (x *Event) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *Event) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_TargetTypeNone
}

func (x *Event) GetActorId() int64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *Event) GetTargetOwnerId() int64 {
	if x != nil {
		return x.TargetOwnerId
	}
	return 0
}

var File_pb_common_common_proto protoreflect.FileDescriptor

var file_pb_common_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x03, 0x52, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x18, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x45, 0x72, 0x72, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0xc8, 0x01,
	0x2a, 0x8f, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x4c, 0x69, 0x6b, 0x65, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x4c, 0x69,
	0x6b, 0x65, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x10, 0x06, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x07,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74,
	0x10, 0x08, 0x2a, 0x48, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x42, 0x1c, 0x5a, 0x1a,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_common_common_proto_rawDescOnce sync.Once
	file_pb_common_common_proto_rawDescData = file_pb_common_common_proto_rawDesc
)

func file_pb_common_common_proto_rawDescGZIP() []byte {
	file_pb_common_common_proto_rawDescOnce.Do(func() {
		file_pb_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_common_common_proto_rawDescData)
	})
	return file_pb_common_common_proto_rawDescData
}

var file_pb_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pb_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_common_common_proto_goTypes = []interface{}{
	(Code)(0),       // 0: Code
	(EventType)(0),  // 1: EventType
	(TargetType)(0), // 2: TargetType
	(*Res)(nil),     // 3: Res
	(*ID)(nil),      // 4: ID
	(*Event)(nil),   // 5: Event
	nil,             // 6: Res.MetadataEntry
}
var file_pb_common_common_proto_depIdxs = []int32{
	6, // 0: Res.metadata:type_name -> Res.MetadataEntry
	1, // 1: Event.eventType:type_name -> EventType
	2, // 2: Event.targetType:type_name -> TargetType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_common_common_proto_init() }
func file_pb_common_common_proto_init() {
	if File_pb_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
		file_pb_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_pb_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_pb_common_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_common_common_proto_goTypes,
		DependencyIndexes: file_pb_common_common_proto_depIdxs,
		EnumInfos:         file_pb_common_common_proto_enumTypes,
		MessageInfos:      file_pb_common_common_proto_msgTypes,
	}.Build()
	File_pb_common_common_proto = out.File
	file_pb_common_common_proto_rawDesc = nil
	file_pb_common_common_proto_goTypes = nil
	file_pb_common_common_proto_depIdxs = nil
}
