// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: template.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type PageTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	ServerId int64 `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *PageTemplateRequest) Reset() {
	*x = PageTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTemplateRequest) ProtoMessage() {}

func (x *PageTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTemplateRequest.ProtoReflect.Descriptor instead.
func (*PageTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0}
}

func (x *PageTemplateRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageTemplateRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageTemplateRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type PageTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*PageTemplateReply_Template `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageTemplateReply) Reset() {
	*x = PageTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTemplateReply) ProtoMessage() {}

func (x *PageTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTemplateReply.ProtoReflect.Descriptor instead.
func (*PageTemplateReply) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{1}
}

func (x *PageTemplateReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageTemplateReply) GetList() []*PageTemplateReply_Template {
	if x != nil {
		return x.List
	}
	return nil
}

// 查询指定模板
type CurrentTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int64 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *CurrentTemplateRequest) Reset() {
	*x = CurrentTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentTemplateRequest) ProtoMessage() {}

func (x *CurrentTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentTemplateRequest.ProtoReflect.Descriptor instead.
func (*CurrentTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{2}
}

func (x *CurrentTemplateRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type CurrentTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId    int64   `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Content     string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Version     string  `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	IsUse       bool    `protobuf:"varint,6,opt,name=is_use,json=isUse,proto3" json:"is_use,omitempty"`
	Operator    *string `protobuf:"bytes,7,opt,name=operator,proto3,oneof" json:"operator,omitempty"`
	OperatorId  *int64  `protobuf:"varint,8,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"`
}

func (x *CurrentTemplateReply) Reset() {
	*x = CurrentTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentTemplateReply) ProtoMessage() {}

func (x *CurrentTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentTemplateReply.ProtoReflect.Descriptor instead.
func (*CurrentTemplateReply) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{3}
}

func (x *CurrentTemplateReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CurrentTemplateReply) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *CurrentTemplateReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CurrentTemplateReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CurrentTemplateReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CurrentTemplateReply) GetIsUse() bool {
	if x != nil {
		return x.IsUse
	}
	return false
}

func (x *CurrentTemplateReply) GetOperator() string {
	if x != nil && x.Operator != nil {
		return *x.Operator
	}
	return ""
}

func (x *CurrentTemplateReply) GetOperatorId() int64 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 查询指定模板
type GetTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTemplateRequest) Reset() {
	*x = GetTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateRequest) ProtoMessage() {}

func (x *GetTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{4}
}

func (x *GetTemplateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId    int64   `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Content     string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Version     string  `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	IsUse       *bool   `protobuf:"varint,6,opt,name=is_use,json=isUse,proto3,oneof" json:"is_use,omitempty"`
	Operator    *string `protobuf:"bytes,7,opt,name=operator,proto3,oneof" json:"operator,omitempty"`
	OperatorId  *int64  `protobuf:"varint,8,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"`
}

func (x *GetTemplateReply) Reset() {
	*x = GetTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateReply) ProtoMessage() {}

func (x *GetTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateReply.ProtoReflect.Descriptor instead.
func (*GetTemplateReply) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{5}
}

func (x *GetTemplateReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetTemplateReply) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *GetTemplateReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetTemplateReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetTemplateReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetTemplateReply) GetIsUse() bool {
	if x != nil && x.IsUse != nil {
		return *x.IsUse
	}
	return false
}

func (x *GetTemplateReply) GetOperator() string {
	if x != nil && x.Operator != nil {
		return *x.Operator
	}
	return ""
}

func (x *GetTemplateReply) GetOperatorId() int64 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 新增模板
type AddTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId    int64  `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Content     string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddTemplateRequest) Reset() {
	*x = AddTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTemplateRequest) ProtoMessage() {}

func (x *AddTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTemplateRequest.ProtoReflect.Descriptor instead.
func (*AddTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{6}
}

func (x *AddTemplateRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *AddTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddTemplateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 修改模板
type UseTemplateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId int64 `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *UseTemplateVersionRequest) Reset() {
	*x = UseTemplateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseTemplateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseTemplateVersionRequest) ProtoMessage() {}

func (x *UseTemplateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseTemplateVersionRequest.ProtoReflect.Descriptor instead.
func (*UseTemplateVersionRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{7}
}

func (x *UseTemplateVersionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UseTemplateVersionRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

// 解析模板
type ParseTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrvKeyword string `protobuf:"bytes,1,opt,name=srv_keyword,json=srvKeyword,proto3" json:"srv_keyword,omitempty"`
	EnvKeyword string `protobuf:"bytes,2,opt,name=env_keyword,json=envKeyword,proto3" json:"env_keyword,omitempty"`
}

func (x *ParseTemplateRequest) Reset() {
	*x = ParseTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTemplateRequest) ProtoMessage() {}

func (x *ParseTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTemplateRequest.ProtoReflect.Descriptor instead.
func (*ParseTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{8}
}

func (x *ParseTemplateRequest) GetSrvKeyword() string {
	if x != nil {
		return x.SrvKeyword
	}
	return ""
}

func (x *ParseTemplateRequest) GetEnvKeyword() string {
	if x != nil {
		return x.EnvKeyword
	}
	return ""
}

type ParseTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ParseTemplateReply) Reset() {
	*x = ParseTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTemplateReply) ProtoMessage() {}

func (x *ParseTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTemplateReply.ProtoReflect.Descriptor instead.
func (*ParseTemplateReply) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{9}
}

func (x *ParseTemplateReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PageTemplateReply_Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId    int64   `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Version     string  `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	IsUse       bool    `protobuf:"varint,5,opt,name=is_use,json=isUse,proto3" json:"is_use,omitempty"`
	CreatedAt   int64   `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Operator    *string `protobuf:"bytes,7,opt,name=operator,proto3,oneof" json:"operator,omitempty"`
	OperatorId  *int64  `protobuf:"varint,8,opt,name=operator_id,json=operatorId,proto3,oneof" json:"operator_id,omitempty"`
}

func (x *PageTemplateReply_Template) Reset() {
	*x = PageTemplateReply_Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTemplateReply_Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTemplateReply_Template) ProtoMessage() {}

func (x *PageTemplateReply_Template) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTemplateReply_Template.ProtoReflect.Descriptor instead.
func (*PageTemplateReply_Template) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PageTemplateReply_Template) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PageTemplateReply_Template) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *PageTemplateReply_Template) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PageTemplateReply_Template) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PageTemplateReply_Template) GetIsUse() bool {
	if x != nil {
		return x.IsUse
	}
	return false
}

func (x *PageTemplateReply_Template) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PageTemplateReply_Template) GetOperator() string {
	if x != nil && x.Operator != nil {
		return *x.Operator
	}
	return ""
}

func (x *PageTemplateReply_Template) GetOperatorId() int64 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_template_proto protoreflect.FileDescriptor

var file_template_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01,
	0x0a, 0x13, 0x50, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x32, 0x20, 0x00,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xed, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x1a, 0x8d, 0x02, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x55, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x22, 0x3e, 0x0a, 0x16, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x94, 0x02, 0x0a, 0x14, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x55, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa0, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x06, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x05, 0x69, 0x73, 0x55, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x40, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x6a, 0x0a, 0x14, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x72,
	0x76, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x72, 0x76, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x5f, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0a, 0x65, 0x6e, 0x76, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2e,
	0x0a, 0x12, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_template_proto_rawDescOnce sync.Once
	file_template_proto_rawDescData = file_template_proto_rawDesc
)

func file_template_proto_rawDescGZIP() []byte {
	file_template_proto_rawDescOnce.Do(func() {
		file_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_proto_rawDescData)
	})
	return file_template_proto_rawDescData
}

var file_template_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_template_proto_goTypes = []interface{}{
	(*PageTemplateRequest)(nil),        // 0: v1.PageTemplateRequest
	(*PageTemplateReply)(nil),          // 1: v1.PageTemplateReply
	(*CurrentTemplateRequest)(nil),     // 2: v1.CurrentTemplateRequest
	(*CurrentTemplateReply)(nil),       // 3: v1.CurrentTemplateReply
	(*GetTemplateRequest)(nil),         // 4: v1.GetTemplateRequest
	(*GetTemplateReply)(nil),           // 5: v1.GetTemplateReply
	(*AddTemplateRequest)(nil),         // 6: v1.AddTemplateRequest
	(*UseTemplateVersionRequest)(nil),  // 7: v1.UseTemplateVersionRequest
	(*ParseTemplateRequest)(nil),       // 8: v1.ParseTemplateRequest
	(*ParseTemplateReply)(nil),         // 9: v1.ParseTemplateReply
	(*PageTemplateReply_Template)(nil), // 10: v1.PageTemplateReply.Template
}
var file_template_proto_depIdxs = []int32{
	10, // 0: v1.PageTemplateReply.list:type_name -> v1.PageTemplateReply.Template
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_template_proto_init() }
func file_template_proto_init() {
	if File_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTemplateRequest); i {
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
		file_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTemplateReply); i {
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
		file_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentTemplateRequest); i {
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
		file_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentTemplateReply); i {
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
		file_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateRequest); i {
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
		file_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateReply); i {
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
		file_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTemplateRequest); i {
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
		file_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseTemplateVersionRequest); i {
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
		file_template_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTemplateRequest); i {
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
		file_template_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTemplateReply); i {
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
		file_template_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTemplateReply_Template); i {
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
	file_template_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_template_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_template_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_proto_goTypes,
		DependencyIndexes: file_template_proto_depIdxs,
		MessageInfos:      file_template_proto_msgTypes,
	}.Build()
	File_template_proto = out.File
	file_template_proto_rawDesc = nil
	file_template_proto_goTypes = nil
	file_template_proto_depIdxs = nil
}