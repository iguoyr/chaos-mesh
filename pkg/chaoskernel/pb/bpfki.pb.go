// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: bpfki.proto

package bpfki

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FailKernRequest_FAILTYPE int32

const (
	FailKernRequest_SLAB FailKernRequest_FAILTYPE = 0
	FailKernRequest_PAGE FailKernRequest_FAILTYPE = 1
	FailKernRequest_BIO  FailKernRequest_FAILTYPE = 2
)

// Enum value maps for FailKernRequest_FAILTYPE.
var (
	FailKernRequest_FAILTYPE_name = map[int32]string{
		0: "SLAB",
		1: "PAGE",
		2: "BIO",
	}
	FailKernRequest_FAILTYPE_value = map[string]int32{
		"SLAB": 0,
		"PAGE": 1,
		"BIO":  2,
	}
)

func (x FailKernRequest_FAILTYPE) Enum() *FailKernRequest_FAILTYPE {
	p := new(FailKernRequest_FAILTYPE)
	*p = x
	return p
}

func (x FailKernRequest_FAILTYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailKernRequest_FAILTYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_bpfki_proto_enumTypes[0].Descriptor()
}

func (FailKernRequest_FAILTYPE) Type() protoreflect.EnumType {
	return &file_bpfki_proto_enumTypes[0]
}

func (x FailKernRequest_FAILTYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailKernRequest_FAILTYPE.Descriptor instead.
func (FailKernRequest_FAILTYPE) EnumDescriptor() ([]byte, []int) {
	return file_bpfki_proto_rawDescGZIP(), []int{1, 0}
}

type BumpTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid         uint32  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Tid         uint32  `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	Second      int32   `protobuf:"varint,3,opt,name=second,proto3" json:"second,omitempty"`
	Subsecond   int32   `protobuf:"varint,4,opt,name=subsecond,proto3" json:"subsecond,omitempty"`
	Probability float32 `protobuf:"fixed32,5,opt,name=probability,proto3" json:"probability,omitempty"`
}

func (x *BumpTimeRequest) Reset() {
	*x = BumpTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfki_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BumpTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BumpTimeRequest) ProtoMessage() {}

func (x *BumpTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfki_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BumpTimeRequest.ProtoReflect.Descriptor instead.
func (*BumpTimeRequest) Descriptor() ([]byte, []int) {
	return file_bpfki_proto_rawDescGZIP(), []int{0}
}

func (x *BumpTimeRequest) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *BumpTimeRequest) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *BumpTimeRequest) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

func (x *BumpTimeRequest) GetSubsecond() int32 {
	if x != nil {
		return x.Subsecond
	}
	return 0
}

func (x *BumpTimeRequest) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

type FailKernRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid         uint32                   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Tid         uint32                   `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	Ftype       FailKernRequest_FAILTYPE `protobuf:"varint,3,opt,name=ftype,proto3,enum=bpfki.FailKernRequest_FAILTYPE" json:"ftype,omitempty"`
	Headers     []string                 `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	Callchain   []*FailKernRequestFrame  `protobuf:"bytes,5,rep,name=callchain,proto3" json:"callchain,omitempty"`
	Probability float32                  `protobuf:"fixed32,6,opt,name=probability,proto3" json:"probability,omitempty"`
	Times       uint32                   `protobuf:"varint,7,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *FailKernRequest) Reset() {
	*x = FailKernRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfki_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailKernRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailKernRequest) ProtoMessage() {}

func (x *FailKernRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfki_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailKernRequest.ProtoReflect.Descriptor instead.
func (*FailKernRequest) Descriptor() ([]byte, []int) {
	return file_bpfki_proto_rawDescGZIP(), []int{1}
}

func (x *FailKernRequest) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *FailKernRequest) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *FailKernRequest) GetFtype() FailKernRequest_FAILTYPE {
	if x != nil {
		return x.Ftype
	}
	return FailKernRequest_SLAB
}

func (x *FailKernRequest) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *FailKernRequest) GetCallchain() []*FailKernRequestFrame {
	if x != nil {
		return x.Callchain
	}
	return nil
}

func (x *FailKernRequest) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

func (x *FailKernRequest) GetTimes() uint32 {
	if x != nil {
		return x.Times
	}
	return 0
}

type FailSyscallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid         uint32   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Tid         uint32   `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	Methods     []string `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
	Err         uint32   `protobuf:"varint,4,opt,name=err,proto3" json:"err,omitempty"`
	Probability float32  `protobuf:"fixed32,5,opt,name=probability,proto3" json:"probability,omitempty"`
}

func (x *FailSyscallRequest) Reset() {
	*x = FailSyscallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfki_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailSyscallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailSyscallRequest) ProtoMessage() {}

func (x *FailSyscallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfki_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailSyscallRequest.ProtoReflect.Descriptor instead.
func (*FailSyscallRequest) Descriptor() ([]byte, []int) {
	return file_bpfki_proto_rawDescGZIP(), []int{2}
}

func (x *FailSyscallRequest) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *FailSyscallRequest) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *FailSyscallRequest) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *FailSyscallRequest) GetErr() uint32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *FailSyscallRequest) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfki_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfki_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_bpfki_proto_rawDescGZIP(), []int{3}
}

func (x *StatusResponse) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *StatusResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type FailKernRequestFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Funcname   string `protobuf:"bytes,1,opt,name=funcname,proto3" json:"funcname,omitempty"`
	Parameters string `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Predicate  string `protobuf:"bytes,3,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *FailKernRequestFrame) Reset() {
	*x = FailKernRequestFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfki_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailKernRequestFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailKernRequestFrame) ProtoMessage() {}

func (x *FailKernRequestFrame) ProtoReflect() protoreflect.Message {
	mi := &file_bpfki_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailKernRequestFrame.ProtoReflect.Descriptor instead.
func (*FailKernRequestFrame) Descriptor() ([]byte, []int) {
	return file_bpfki_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FailKernRequestFrame) GetFuncname() string {
	if x != nil {
		return x.Funcname
	}
	return ""
}

func (x *FailKernRequestFrame) GetParameters() string {
	if x != nil {
		return x.Parameters
	}
	return ""
}

func (x *FailKernRequestFrame) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

var File_bpfki_proto protoreflect.FileDescriptor

var file_bpfki_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x70, 0x66, 0x6b, 0x69, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x42, 0x75, 0x6d, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x86, 0x03, 0x0a, 0x0f, 0x46, 0x61, 0x69, 0x6c, 0x4b, 0x65, 0x72,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x05,
	0x66, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x70,
	0x66, 0x6b, 0x69, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x4b, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x46, 0x41, 0x49, 0x4c, 0x54, 0x59, 0x50, 0x45, 0x52, 0x05, 0x66, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a,
	0x09, 0x63, 0x61, 0x6c, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x4b, 0x65, 0x72,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x09,
	0x63, 0x61, 0x6c, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x1a, 0x61, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x08, 0x46, 0x41, 0x49, 0x4c, 0x54, 0x59, 0x50, 0x45,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x4c, 0x41, 0x42, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41,
	0x47, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x4f, 0x10, 0x02, 0x22, 0x86, 0x01,
	0x0a, 0x12, 0x46, 0x61, 0x69, 0x6c, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x34, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa0, 0x04, 0x0a,
	0x0c, 0x42, 0x50, 0x46, 0x4b, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x62, 0x70,
	0x66, 0x6b, 0x69, 0x2e, 0x42, 0x75, 0x6d, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x16,
	0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x42, 0x75, 0x6d, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16,
	0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x42, 0x75, 0x6d, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x16, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x42, 0x75, 0x6d, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66,
	0x6b, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x46, 0x61, 0x69, 0x6c, 0x4d, 0x4d, 0x4f, 0x72, 0x42,
	0x49, 0x4f, 0x12, 0x16, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x4b,
	0x65, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66,
	0x6b, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x4d, 0x4d,
	0x4f, 0x72, 0x42, 0x49, 0x4f, 0x12, 0x16, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x4b, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x46, 0x61, 0x69, 0x6c, 0x53, 0x79,
	0x73, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x70,
	0x66, 0x6b, 0x69, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x6b, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bpfki_proto_rawDescOnce sync.Once
	file_bpfki_proto_rawDescData = file_bpfki_proto_rawDesc
)

func file_bpfki_proto_rawDescGZIP() []byte {
	file_bpfki_proto_rawDescOnce.Do(func() {
		file_bpfki_proto_rawDescData = protoimpl.X.CompressGZIP(file_bpfki_proto_rawDescData)
	})
	return file_bpfki_proto_rawDescData
}

var file_bpfki_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bpfki_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bpfki_proto_goTypes = []interface{}{
	(FailKernRequest_FAILTYPE)(0), // 0: bpfki.FailKernRequest.FAILTYPE
	(*BumpTimeRequest)(nil),       // 1: bpfki.BumpTimeRequest
	(*FailKernRequest)(nil),       // 2: bpfki.FailKernRequest
	(*FailSyscallRequest)(nil),    // 3: bpfki.FailSyscallRequest
	(*StatusResponse)(nil),        // 4: bpfki.StatusResponse
	(*FailKernRequestFrame)(nil),  // 5: bpfki.FailKernRequest.frame
}
var file_bpfki_proto_depIdxs = []int32{
	0,  // 0: bpfki.FailKernRequest.ftype:type_name -> bpfki.FailKernRequest.FAILTYPE
	5,  // 1: bpfki.FailKernRequest.callchain:type_name -> bpfki.FailKernRequest.frame
	1,  // 2: bpfki.BPFKIService.SetTimeVal:input_type -> bpfki.BumpTimeRequest
	1,  // 3: bpfki.BPFKIService.RecoverTimeVal:input_type -> bpfki.BumpTimeRequest
	1,  // 4: bpfki.BPFKIService.SetTimeSpec:input_type -> bpfki.BumpTimeRequest
	1,  // 5: bpfki.BPFKIService.RecoverTimeSpec:input_type -> bpfki.BumpTimeRequest
	2,  // 6: bpfki.BPFKIService.FailMMOrBIO:input_type -> bpfki.FailKernRequest
	2,  // 7: bpfki.BPFKIService.RecoverMMOrBIO:input_type -> bpfki.FailKernRequest
	3,  // 8: bpfki.BPFKIService.FailSyscall:input_type -> bpfki.FailSyscallRequest
	3,  // 9: bpfki.BPFKIService.RecoverSyscall:input_type -> bpfki.FailSyscallRequest
	4,  // 10: bpfki.BPFKIService.SetTimeVal:output_type -> bpfki.StatusResponse
	4,  // 11: bpfki.BPFKIService.RecoverTimeVal:output_type -> bpfki.StatusResponse
	4,  // 12: bpfki.BPFKIService.SetTimeSpec:output_type -> bpfki.StatusResponse
	4,  // 13: bpfki.BPFKIService.RecoverTimeSpec:output_type -> bpfki.StatusResponse
	4,  // 14: bpfki.BPFKIService.FailMMOrBIO:output_type -> bpfki.StatusResponse
	4,  // 15: bpfki.BPFKIService.RecoverMMOrBIO:output_type -> bpfki.StatusResponse
	4,  // 16: bpfki.BPFKIService.FailSyscall:output_type -> bpfki.StatusResponse
	4,  // 17: bpfki.BPFKIService.RecoverSyscall:output_type -> bpfki.StatusResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_bpfki_proto_init() }
func file_bpfki_proto_init() {
	if File_bpfki_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bpfki_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BumpTimeRequest); i {
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
		file_bpfki_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailKernRequest); i {
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
		file_bpfki_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailSyscallRequest); i {
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
		file_bpfki_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_bpfki_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailKernRequestFrame); i {
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
			RawDescriptor: file_bpfki_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bpfki_proto_goTypes,
		DependencyIndexes: file_bpfki_proto_depIdxs,
		EnumInfos:         file_bpfki_proto_enumTypes,
		MessageInfos:      file_bpfki_proto_msgTypes,
	}.Build()
	File_bpfki_proto = out.File
	file_bpfki_proto_rawDesc = nil
	file_bpfki_proto_goTypes = nil
	file_bpfki_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BPFKIServiceClient is the client API for BPFKIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BPFKIServiceClient interface {
	SetTimeVal(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	RecoverTimeVal(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	SetTimeSpec(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	RecoverTimeSpec(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	FailMMOrBIO(ctx context.Context, in *FailKernRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	RecoverMMOrBIO(ctx context.Context, in *FailKernRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	FailSyscall(ctx context.Context, in *FailSyscallRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	RecoverSyscall(ctx context.Context, in *FailSyscallRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type bPFKIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBPFKIServiceClient(cc grpc.ClientConnInterface) BPFKIServiceClient {
	return &bPFKIServiceClient{cc}
}

func (c *bPFKIServiceClient) SetTimeVal(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/SetTimeVal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) RecoverTimeVal(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/RecoverTimeVal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) SetTimeSpec(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/SetTimeSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) RecoverTimeSpec(ctx context.Context, in *BumpTimeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/RecoverTimeSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) FailMMOrBIO(ctx context.Context, in *FailKernRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/FailMMOrBIO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) RecoverMMOrBIO(ctx context.Context, in *FailKernRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/RecoverMMOrBIO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) FailSyscall(ctx context.Context, in *FailSyscallRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/FailSyscall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bPFKIServiceClient) RecoverSyscall(ctx context.Context, in *FailSyscallRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/bpfki.BPFKIService/RecoverSyscall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BPFKIServiceServer is the server API for BPFKIService service.
type BPFKIServiceServer interface {
	SetTimeVal(context.Context, *BumpTimeRequest) (*StatusResponse, error)
	RecoverTimeVal(context.Context, *BumpTimeRequest) (*StatusResponse, error)
	SetTimeSpec(context.Context, *BumpTimeRequest) (*StatusResponse, error)
	RecoverTimeSpec(context.Context, *BumpTimeRequest) (*StatusResponse, error)
	FailMMOrBIO(context.Context, *FailKernRequest) (*StatusResponse, error)
	RecoverMMOrBIO(context.Context, *FailKernRequest) (*StatusResponse, error)
	FailSyscall(context.Context, *FailSyscallRequest) (*StatusResponse, error)
	RecoverSyscall(context.Context, *FailSyscallRequest) (*StatusResponse, error)
}

// UnimplementedBPFKIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBPFKIServiceServer struct {
}

func (*UnimplementedBPFKIServiceServer) SetTimeVal(context.Context, *BumpTimeRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTimeVal not implemented")
}
func (*UnimplementedBPFKIServiceServer) RecoverTimeVal(context.Context, *BumpTimeRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverTimeVal not implemented")
}
func (*UnimplementedBPFKIServiceServer) SetTimeSpec(context.Context, *BumpTimeRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTimeSpec not implemented")
}
func (*UnimplementedBPFKIServiceServer) RecoverTimeSpec(context.Context, *BumpTimeRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverTimeSpec not implemented")
}
func (*UnimplementedBPFKIServiceServer) FailMMOrBIO(context.Context, *FailKernRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FailMMOrBIO not implemented")
}
func (*UnimplementedBPFKIServiceServer) RecoverMMOrBIO(context.Context, *FailKernRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverMMOrBIO not implemented")
}
func (*UnimplementedBPFKIServiceServer) FailSyscall(context.Context, *FailSyscallRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FailSyscall not implemented")
}
func (*UnimplementedBPFKIServiceServer) RecoverSyscall(context.Context, *FailSyscallRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverSyscall not implemented")
}

func RegisterBPFKIServiceServer(s *grpc.Server, srv BPFKIServiceServer) {
	s.RegisterService(&_BPFKIService_serviceDesc, srv)
}

func _BPFKIService_SetTimeVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BumpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).SetTimeVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/SetTimeVal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).SetTimeVal(ctx, req.(*BumpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_RecoverTimeVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BumpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).RecoverTimeVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/RecoverTimeVal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).RecoverTimeVal(ctx, req.(*BumpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_SetTimeSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BumpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).SetTimeSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/SetTimeSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).SetTimeSpec(ctx, req.(*BumpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_RecoverTimeSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BumpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).RecoverTimeSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/RecoverTimeSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).RecoverTimeSpec(ctx, req.(*BumpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_FailMMOrBIO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailKernRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).FailMMOrBIO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/FailMMOrBIO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).FailMMOrBIO(ctx, req.(*FailKernRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_RecoverMMOrBIO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailKernRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).RecoverMMOrBIO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/RecoverMMOrBIO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).RecoverMMOrBIO(ctx, req.(*FailKernRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_FailSyscall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailSyscallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).FailSyscall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/FailSyscall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).FailSyscall(ctx, req.(*FailSyscallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BPFKIService_RecoverSyscall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailSyscallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BPFKIServiceServer).RecoverSyscall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfki.BPFKIService/RecoverSyscall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BPFKIServiceServer).RecoverSyscall(ctx, req.(*FailSyscallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BPFKIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bpfki.BPFKIService",
	HandlerType: (*BPFKIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTimeVal",
			Handler:    _BPFKIService_SetTimeVal_Handler,
		},
		{
			MethodName: "RecoverTimeVal",
			Handler:    _BPFKIService_RecoverTimeVal_Handler,
		},
		{
			MethodName: "SetTimeSpec",
			Handler:    _BPFKIService_SetTimeSpec_Handler,
		},
		{
			MethodName: "RecoverTimeSpec",
			Handler:    _BPFKIService_RecoverTimeSpec_Handler,
		},
		{
			MethodName: "FailMMOrBIO",
			Handler:    _BPFKIService_FailMMOrBIO_Handler,
		},
		{
			MethodName: "RecoverMMOrBIO",
			Handler:    _BPFKIService_RecoverMMOrBIO_Handler,
		},
		{
			MethodName: "FailSyscall",
			Handler:    _BPFKIService_FailSyscall_Handler,
		},
		{
			MethodName: "RecoverSyscall",
			Handler:    _BPFKIService_RecoverSyscall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bpfki.proto",
}
