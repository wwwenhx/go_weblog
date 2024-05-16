// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: userService.proto

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

type UserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tah: json:"id" form:"id"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tah: json:"user_name" form:"user_name"
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// @inject_tah: json:"created_at" form:"created_at"
	CreatedAt int64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tah: json:"updated_at" form:"updated_at"
	UpdatedAt int64 `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// @inject_tah: json:"delete_at" form:"delete_at"
	DeleteAt int64 `protobuf:"varint,5,opt,name=delete_at,json=deleteAt,proto3" json:"delete_at,omitempty"`
	// @inject_tah: json:"password" form:"password"
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{0}
}

func (x *UserModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserModel) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *UserModel) GetDeleteAt() int64 {
	if x != nil {
		return x.DeleteAt
	}
	return 0
}

func (x *UserModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tah: json:"user_name" form:"user_name"
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// @inject_tah: json:"password" form:"password"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{1}
}

func (x *UserReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tah: json:"user_detail" form:"user_detail"
	UserDetail *UserModel   `protobuf:"bytes,1,opt,name=user_detail,json=userDetail,proto3" json:"user_detail,omitempty"`
	UserList   []*UserModel `protobuf:"bytes,2,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
	// @inject_tah: json:"code" form:"code"
	Code uint32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UserRes) Reset() {
	*x = UserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRes) ProtoMessage() {}

func (x *UserRes) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRes.ProtoReflect.Descriptor instead.
func (*UserRes) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{2}
}

func (x *UserRes) GetUserDetail() *UserModel {
	if x != nil {
		return x.UserDetail
	}
	return nil
}

func (x *UserRes) GetUserList() []*UserModel {
	if x != nil {
		return x.UserList
	}
	return nil
}

func (x *UserRes) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_userService_proto protoreflect.FileDescriptor

var file_userService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x73, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x27, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x6d,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userService_proto_rawDescOnce sync.Once
	file_userService_proto_rawDescData = file_userService_proto_rawDesc
)

func file_userService_proto_rawDescGZIP() []byte {
	file_userService_proto_rawDescOnce.Do(func() {
		file_userService_proto_rawDescData = protoimpl.X.CompressGZIP(file_userService_proto_rawDescData)
	})
	return file_userService_proto_rawDescData
}

var file_userService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_userService_proto_goTypes = []interface{}{
	(*UserModel)(nil), // 0: UserModel
	(*UserReq)(nil),   // 1: UserReq
	(*UserRes)(nil),   // 2: UserRes
}
var file_userService_proto_depIdxs = []int32{
	0, // 0: UserRes.user_detail:type_name -> UserModel
	0, // 1: UserRes.user_list:type_name -> UserModel
	1, // 2: User.UserLogin:input_type -> UserReq
	1, // 3: User.UserRegister:input_type -> UserReq
	1, // 4: User.GetAllUser:input_type -> UserReq
	2, // 5: User.UserLogin:output_type -> UserRes
	2, // 6: User.UserRegister:output_type -> UserRes
	2, // 7: User.GetAllUser:output_type -> UserRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userService_proto_init() }
func file_userService_proto_init() {
	if File_userService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModel); i {
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
		file_userService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
		file_userService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRes); i {
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
			RawDescriptor: file_userService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userService_proto_goTypes,
		DependencyIndexes: file_userService_proto_depIdxs,
		MessageInfos:      file_userService_proto_msgTypes,
	}.Build()
	File_userService_proto = out.File
	file_userService_proto_rawDesc = nil
	file_userService_proto_goTypes = nil
	file_userService_proto_depIdxs = nil
}