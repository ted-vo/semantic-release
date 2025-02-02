// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/publisher/publisher.proto

package publisher

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

type PublisherInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublisherInit) Reset() {
	*x = PublisherInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherInit) ProtoMessage() {}

func (x *PublisherInit) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherInit.ProtoReflect.Descriptor instead.
func (*PublisherInit) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{0}
}

type PublisherName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublisherName) Reset() {
	*x = PublisherName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherName) ProtoMessage() {}

func (x *PublisherName) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherName.ProtoReflect.Descriptor instead.
func (*PublisherName) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{1}
}

type PublisherVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublisherVersion) Reset() {
	*x = PublisherVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherVersion) ProtoMessage() {}

func (x *PublisherVersion) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherVersion.ProtoReflect.Descriptor instead.
func (*PublisherVersion) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{2}
}

type PublisherPublish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublisherPublish) Reset() {
	*x = PublisherPublish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherPublish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherPublish) ProtoMessage() {}

func (x *PublisherPublish) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherPublish.ProtoReflect.Descriptor instead.
func (*PublisherPublish) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{3}
}

type PublisherInit_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config map[string]string `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PublisherInit_Request) Reset() {
	*x = PublisherInit_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherInit_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherInit_Request) ProtoMessage() {}

func (x *PublisherInit_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherInit_Request.ProtoReflect.Descriptor instead.
func (*PublisherInit_Request) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PublisherInit_Request) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type PublisherInit_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PublisherInit_Response) Reset() {
	*x = PublisherInit_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherInit_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherInit_Response) ProtoMessage() {}

func (x *PublisherInit_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherInit_Response.ProtoReflect.Descriptor instead.
func (*PublisherInit_Response) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PublisherInit_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type PublisherName_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublisherName_Request) Reset() {
	*x = PublisherName_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherName_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherName_Request) ProtoMessage() {}

func (x *PublisherName_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherName_Request.ProtoReflect.Descriptor instead.
func (*PublisherName_Request) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{1, 0}
}

type PublisherName_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PublisherName_Response) Reset() {
	*x = PublisherName_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherName_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherName_Response) ProtoMessage() {}

func (x *PublisherName_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherName_Response.ProtoReflect.Descriptor instead.
func (*PublisherName_Response) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PublisherName_Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PublisherVersion_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublisherVersion_Request) Reset() {
	*x = PublisherVersion_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherVersion_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherVersion_Request) ProtoMessage() {}

func (x *PublisherVersion_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherVersion_Request.ProtoReflect.Descriptor instead.
func (*PublisherVersion_Request) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{2, 0}
}

type PublisherVersion_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PublisherVersion_Response) Reset() {
	*x = PublisherVersion_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherVersion_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherVersion_Response) ProtoMessage() {}

func (x *PublisherVersion_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherVersion_Response.ProtoReflect.Descriptor instead.
func (*PublisherVersion_Response) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{2, 1}
}

func (x *PublisherVersion_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type PublisherPublish_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewRelease string `protobuf:"bytes,1,opt,name=new_release,json=newRelease,proto3" json:"new_release,omitempty"`
}

func (x *PublisherPublish_Request) Reset() {
	*x = PublisherPublish_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherPublish_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherPublish_Request) ProtoMessage() {}

func (x *PublisherPublish_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherPublish_Request.ProtoReflect.Descriptor instead.
func (*PublisherPublish_Request) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{3, 0}
}

func (x *PublisherPublish_Request) GetNewRelease() string {
	if x != nil {
		return x.NewRelease
	}
	return ""
}

type PublisherPublish_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PublisherPublish_Response) Reset() {
	*x = PublisherPublish_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_publisher_publisher_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherPublish_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherPublish_Response) ProtoMessage() {}

func (x *PublisherPublish_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_publisher_publisher_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherPublish_Response.ProtoReflect.Descriptor instead.
func (*PublisherPublish_Response) Descriptor() ([]byte, []int) {
	return file_pkg_publisher_publisher_proto_rawDescGZIP(), []int{3, 1}
}

func (x *PublisherPublish_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_publisher_publisher_proto protoreflect.FileDescriptor

var file_pkg_publisher_publisher_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb4, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x69,
	0x74, 0x1a, 0x80, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3a, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x43, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x1a, 0x2a, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x87, 0x02, 0x0a, 0x0f, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x37, 0x0a,
	0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x64, 0x2d, 0x76, 0x6f, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_publisher_publisher_proto_rawDescOnce sync.Once
	file_pkg_publisher_publisher_proto_rawDescData = file_pkg_publisher_publisher_proto_rawDesc
)

func file_pkg_publisher_publisher_proto_rawDescGZIP() []byte {
	file_pkg_publisher_publisher_proto_rawDescOnce.Do(func() {
		file_pkg_publisher_publisher_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_publisher_publisher_proto_rawDescData)
	})
	return file_pkg_publisher_publisher_proto_rawDescData
}

var file_pkg_publisher_publisher_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_publisher_publisher_proto_goTypes = []interface{}{
	(*PublisherInit)(nil),             // 0: PublisherInit
	(*PublisherName)(nil),             // 1: PublisherName
	(*PublisherVersion)(nil),          // 2: PublisherVersion
	(*PublisherPublish)(nil),          // 3: PublisherPublish
	(*PublisherInit_Request)(nil),     // 4: PublisherInit.Request
	(*PublisherInit_Response)(nil),    // 5: PublisherInit.Response
	nil,                               // 6: PublisherInit.Request.ConfigEntry
	(*PublisherName_Request)(nil),     // 7: PublisherName.Request
	(*PublisherName_Response)(nil),    // 8: PublisherName.Response
	(*PublisherVersion_Request)(nil),  // 9: PublisherVersion.Request
	(*PublisherVersion_Response)(nil), // 10: PublisherVersion.Response
	(*PublisherPublish_Request)(nil),  // 11: PublisherPublish.Request
	(*PublisherPublish_Response)(nil), // 12: PublisherPublish.Response
}
var file_pkg_publisher_publisher_proto_depIdxs = []int32{
	6,  // 0: PublisherInit.Request.config:type_name -> PublisherInit.Request.ConfigEntry
	4,  // 1: PublisherPlugin.Init:input_type -> PublisherInit.Request
	7,  // 2: PublisherPlugin.Name:input_type -> PublisherName.Request
	9,  // 3: PublisherPlugin.Version:input_type -> PublisherVersion.Request
	11, // 4: PublisherPlugin.Publish:input_type -> PublisherPublish.Request
	5,  // 5: PublisherPlugin.Init:output_type -> PublisherInit.Response
	8,  // 6: PublisherPlugin.Name:output_type -> PublisherName.Response
	10, // 7: PublisherPlugin.Version:output_type -> PublisherVersion.Response
	12, // 8: PublisherPlugin.Publish:output_type -> PublisherPublish.Response
	5,  // [5:9] is the sub-list for method output_type
	1,  // [1:5] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_publisher_publisher_proto_init() }
func file_pkg_publisher_publisher_proto_init() {
	if File_pkg_publisher_publisher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_publisher_publisher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherInit); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherName); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherVersion); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherPublish); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherInit_Request); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherInit_Response); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherName_Request); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherName_Response); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherVersion_Request); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherVersion_Response); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherPublish_Request); i {
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
		file_pkg_publisher_publisher_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherPublish_Response); i {
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
			RawDescriptor: file_pkg_publisher_publisher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_publisher_publisher_proto_goTypes,
		DependencyIndexes: file_pkg_publisher_publisher_proto_depIdxs,
		MessageInfos:      file_pkg_publisher_publisher_proto_msgTypes,
	}.Build()
	File_pkg_publisher_publisher_proto = out.File
	file_pkg_publisher_publisher_proto_rawDesc = nil
	file_pkg_publisher_publisher_proto_goTypes = nil
	file_pkg_publisher_publisher_proto_depIdxs = nil
}
