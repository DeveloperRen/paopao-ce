// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/auth.proto

package v1

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNum string `protobuf:"bytes,1,opt,name=phoneNum,proto3" json:"phoneNum,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

type UserVerify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNum         string `protobuf:"bytes,1,opt,name=phoneNum,proto3" json:"phoneNum,omitempty"`
	VerificationCode string `protobuf:"bytes,2,opt,name=VerificationCode,proto3" json:"VerificationCode,omitempty"`
}

func (x *UserVerify) Reset() {
	*x = UserVerify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVerify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVerify) ProtoMessage() {}

func (x *UserVerify) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVerify.ProtoReflect.Descriptor instead.
func (*UserVerify) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserVerify) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

func (x *UserVerify) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatuCode int32  `protobuf:"varint,1,opt,name=statuCode,proto3" json:"statuCode,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReply) GetStatuCode() int32 {
	if x != nil {
		return x.StatuCode
	}
	return 0
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ActionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *ActionReply) Reset() {
	*x = ActionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionReply) ProtoMessage() {}

func (x *ActionReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionReply.ProtoReflect.Descriptor instead.
func (*ActionReply) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ActionReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_v1_auth_proto protoreflect.FileDescriptor

var file_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x22, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x54, 0x0a, 0x0a, 0x55, 0x73, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x40, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x2d, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x89, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x6f, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x6f, 0x63, 0x62, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x61, 0x6f, 0x70, 0x61, 0x6f,
	0x2d, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0xca, 0x02, 0x04, 0x41,
	0x75, 0x74, 0x68, 0xe2, 0x02, 0x10, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_auth_proto_rawDescOnce sync.Once
	file_v1_auth_proto_rawDescData = file_v1_auth_proto_rawDesc
)

func file_v1_auth_proto_rawDescGZIP() []byte {
	file_v1_auth_proto_rawDescOnce.Do(func() {
		file_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_auth_proto_rawDescData)
	})
	return file_v1_auth_proto_rawDescData
}

var file_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_auth_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: auth.User
	(*UserVerify)(nil),  // 1: auth.UserVerify
	(*LoginReply)(nil),  // 2: auth.LoginReply
	(*ActionReply)(nil), // 3: auth.ActionReply
}
var file_v1_auth_proto_depIdxs = []int32{
	0, // 0: auth.Authenticate.preLogin:input_type -> auth.User
	0, // 1: auth.Authenticate.login:input_type -> auth.User
	0, // 2: auth.Authenticate.logout:input_type -> auth.User
	3, // 3: auth.Authenticate.preLogin:output_type -> auth.ActionReply
	2, // 4: auth.Authenticate.login:output_type -> auth.LoginReply
	3, // 5: auth.Authenticate.logout:output_type -> auth.ActionReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_auth_proto_init() }
func file_v1_auth_proto_init() {
	if File_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVerify); i {
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
		file_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionReply); i {
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
			RawDescriptor: file_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_auth_proto_goTypes,
		DependencyIndexes: file_v1_auth_proto_depIdxs,
		MessageInfos:      file_v1_auth_proto_msgTypes,
	}.Build()
	File_v1_auth_proto = out.File
	file_v1_auth_proto_rawDesc = nil
	file_v1_auth_proto_goTypes = nil
	file_v1_auth_proto_depIdxs = nil
}
