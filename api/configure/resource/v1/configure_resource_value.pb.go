// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/configure/resource/configure_resource_value.proto

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

type ListResourceValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId uint32 `protobuf:"varint,3,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *ListResourceValueRequest) Reset() {
	*x = ListResourceValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceValueRequest) ProtoMessage() {}

func (x *ListResourceValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceValueRequest.ProtoReflect.Descriptor instead.
func (*ListResourceValueRequest) Descriptor() ([]byte, []int) {
	return file_api_configure_resource_configure_resource_value_proto_rawDescGZIP(), []int{0}
}

func (x *ListResourceValueRequest) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

type ListResourceValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32                                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*ListResourceValueReply_ResourceValue `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListResourceValueReply) Reset() {
	*x = ListResourceValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceValueReply) ProtoMessage() {}

func (x *ListResourceValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceValueReply.ProtoReflect.Descriptor instead.
func (*ListResourceValueReply) Descriptor() ([]byte, []int) {
	return file_api_configure_resource_configure_resource_value_proto_rawDescGZIP(), []int{1}
}

func (x *ListResourceValueReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResourceValueReply) GetList() []*ListResourceValueReply_ResourceValue {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateResourceValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List       []*UpdateResourceValueRequest_Value `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	ResourceId uint32                              `protobuf:"varint,2,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *UpdateResourceValueRequest) Reset() {
	*x = UpdateResourceValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceValueRequest) ProtoMessage() {}

func (x *UpdateResourceValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceValueRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceValueRequest) Descriptor() ([]byte, []int) {
	return file_api_configure_resource_configure_resource_value_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateResourceValueRequest) GetList() []*UpdateResourceValueRequest_Value {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UpdateResourceValueRequest) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

type UpdateResourceValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResourceValueReply) Reset() {
	*x = UpdateResourceValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceValueReply) ProtoMessage() {}

func (x *UpdateResourceValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceValueReply.ProtoReflect.Descriptor instead.
func (*UpdateResourceValueReply) Descriptor() ([]byte, []int) {
	return file_api_configure_resource_configure_resource_value_proto_rawDescGZIP(), []int{3}
}

type ListResourceValueReply_ResourceValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvId      uint32 `protobuf:"varint,2,opt,name=envId,proto3" json:"envId,omitempty"`
	ResourceId uint32 `protobuf:"varint,3,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Value      string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	CreatedAt  uint32 `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt  uint32 `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ListResourceValueReply_ResourceValue) Reset() {
	*x = ListResourceValueReply_ResourceValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceValueReply_ResourceValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceValueReply_ResourceValue) ProtoMessage() {}

func (x *ListResourceValueReply_ResourceValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceValueReply_ResourceValue.ProtoReflect.Descriptor instead.
func (*ListResourceValueReply_ResourceValue) Descriptor() ([]byte, []int) {
	return file_api_configure_resource_configure_resource_value_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListResourceValueReply_ResourceValue) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListResourceValueReply_ResourceValue) GetEnvId() uint32 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *ListResourceValueReply_ResourceValue) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *ListResourceValueReply_ResourceValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ListResourceValueReply_ResourceValue) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListResourceValueReply_ResourceValue) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UpdateResourceValueRequest_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvId uint32 `protobuf:"varint,2,opt,name=envId,proto3" json:"envId,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateResourceValueRequest_Value) Reset() {
	*x = UpdateResourceValueRequest_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceValueRequest_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceValueRequest_Value) ProtoMessage() {}

func (x *UpdateResourceValueRequest_Value) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_resource_configure_resource_value_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceValueRequest_Value.ProtoReflect.Descriptor instead.
func (*UpdateResourceValueRequest_Value) Descriptor() ([]byte, []int) {
	return file_api_configure_resource_configure_resource_value_proto_rawDescGZIP(), []int{2, 0}
}

func (x *UpdateResourceValueRequest_Value) GetEnvId() uint32 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *UpdateResourceValueRequest_Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_api_configure_resource_configure_resource_value_proto protoreflect.FileDescriptor

var file_api_configure_resource_configure_resource_value_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xb7, 0x02, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x5d, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0xa7, 0x01, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6e, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x6e, 0x76,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xf1, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x1a, 0x45, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6e,
	0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x20, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x3c, 0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_configure_resource_configure_resource_value_proto_rawDescOnce sync.Once
	file_api_configure_resource_configure_resource_value_proto_rawDescData = file_api_configure_resource_configure_resource_value_proto_rawDesc
)

func file_api_configure_resource_configure_resource_value_proto_rawDescGZIP() []byte {
	file_api_configure_resource_configure_resource_value_proto_rawDescOnce.Do(func() {
		file_api_configure_resource_configure_resource_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_configure_resource_configure_resource_value_proto_rawDescData)
	})
	return file_api_configure_resource_configure_resource_value_proto_rawDescData
}

var file_api_configure_resource_configure_resource_value_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_configure_resource_configure_resource_value_proto_goTypes = []interface{}{
	(*ListResourceValueRequest)(nil),             // 0: configure.api.configure.resource.v1.ListResourceValueRequest
	(*ListResourceValueReply)(nil),               // 1: configure.api.configure.resource.v1.ListResourceValueReply
	(*UpdateResourceValueRequest)(nil),           // 2: configure.api.configure.resource.v1.UpdateResourceValueRequest
	(*UpdateResourceValueReply)(nil),             // 3: configure.api.configure.resource.v1.UpdateResourceValueReply
	(*ListResourceValueReply_ResourceValue)(nil), // 4: configure.api.configure.resource.v1.ListResourceValueReply.ResourceValue
	(*UpdateResourceValueRequest_Value)(nil),     // 5: configure.api.configure.resource.v1.UpdateResourceValueRequest.Value
}
var file_api_configure_resource_configure_resource_value_proto_depIdxs = []int32{
	4, // 0: configure.api.configure.resource.v1.ListResourceValueReply.list:type_name -> configure.api.configure.resource.v1.ListResourceValueReply.ResourceValue
	5, // 1: configure.api.configure.resource.v1.UpdateResourceValueRequest.list:type_name -> configure.api.configure.resource.v1.UpdateResourceValueRequest.Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_configure_resource_configure_resource_value_proto_init() }
func file_api_configure_resource_configure_resource_value_proto_init() {
	if File_api_configure_resource_configure_resource_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_configure_resource_configure_resource_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceValueRequest); i {
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
		file_api_configure_resource_configure_resource_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceValueReply); i {
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
		file_api_configure_resource_configure_resource_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceValueRequest); i {
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
		file_api_configure_resource_configure_resource_value_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceValueReply); i {
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
		file_api_configure_resource_configure_resource_value_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceValueReply_ResourceValue); i {
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
		file_api_configure_resource_configure_resource_value_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceValueRequest_Value); i {
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
			RawDescriptor: file_api_configure_resource_configure_resource_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_configure_resource_configure_resource_value_proto_goTypes,
		DependencyIndexes: file_api_configure_resource_configure_resource_value_proto_depIdxs,
		MessageInfos:      file_api_configure_resource_configure_resource_value_proto_msgTypes,
	}.Build()
	File_api_configure_resource_configure_resource_value_proto = out.File
	file_api_configure_resource_configure_resource_value_proto_rawDesc = nil
	file_api_configure_resource_configure_resource_value_proto_goTypes = nil
	file_api_configure_resource_configure_resource_value_proto_depIdxs = nil
}