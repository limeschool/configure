// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resource_value.proto

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

type AllResourceValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId int64 `protobuf:"varint,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *AllResourceValueRequest) Reset() {
	*x = AllResourceValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllResourceValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllResourceValueRequest) ProtoMessage() {}

func (x *AllResourceValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllResourceValueRequest.ProtoReflect.Descriptor instead.
func (*AllResourceValueRequest) Descriptor() ([]byte, []int) {
	return file_resource_value_proto_rawDescGZIP(), []int{0}
}

func (x *AllResourceValueRequest) GetResourceId() int64 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

type AllResourceValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AllResourceValueReply_ResourceValue `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AllResourceValueReply) Reset() {
	*x = AllResourceValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllResourceValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllResourceValueReply) ProtoMessage() {}

func (x *AllResourceValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_resource_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllResourceValueReply.ProtoReflect.Descriptor instead.
func (*AllResourceValueReply) Descriptor() ([]byte, []int) {
	return file_resource_value_proto_rawDescGZIP(), []int{1}
}

func (x *AllResourceValueReply) GetList() []*AllResourceValueReply_ResourceValue {
	if x != nil {
		return x.List
	}
	return nil
}

// 新增资源值
type AddResourceValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId int64  `protobuf:"varint,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ResourceId    int64  `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Values        string `protobuf:"bytes,3,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *AddResourceValueRequest) Reset() {
	*x = AddResourceValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResourceValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResourceValueRequest) ProtoMessage() {}

func (x *AddResourceValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResourceValueRequest.ProtoReflect.Descriptor instead.
func (*AddResourceValueRequest) Descriptor() ([]byte, []int) {
	return file_resource_value_proto_rawDescGZIP(), []int{2}
}

func (x *AddResourceValueRequest) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *AddResourceValueRequest) GetResourceId() int64 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *AddResourceValueRequest) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

// 修改资源值
type UpdateResourceValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Values string `protobuf:"bytes,4,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *UpdateResourceValueRequest) Reset() {
	*x = UpdateResourceValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_value_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceValueRequest) ProtoMessage() {}

func (x *UpdateResourceValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_value_proto_msgTypes[3]
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
	return file_resource_value_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResourceValueRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateResourceValueRequest) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

type AllResourceValueReply_ResourceValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvironmentId int64        `protobuf:"varint,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ResourceId    int64        `protobuf:"varint,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Values        string       `protobuf:"bytes,4,opt,name=values,proto3" json:"values,omitempty"`
	Environment   *Environment `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *AllResourceValueReply_ResourceValue) Reset() {
	*x = AllResourceValueReply_ResourceValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_value_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllResourceValueReply_ResourceValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllResourceValueReply_ResourceValue) ProtoMessage() {}

func (x *AllResourceValueReply_ResourceValue) ProtoReflect() protoreflect.Message {
	mi := &file_resource_value_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllResourceValueReply_ResourceValue.ProtoReflect.Descriptor instead.
func (*AllResourceValueReply_ResourceValue) Descriptor() ([]byte, []int) {
	return file_resource_value_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AllResourceValueReply_ResourceValue) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AllResourceValueReply_ResourceValue) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *AllResourceValueReply_ResourceValue) GetResourceId() int64 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *AllResourceValueReply_ResourceValue) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

func (x *AllResourceValueReply_ResourceValue) GetEnvironment() *Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

var File_resource_value_proto protoreflect.FileDescriptor

var file_resource_value_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x17, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x89, 0x02, 0x0a, 0x15,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x1a, 0xb2, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x07, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_value_proto_rawDescOnce sync.Once
	file_resource_value_proto_rawDescData = file_resource_value_proto_rawDesc
)

func file_resource_value_proto_rawDescGZIP() []byte {
	file_resource_value_proto_rawDescOnce.Do(func() {
		file_resource_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_value_proto_rawDescData)
	})
	return file_resource_value_proto_rawDescData
}

var file_resource_value_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_resource_value_proto_goTypes = []interface{}{
	(*AllResourceValueRequest)(nil),             // 0: v1.AllResourceValueRequest
	(*AllResourceValueReply)(nil),               // 1: v1.AllResourceValueReply
	(*AddResourceValueRequest)(nil),             // 2: v1.AddResourceValueRequest
	(*UpdateResourceValueRequest)(nil),          // 3: v1.UpdateResourceValueRequest
	(*AllResourceValueReply_ResourceValue)(nil), // 4: v1.AllResourceValueReply.ResourceValue
	(*Environment)(nil),                         // 5: v1.Environment
}
var file_resource_value_proto_depIdxs = []int32{
	4, // 0: v1.AllResourceValueReply.list:type_name -> v1.AllResourceValueReply.ResourceValue
	5, // 1: v1.AllResourceValueReply.ResourceValue.environment:type_name -> v1.Environment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resource_value_proto_init() }
func file_resource_value_proto_init() {
	if File_resource_value_proto != nil {
		return
	}
	file_resource_proto_init()
	file_environment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_resource_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllResourceValueRequest); i {
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
		file_resource_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllResourceValueReply); i {
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
		file_resource_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResourceValueRequest); i {
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
		file_resource_value_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_resource_value_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllResourceValueReply_ResourceValue); i {
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
			RawDescriptor: file_resource_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_value_proto_goTypes,
		DependencyIndexes: file_resource_value_proto_depIdxs,
		MessageInfos:      file_resource_value_proto_msgTypes,
	}.Build()
	File_resource_value_proto = out.File
	file_resource_value_proto_rawDesc = nil
	file_resource_value_proto_goTypes = nil
	file_resource_value_proto_depIdxs = nil
}