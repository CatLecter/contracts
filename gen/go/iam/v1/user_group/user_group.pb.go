// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: iam/v1/user_group/user_group.proto

package user_group

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

type UserGroupBriefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *UserGroupBriefRequest) Reset() {
	*x = UserGroupBriefRequest{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupBriefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupBriefRequest) ProtoMessage() {}

func (x *UserGroupBriefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupBriefRequest.ProtoReflect.Descriptor instead.
func (*UserGroupBriefRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{0}
}

func (x *UserGroupBriefRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type UserGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Key         string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Protected   string `protobuf:"bytes,5,opt,name=protected,proto3" json:"protected,omitempty"`
}

func (x *UserGroupResponse) Reset() {
	*x = UserGroupResponse{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupResponse) ProtoMessage() {}

func (x *UserGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupResponse.ProtoReflect.Descriptor instead.
func (*UserGroupResponse) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{1}
}

func (x *UserGroupResponse) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UserGroupResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UserGroupResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserGroupResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserGroupResponse) GetProtected() string {
	if x != nil {
		return x.Protected
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page     uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPage  uint32 `protobuf:"varint,1,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	TotalItems uint32 `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetTotalPage() uint32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *Pagination) GetTotalItems() uint32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{4}
}

func (x *Metadata) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []*UserGroupResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Metadata *Metadata            `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResponse) GetData() []*UserGroupResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SearchResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CreateUserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Protected   string `protobuf:"bytes,4,opt,name=protected,proto3" json:"protected,omitempty"`
}

func (x *CreateUserGroupRequest) Reset() {
	*x = CreateUserGroupRequest{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserGroupRequest) ProtoMessage() {}

func (x *CreateUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserGroupRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CreateUserGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateUserGroupRequest) GetProtected() string {
	if x != nil {
		return x.Protected
	}
	return ""
}

type UpdateUserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateUserGroupRequest) Reset() {
	*x = UpdateUserGroupRequest{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserGroupRequest) ProtoMessage() {}

func (x *UpdateUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserGroupRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateUserGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UsersInGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  []string `protobuf:"bytes,1,rep,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId []string `protobuf:"bytes,2,rep,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *UsersInGroups) Reset() {
	*x = UsersInGroups{}
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersInGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersInGroups) ProtoMessage() {}

func (x *UsersInGroups) ProtoReflect() protoreflect.Message {
	mi := &file_iam_v1_user_group_user_group_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersInGroups.ProtoReflect.Descriptor instead.
func (*UsersInGroups) Descriptor() ([]byte, []int) {
	return file_iam_v1_user_group_user_group_proto_rawDescGZIP(), []int{8}
}

func (x *UsersInGroups) GetUserId() []string {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UsersInGroups) GetGroupId() []string {
	if x != nil {
		return x.GroupId
	}
	return nil
}

var File_iam_v1_user_group_user_group_proto protoreflect.FileDescriptor

var file_iam_v1_user_group_user_group_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x61, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x22, 0x56, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4c, 0x0a, 0x0a, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3b, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7e, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x60, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x43, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x32, 0xae, 0x03, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x4c, 0x65, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iam_v1_user_group_user_group_proto_rawDescOnce sync.Once
	file_iam_v1_user_group_user_group_proto_rawDescData = file_iam_v1_user_group_user_group_proto_rawDesc
)

func file_iam_v1_user_group_user_group_proto_rawDescGZIP() []byte {
	file_iam_v1_user_group_user_group_proto_rawDescOnce.Do(func() {
		file_iam_v1_user_group_user_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_v1_user_group_user_group_proto_rawDescData)
	})
	return file_iam_v1_user_group_user_group_proto_rawDescData
}

var file_iam_v1_user_group_user_group_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_iam_v1_user_group_user_group_proto_goTypes = []any{
	(*UserGroupBriefRequest)(nil),  // 0: iam.UserGroupBriefRequest
	(*UserGroupResponse)(nil),      // 1: iam.UserGroupResponse
	(*SearchRequest)(nil),          // 2: iam.SearchRequest
	(*Pagination)(nil),             // 3: iam.Pagination
	(*Metadata)(nil),               // 4: iam.Metadata
	(*SearchResponse)(nil),         // 5: iam.SearchResponse
	(*CreateUserGroupRequest)(nil), // 6: iam.CreateUserGroupRequest
	(*UpdateUserGroupRequest)(nil), // 7: iam.UpdateUserGroupRequest
	(*UsersInGroups)(nil),          // 8: iam.UsersInGroups
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_iam_v1_user_group_user_group_proto_depIdxs = []int32{
	3,  // 0: iam.Metadata.pagination:type_name -> iam.Pagination
	1,  // 1: iam.SearchResponse.data:type_name -> iam.UserGroupResponse
	4,  // 2: iam.SearchResponse.metadata:type_name -> iam.Metadata
	0,  // 3: iam.UserGroups.Get:input_type -> iam.UserGroupBriefRequest
	2,  // 4: iam.UserGroups.Search:input_type -> iam.SearchRequest
	6,  // 5: iam.UserGroups.Create:input_type -> iam.CreateUserGroupRequest
	7,  // 6: iam.UserGroups.Update:input_type -> iam.UpdateUserGroupRequest
	0,  // 7: iam.UserGroups.Delete:input_type -> iam.UserGroupBriefRequest
	8,  // 8: iam.UserGroups.AddingUsers:input_type -> iam.UsersInGroups
	8,  // 9: iam.UserGroups.DeletingUsers:input_type -> iam.UsersInGroups
	1,  // 10: iam.UserGroups.Get:output_type -> iam.UserGroupResponse
	5,  // 11: iam.UserGroups.Search:output_type -> iam.SearchResponse
	9,  // 12: iam.UserGroups.Create:output_type -> google.protobuf.Empty
	9,  // 13: iam.UserGroups.Update:output_type -> google.protobuf.Empty
	9,  // 14: iam.UserGroups.Delete:output_type -> google.protobuf.Empty
	9,  // 15: iam.UserGroups.AddingUsers:output_type -> google.protobuf.Empty
	9,  // 16: iam.UserGroups.DeletingUsers:output_type -> google.protobuf.Empty
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_iam_v1_user_group_user_group_proto_init() }
func file_iam_v1_user_group_user_group_proto_init() {
	if File_iam_v1_user_group_user_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iam_v1_user_group_user_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iam_v1_user_group_user_group_proto_goTypes,
		DependencyIndexes: file_iam_v1_user_group_user_group_proto_depIdxs,
		MessageInfos:      file_iam_v1_user_group_user_group_proto_msgTypes,
	}.Build()
	File_iam_v1_user_group_user_group_proto = out.File
	file_iam_v1_user_group_user_group_proto_rawDesc = nil
	file_iam_v1_user_group_user_group_proto_goTypes = nil
	file_iam_v1_user_group_user_group_proto_depIdxs = nil
}