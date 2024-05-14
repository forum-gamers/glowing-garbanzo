// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: community.proto

package community

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

type CreateCommunityInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl      string `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	ImageId       string `protobuf:"bytes,3,opt,name=imageId,proto3" json:"imageId,omitempty"`
	Desc          string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	BackgroundUrl string `protobuf:"bytes,5,opt,name=backgroundUrl,proto3" json:"backgroundUrl,omitempty"`
	BackgroundId  string `protobuf:"bytes,6,opt,name=backgroundId,proto3" json:"backgroundId,omitempty"`
}

func (x *CreateCommunityInput) Reset() {
	*x = CreateCommunityInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommunityInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommunityInput) ProtoMessage() {}

func (x *CreateCommunityInput) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommunityInput.ProtoReflect.Descriptor instead.
func (*CreateCommunityInput) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCommunityInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCommunityInput) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateCommunityInput) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *CreateCommunityInput) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateCommunityInput) GetBackgroundUrl() string {
	if x != nil {
		return x.BackgroundUrl
	}
	return ""
}

func (x *CreateCommunityInput) GetBackgroundId() string {
	if x != nil {
		return x.BackgroundId
	}
	return ""
}

type Community struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl      string `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	ImageId       string `protobuf:"bytes,4,opt,name=imageId,proto3" json:"imageId,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Owner         string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	CreatedAt     string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     string `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	BackgroundUrl string `protobuf:"bytes,9,opt,name=backgroundUrl,proto3" json:"backgroundUrl,omitempty"`
	BackgroundId  string `protobuf:"bytes,10,opt,name=backgroundId,proto3" json:"backgroundId,omitempty"`
}

func (x *Community) Reset() {
	*x = Community{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Community) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Community) ProtoMessage() {}

func (x *Community) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Community.ProtoReflect.Descriptor instead.
func (*Community) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{1}
}

func (x *Community) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Community) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Community) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Community) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Community) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Community) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Community) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Community) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Community) GetBackgroundUrl() string {
	if x != nil {
		return x.BackgroundUrl
	}
	return ""
}

func (x *Community) GetBackgroundId() string {
	if x != nil {
		return x.BackgroundId
	}
	return ""
}

type ImageIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId      string `protobuf:"bytes,1,opt,name=imageId,proto3" json:"imageId,omitempty"`
	BackgroundId string `protobuf:"bytes,2,opt,name=backgroundId,proto3" json:"backgroundId,omitempty"`
}

func (x *ImageIdResp) Reset() {
	*x = ImageIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageIdResp) ProtoMessage() {}

func (x *ImageIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageIdResp.ProtoReflect.Descriptor instead.
func (*ImageIdResp) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{2}
}

func (x *ImageIdResp) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ImageIdResp) GetBackgroundId() string {
	if x != nil {
		return x.BackgroundId
	}
	return ""
}

type CommunityIdInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty"`
}

func (x *CommunityIdInput) Reset() {
	*x = CommunityIdInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityIdInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityIdInput) ProtoMessage() {}

func (x *CommunityIdInput) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[3]
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
	return file_community_proto_rawDescGZIP(), []int{3}
}

func (x *CommunityIdInput) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

type UpdateImgInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	CommunityId string `protobuf:"bytes,3,opt,name=communityId,proto3" json:"communityId,omitempty"`
}

func (x *UpdateImgInput) Reset() {
	*x = UpdateImgInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImgInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImgInput) ProtoMessage() {}

func (x *UpdateImgInput) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImgInput.ProtoReflect.Descriptor instead.
func (*UpdateImgInput) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateImgInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateImgInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateImgInput) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

type TextInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text        string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	CommunityId string `protobuf:"bytes,2,opt,name=communityId,proto3" json:"communityId,omitempty"`
}

func (x *TextInput) Reset() {
	*x = TextInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextInput) ProtoMessage() {}

func (x *TextInput) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextInput.ProtoReflect.Descriptor instead.
func (*TextInput) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{5}
}

func (x *TextInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TextInput) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

type ChangeOwnershipInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty"`
	TargetId    string `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
}

func (x *ChangeOwnershipInput) Reset() {
	*x = ChangeOwnershipInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeOwnershipInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeOwnershipInput) ProtoMessage() {}

func (x *ChangeOwnershipInput) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeOwnershipInput.ProtoReflect.Descriptor instead.
func (*ChangeOwnershipInput) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeOwnershipInput) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *ChangeOwnershipInput) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

var File_community_proto protoreflect.FileDescriptor

var file_community_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x22, 0xbe, 0x01, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x22, 0xa3, 0x02,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x22, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64,
	0x22, 0x34, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6d, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x09,
	0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22,
	0x54, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x32, 0xb9, 0x03, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6d, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_community_proto_rawDescOnce sync.Once
	file_community_proto_rawDescData = file_community_proto_rawDesc
)

func file_community_proto_rawDescGZIP() []byte {
	file_community_proto_rawDescOnce.Do(func() {
		file_community_proto_rawDescData = protoimpl.X.CompressGZIP(file_community_proto_rawDescData)
	})
	return file_community_proto_rawDescData
}

var file_community_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_community_proto_goTypes = []interface{}{
	(*CreateCommunityInput)(nil), // 0: community.CreateCommunityInput
	(*Community)(nil),            // 1: community.Community
	(*ImageIdResp)(nil),          // 2: community.ImageIdResp
	(*CommunityIdInput)(nil),     // 3: community.CommunityIdInput
	(*UpdateImgInput)(nil),       // 4: community.UpdateImgInput
	(*TextInput)(nil),            // 5: community.TextInput
	(*ChangeOwnershipInput)(nil), // 6: community.ChangeOwnershipInput
}
var file_community_proto_depIdxs = []int32{
	0, // 0: community.CommunityService.CreateCommunity:input_type -> community.CreateCommunityInput
	3, // 1: community.CommunityService.DeleteCommunity:input_type -> community.CommunityIdInput
	4, // 2: community.CommunityService.UpdateImage:input_type -> community.UpdateImgInput
	4, // 3: community.CommunityService.UpdateBackground:input_type -> community.UpdateImgInput
	5, // 4: community.CommunityService.UpdateDesc:input_type -> community.TextInput
	6, // 5: community.CommunityService.ChangeOwnership:input_type -> community.ChangeOwnershipInput
	1, // 6: community.CommunityService.CreateCommunity:output_type -> community.Community
	2, // 7: community.CommunityService.DeleteCommunity:output_type -> community.ImageIdResp
	1, // 8: community.CommunityService.UpdateImage:output_type -> community.Community
	1, // 9: community.CommunityService.UpdateBackground:output_type -> community.Community
	1, // 10: community.CommunityService.UpdateDesc:output_type -> community.Community
	1, // 11: community.CommunityService.ChangeOwnership:output_type -> community.Community
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_community_proto_init() }
func file_community_proto_init() {
	if File_community_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_community_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommunityInput); i {
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
		file_community_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Community); i {
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
		file_community_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageIdResp); i {
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
		file_community_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_community_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImgInput); i {
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
		file_community_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextInput); i {
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
		file_community_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeOwnershipInput); i {
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
			RawDescriptor: file_community_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_community_proto_goTypes,
		DependencyIndexes: file_community_proto_depIdxs,
		MessageInfos:      file_community_proto_msgTypes,
	}.Build()
	File_community_proto = out.File
	file_community_proto_rawDesc = nil
	file_community_proto_goTypes = nil
	file_community_proto_depIdxs = nil
}
