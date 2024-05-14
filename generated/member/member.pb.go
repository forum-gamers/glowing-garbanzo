// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: member.proto

package member

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

type CommunityIdInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty"`
}

func (x *CommunityIdInput) Reset() {
	*x = CommunityIdInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityIdInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityIdInput) ProtoMessage() {}

func (x *CommunityIdInput) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunityIdInput.ProtoReflect.Descriptor instead.
func (*CommunityIdInput) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{0}
}

func (x *CommunityIdInput) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_member_proto protoreflect.FileDescriptor

var file_member_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x8c, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f,
	0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_proto_rawDescOnce sync.Once
	file_member_proto_rawDescData = file_member_proto_rawDesc
)

func file_member_proto_rawDescGZIP() []byte {
	file_member_proto_rawDescOnce.Do(func() {
		file_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_proto_rawDescData)
	})
	return file_member_proto_rawDescData
}

var file_member_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_member_proto_goTypes = []interface{}{
	(*CommunityIdInput)(nil), // 0: member.CommunityIdInput
	(*Message)(nil),          // 1: member.Message
}
var file_member_proto_depIdxs = []int32{
	0, // 0: member.MemberService.JoinCommunity:input_type -> member.CommunityIdInput
	0, // 1: member.MemberService.LeaveCommunity:input_type -> member.CommunityIdInput
	1, // 2: member.MemberService.JoinCommunity:output_type -> member.Message
	1, // 3: member.MemberService.LeaveCommunity:output_type -> member.Message
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_member_proto_init() }
func file_member_proto_init() {
	if File_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityIdInput); i {
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
		file_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_member_proto_goTypes,
		DependencyIndexes: file_member_proto_depIdxs,
		MessageInfos:      file_member_proto_msgTypes,
	}.Build()
	File_member_proto = out.File
	file_member_proto_rawDesc = nil
	file_member_proto_goTypes = nil
	file_member_proto_depIdxs = nil
}
