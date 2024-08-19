// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: msgService.proto

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

type MsgModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tah: json:"id" form:"id"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tah: json:"user_id" form:"user_id"
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tah: json:"created_at" form:"created_at"
	CreatedAt int64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tah: json:"updated_at" form:"updated_at"
	UpdatedAt int64 `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// @inject_tah: json:"delete_at" form:"delete_at"
	DeleteAt int64 `protobuf:"varint,5,opt,name=delete_at,json=deleteAt,proto3" json:"delete_at,omitempty"`
	// @inject_tah: json:"from" form:"from"
	From uint32 `protobuf:"varint,6,opt,name=from,proto3" json:"from,omitempty"`
	// @inject_tah: json:"body" form:"body"
	Body string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	// @inject_tah: json:"to" form:"to"
	To uint32 `protobuf:"varint,8,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MsgModel) Reset() {
	*x = MsgModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgModel) ProtoMessage() {}

func (x *MsgModel) ProtoReflect() protoreflect.Message {
	mi := &file_msgService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgModel.ProtoReflect.Descriptor instead.
func (*MsgModel) Descriptor() ([]byte, []int) {
	return file_msgService_proto_rawDescGZIP(), []int{0}
}

func (x *MsgModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgModel) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *MsgModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MsgModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *MsgModel) GetDeleteAt() int64 {
	if x != nil {
		return x.DeleteAt
	}
	return 0
}

func (x *MsgModel) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *MsgModel) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MsgModel) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

type MsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tah: json:"id" form:"id"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tah: json:"user_id" form:"user_id"
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tah: json:"body" form:"body"
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// @inject_tah: json:"from" form:"from"
	From uint32 `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	// @inject_tah: json:"to" form:"to"
	To uint32 `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MsgReq) Reset() {
	*x = MsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReq) ProtoMessage() {}

func (x *MsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msgService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReq.ProtoReflect.Descriptor instead.
func (*MsgReq) Descriptor() ([]byte, []int) {
	return file_msgService_proto_rawDescGZIP(), []int{1}
}

func (x *MsgReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *MsgReq) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MsgReq) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *MsgReq) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

type MsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tah: json:"code" form:"code"
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tah: json:"msg_list" form:"msg_list"
	MsgList []*MsgModel `protobuf:"bytes,2,rep,name=msg_list,json=msgList,proto3" json:"msg_list,omitempty"`
	// @inject_tah: json:"msg" form:"msg"
	Msg *MsgModel `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MsgRes) Reset() {
	*x = MsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRes) ProtoMessage() {}

func (x *MsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_msgService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRes.ProtoReflect.Descriptor instead.
func (*MsgRes) Descriptor() ([]byte, []int) {
	return file_msgService_proto_rawDescGZIP(), []int{2}
}

func (x *MsgRes) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgRes) GetMsgList() []*MsgModel {
	if x != nil {
		return x.MsgList
	}
	return nil
}

func (x *MsgRes) GetMsg() *MsgModel {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_msgService_proto protoreflect.FileDescriptor

var file_msgService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x06, 0x6d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x3e, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1a,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x07, 0x2e, 0x6d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x07, 0x2e, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x07, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x07, 0x2e, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x07,
	0x2e, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgService_proto_rawDescOnce sync.Once
	file_msgService_proto_rawDescData = file_msgService_proto_rawDesc
)

func file_msgService_proto_rawDescGZIP() []byte {
	file_msgService_proto_rawDescOnce.Do(func() {
		file_msgService_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgService_proto_rawDescData)
	})
	return file_msgService_proto_rawDescData
}

var file_msgService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_msgService_proto_goTypes = []interface{}{
	(*MsgModel)(nil), // 0: MsgModel
	(*MsgReq)(nil),   // 1: msgReq
	(*MsgRes)(nil),   // 2: msgRes
}
var file_msgService_proto_depIdxs = []int32{
	0, // 0: msgRes.msg_list:type_name -> MsgModel
	0, // 1: msgRes.msg:type_name -> MsgModel
	1, // 2: Msg.GetMsg:input_type -> msgReq
	1, // 3: Msg.SendMsg:input_type -> msgReq
	2, // 4: Msg.GetMsg:output_type -> msgRes
	2, // 5: Msg.SendMsg:output_type -> msgRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_msgService_proto_init() }
func file_msgService_proto_init() {
	if File_msgService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msgService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgModel); i {
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
		file_msgService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReq); i {
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
		file_msgService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRes); i {
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
			RawDescriptor: file_msgService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msgService_proto_goTypes,
		DependencyIndexes: file_msgService_proto_depIdxs,
		MessageInfos:      file_msgService_proto_msgTypes,
	}.Build()
	File_msgService_proto = out.File
	file_msgService_proto_rawDesc = nil
	file_msgService_proto_goTypes = nil
	file_msgService_proto_depIdxs = nil
}
