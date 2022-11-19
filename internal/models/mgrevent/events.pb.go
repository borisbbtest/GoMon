// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: events.proto

package mgrevent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Source      string                 `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Status      int32                  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Update      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update,proto3" json:"update,omitempty"`
	Key         string                 `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	KeyClose    string                 `protobuf:"bytes,8,opt,name=key_close,json=keyClose,proto3" json:"key_close,omitempty"`
	Assigned    []string               `protobuf:"bytes,9,rep,name=assigned,proto3" json:"assigned,omitempty"`
	AutoRunner  string                 `protobuf:"bytes,10,opt,name=auto_runner,json=autoRunner,proto3" json:"auto_runner,omitempty"`
	Severity    int32                  `protobuf:"varint,11,opt,name=severity,proto3" json:"severity,omitempty"`
	RelarionCi  []string               `protobuf:"bytes,12,rep,name=relarion_ci,json=relarionCi,proto3" json:"relarion_ci,omitempty"`
	CreatedBy   string                 `protobuf:"bytes,13,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Count       int32                  `protobuf:"varint,14,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Event) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Event) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Event) GetUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Event) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Event) GetKeyClose() string {
	if x != nil {
		return x.KeyClose
	}
	return ""
}

func (x *Event) GetAssigned() []string {
	if x != nil {
		return x.Assigned
	}
	return nil
}

func (x *Event) GetAutoRunner() string {
	if x != nil {
		return x.AutoRunner
	}
	return ""
}

func (x *Event) GetSeverity() int32 {
	if x != nil {
		return x.Severity
	}
	return 0
}

func (x *Event) GetRelarionCi() []string {
	if x != nil {
		return x.RelarionCi
	}
	return nil
}

func (x *Event) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Event) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Ev   *Event `protobuf:"bytes,2,opt,name=ev,proto3" json:"ev,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *PushRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PushRequest) GetEv() *Event {
	if x != nil {
		return x.Ev
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *PushResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Ev   *Event `protobuf:"bytes,2,opt,name=ev,proto3" json:"ev,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetResponse) GetEv() *Event {
	if x != nil {
		return x.Ev
	}
	return nil
}

type PushBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Ev   []*Event `protobuf:"bytes,2,rep,name=ev,proto3" json:"ev,omitempty"`
}

func (x *PushBatchRequest) Reset() {
	*x = PushBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBatchRequest) ProtoMessage() {}

func (x *PushBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBatchRequest.ProtoReflect.Descriptor instead.
func (*PushBatchRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{5}
}

func (x *PushBatchRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PushBatchRequest) GetEv() []*Event {
	if x != nil {
		return x.Ev
	}
	return nil
}

type PushBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PushBatchResponse) Reset() {
	*x = PushBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBatchResponse) ProtoMessage() {}

func (x *PushBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBatchResponse.ProtoReflect.Descriptor instead.
func (*PushBatchResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{6}
}

func (x *PushBatchResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ev []*Event `protobuf:"bytes,1,rep,name=ev,proto3" json:"ev,omitempty"`
}

func (x *GetBatchRequest) Reset() {
	*x = GetBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchRequest) ProtoMessage() {}

func (x *GetBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchRequest.ProtoReflect.Descriptor instead.
func (*GetBatchRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{7}
}

func (x *GetBatchRequest) GetEv() []*Event {
	if x != nil {
		return x.Ev
	}
	return nil
}

type GetBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Ev   []*Event `protobuf:"bytes,2,rep,name=ev,proto3" json:"ev,omitempty"`
}

func (x *GetBatchResponse) Reset() {
	*x = GetBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchResponse) ProtoMessage() {}

func (x *GetBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchResponse.ProtoReflect.Descriptor instead.
func (*GetBatchResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{8}
}

func (x *GetBatchResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetBatchResponse) GetEv() []*Event {
	if x != nil {
		return x.Ev
	}
	return nil
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79,
	0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x72, 0x69, 0x6f, 0x6e, 0x43, 0x69,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x02, 0x65, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x02, 0x65, 0x76, 0x22, 0x22, 0x0a, 0x0c, 0x50,
	0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x02, 0x65, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x02, 0x65, 0x76, 0x22, 0x49, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x02, 0x65,
	0x76, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x02, 0x65, 0x76, 0x22, 0x27,
	0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x02, 0x65, 0x76,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x02, 0x65, 0x76, 0x22, 0x49, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x02, 0x65, 0x76, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x02, 0x65, 0x76, 0x32, 0x94, 0x02, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x17, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x50, 0x75,
	0x73, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x6d, 0x67, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_events_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: eventsgrpc.Event
	(*PushRequest)(nil),           // 1: eventsgrpc.PushRequest
	(*PushResponse)(nil),          // 2: eventsgrpc.PushResponse
	(*GetRequest)(nil),            // 3: eventsgrpc.GetRequest
	(*GetResponse)(nil),           // 4: eventsgrpc.GetResponse
	(*PushBatchRequest)(nil),      // 5: eventsgrpc.PushBatchRequest
	(*PushBatchResponse)(nil),     // 6: eventsgrpc.PushBatchResponse
	(*GetBatchRequest)(nil),       // 7: eventsgrpc.GetBatchRequest
	(*GetBatchResponse)(nil),      // 8: eventsgrpc.GetBatchResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_events_proto_depIdxs = []int32{
	9,  // 0: eventsgrpc.Event.created:type_name -> google.protobuf.Timestamp
	9,  // 1: eventsgrpc.Event.update:type_name -> google.protobuf.Timestamp
	0,  // 2: eventsgrpc.PushRequest.ev:type_name -> eventsgrpc.Event
	0,  // 3: eventsgrpc.GetResponse.ev:type_name -> eventsgrpc.Event
	0,  // 4: eventsgrpc.PushBatchRequest.ev:type_name -> eventsgrpc.Event
	0,  // 5: eventsgrpc.GetBatchRequest.ev:type_name -> eventsgrpc.Event
	0,  // 6: eventsgrpc.GetBatchResponse.ev:type_name -> eventsgrpc.Event
	1,  // 7: eventsgrpc.events.Push:input_type -> eventsgrpc.PushRequest
	3,  // 8: eventsgrpc.events.Get:input_type -> eventsgrpc.GetRequest
	5,  // 9: eventsgrpc.events.PushBatch:input_type -> eventsgrpc.PushBatchRequest
	7,  // 10: eventsgrpc.events.GetBatch:input_type -> eventsgrpc.GetBatchRequest
	2,  // 11: eventsgrpc.events.Push:output_type -> eventsgrpc.PushResponse
	4,  // 12: eventsgrpc.events.Get:output_type -> eventsgrpc.GetResponse
	6,  // 13: eventsgrpc.events.PushBatch:output_type -> eventsgrpc.PushBatchResponse
	8,  // 14: eventsgrpc.events.GetBatch:output_type -> eventsgrpc.GetBatchResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBatchRequest); i {
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
		file_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBatchResponse); i {
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
		file_events_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchRequest); i {
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
		file_events_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchResponse); i {
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
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}
