//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// Protocol Buffers describing the Runner API, which is the runner-independent,
// SDK-independent definition of the Beam model.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: org/apache/beam/model/pipeline/v1/standard_window_fns.proto

package pipeline_v1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GlobalWindowsPayload_Enum int32

const (
	GlobalWindowsPayload_PROPERTIES GlobalWindowsPayload_Enum = 0
)

// Enum value maps for GlobalWindowsPayload_Enum.
var (
	GlobalWindowsPayload_Enum_name = map[int32]string{
		0: "PROPERTIES",
	}
	GlobalWindowsPayload_Enum_value = map[string]int32{
		"PROPERTIES": 0,
	}
)

func (x GlobalWindowsPayload_Enum) Enum() *GlobalWindowsPayload_Enum {
	p := new(GlobalWindowsPayload_Enum)
	*p = x
	return p
}

func (x GlobalWindowsPayload_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GlobalWindowsPayload_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[0].Descriptor()
}

func (GlobalWindowsPayload_Enum) Type() protoreflect.EnumType {
	return &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[0]
}

func (x GlobalWindowsPayload_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GlobalWindowsPayload_Enum.Descriptor instead.
func (GlobalWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{0, 0}
}

type FixedWindowsPayload_Enum int32

const (
	FixedWindowsPayload_PROPERTIES FixedWindowsPayload_Enum = 0
)

// Enum value maps for FixedWindowsPayload_Enum.
var (
	FixedWindowsPayload_Enum_name = map[int32]string{
		0: "PROPERTIES",
	}
	FixedWindowsPayload_Enum_value = map[string]int32{
		"PROPERTIES": 0,
	}
)

func (x FixedWindowsPayload_Enum) Enum() *FixedWindowsPayload_Enum {
	p := new(FixedWindowsPayload_Enum)
	*p = x
	return p
}

func (x FixedWindowsPayload_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FixedWindowsPayload_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[1].Descriptor()
}

func (FixedWindowsPayload_Enum) Type() protoreflect.EnumType {
	return &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[1]
}

func (x FixedWindowsPayload_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FixedWindowsPayload_Enum.Descriptor instead.
func (FixedWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{1, 0}
}

type SlidingWindowsPayload_Enum int32

const (
	SlidingWindowsPayload_PROPERTIES SlidingWindowsPayload_Enum = 0
)

// Enum value maps for SlidingWindowsPayload_Enum.
var (
	SlidingWindowsPayload_Enum_name = map[int32]string{
		0: "PROPERTIES",
	}
	SlidingWindowsPayload_Enum_value = map[string]int32{
		"PROPERTIES": 0,
	}
)

func (x SlidingWindowsPayload_Enum) Enum() *SlidingWindowsPayload_Enum {
	p := new(SlidingWindowsPayload_Enum)
	*p = x
	return p
}

func (x SlidingWindowsPayload_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SlidingWindowsPayload_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[2].Descriptor()
}

func (SlidingWindowsPayload_Enum) Type() protoreflect.EnumType {
	return &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[2]
}

func (x SlidingWindowsPayload_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SlidingWindowsPayload_Enum.Descriptor instead.
func (SlidingWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{2, 0}
}

type SessionWindowsPayload_Enum int32

const (
	SessionWindowsPayload_PROPERTIES SessionWindowsPayload_Enum = 0
)

// Enum value maps for SessionWindowsPayload_Enum.
var (
	SessionWindowsPayload_Enum_name = map[int32]string{
		0: "PROPERTIES",
	}
	SessionWindowsPayload_Enum_value = map[string]int32{
		"PROPERTIES": 0,
	}
)

func (x SessionWindowsPayload_Enum) Enum() *SessionWindowsPayload_Enum {
	p := new(SessionWindowsPayload_Enum)
	*p = x
	return p
}

func (x SessionWindowsPayload_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SessionWindowsPayload_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[3].Descriptor()
}

func (SessionWindowsPayload_Enum) Type() protoreflect.EnumType {
	return &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes[3]
}

func (x SessionWindowsPayload_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SessionWindowsPayload_Enum.Descriptor instead.
func (SessionWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{3, 0}
}

// By default, all data in a PCollection is assigned to the single global
// window. See BeamConstants for the time span this window encompasses.
//
// See https://beam.apache.org/documentation/programming-guide/#single-global-window
// for additional details.
type GlobalWindowsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GlobalWindowsPayload) Reset() {
	*x = GlobalWindowsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalWindowsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalWindowsPayload) ProtoMessage() {}

func (x *GlobalWindowsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalWindowsPayload.ProtoReflect.Descriptor instead.
func (*GlobalWindowsPayload) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{0}
}

// A fixed time window represents a consistent duration size, non overlapping
// time interval in the data stream.
//
// See https://beam.apache.org/documentation/programming-guide/#fixed-time-windows
// for additional details.
type FixedWindowsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Represents the size of the window.
	Size *duration.Duration `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	// (Required) Represents the timestamp of when the first window begins.
	// Window N will start at offset + N * size.
	Offset *timestamp.Timestamp `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FixedWindowsPayload) Reset() {
	*x = FixedWindowsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixedWindowsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedWindowsPayload) ProtoMessage() {}

func (x *FixedWindowsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedWindowsPayload.ProtoReflect.Descriptor instead.
func (*FixedWindowsPayload) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{1}
}

func (x *FixedWindowsPayload) GetSize() *duration.Duration {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *FixedWindowsPayload) GetOffset() *timestamp.Timestamp {
	if x != nil {
		return x.Offset
	}
	return nil
}

// A sliding time window represents time intervals in the data stream that can
// overlap. For example, each window might capture 60 seconds worth of data, but
// a new window starts every 30 seconds. The frequency with which sliding
// windows begin is called the period. Therefore, our example would have a
// window size of 60 seconds and a period of 30 seconds.
//
// Because multiple windows overlap, most elements in a data set will belong to
// more than one window. This kind of windowing is useful for taking running
// averages of data; using sliding time windows, you can compute a running
// average of the past 60 seconds’ worth of data, updated every 30 seconds, in
// our example.
//
// See https://beam.apache.org/documentation/programming-guide/#sliding-time-windows
// for additional details.
type SlidingWindowsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Represents the size of the window.
	Size *duration.Duration `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	// (Required) Represents the timestamp of when the first window begins.
	// Window N will start at offset + N * period.
	Offset *timestamp.Timestamp `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// (Required) Represents the amount of time between each start of a window.
	Period *duration.Duration `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *SlidingWindowsPayload) Reset() {
	*x = SlidingWindowsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlidingWindowsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlidingWindowsPayload) ProtoMessage() {}

func (x *SlidingWindowsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlidingWindowsPayload.ProtoReflect.Descriptor instead.
func (*SlidingWindowsPayload) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{2}
}

func (x *SlidingWindowsPayload) GetSize() *duration.Duration {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *SlidingWindowsPayload) GetOffset() *timestamp.Timestamp {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *SlidingWindowsPayload) GetPeriod() *duration.Duration {
	if x != nil {
		return x.Period
	}
	return nil
}

// A session window function defines windows that contain elements that are
// within a certain gap size of another element. Session windowing applies
// on a per-key basis and is useful for data that is irregularly distributed
// with respect to time. For example, a data stream representing user mouse
// activity may have long periods of idle time interspersed with high
// concentrations of clicks. If data arrives after the minimum specified gap
// size duration, this initiates the start of a new window.
//
// See https://beam.apache.org/documentation/programming-guide/#session-windows
// for additional details.
type SessionWindowsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) Minimum duration of gaps between sessions.
	GapSize *duration.Duration `protobuf:"bytes,1,opt,name=gap_size,json=gapSize,proto3" json:"gap_size,omitempty"`
}

func (x *SessionWindowsPayload) Reset() {
	*x = SessionWindowsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionWindowsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionWindowsPayload) ProtoMessage() {}

func (x *SessionWindowsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionWindowsPayload.ProtoReflect.Descriptor instead.
func (*SessionWindowsPayload) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP(), []int{3}
}

func (x *SessionWindowsPayload) GetGapSize() *duration.Duration {
	if x != nil {
		return x.GapSize
	}
	return nil
}

var File_org_apache_beam_model_pipeline_v1_standard_window_fns_proto protoreflect.FileDescriptor

var file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61,
	0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x5f, 0x66, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x37, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61,
	0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x3e, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x36, 0x0a, 0x0a, 0x50, 0x52,
	0x4f, 0x50, 0x45, 0x52, 0x54, 0x49, 0x45, 0x53, 0x10, 0x00, 0x1a, 0x26, 0xa2, 0xb4, 0xfa, 0xc2,
	0x05, 0x20, 0x62, 0x65, 0x61, 0x6d, 0x3a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x6e,
	0x3a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x3a,
	0x76, 0x31, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x78, 0x65, 0x64, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3d, 0x0a,
	0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54,
	0x49, 0x45, 0x53, 0x10, 0x00, 0x1a, 0x25, 0xa2, 0xb4, 0xfa, 0xc2, 0x05, 0x1f, 0x62, 0x65, 0x61,
	0x6d, 0x3a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x6e, 0x3a, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x3a, 0x76, 0x31, 0x22, 0xee, 0x01, 0x0a,
	0x15, 0x53, 0x6c, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x3f, 0x0a, 0x04,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x49,
	0x45, 0x53, 0x10, 0x00, 0x1a, 0x27, 0xa2, 0xb4, 0xfa, 0xc2, 0x05, 0x21, 0x62, 0x65, 0x61, 0x6d,
	0x3a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x6e, 0x3a, 0x73, 0x6c, 0x69, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x3a, 0x76, 0x31, 0x22, 0x8e, 0x01,
	0x0a, 0x15, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x67, 0x61, 0x70, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x67, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x3f, 0x0a,
	0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54,
	0x49, 0x45, 0x53, 0x10, 0x00, 0x1a, 0x27, 0xa2, 0xb4, 0xfa, 0xc2, 0x05, 0x21, 0x62, 0x65, 0x61,
	0x6d, 0x3a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x6e, 0x3a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x3a, 0x76, 0x31, 0x42, 0x80,
	0x01, 0x0a, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65,
	0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x46, 0x6e, 0x73, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f,
	0x73, 0x64, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62,
	0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x76, 0x31, 0x3b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescOnce sync.Once
	file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescData = file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDesc
)

func file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescGZIP() []byte {
	file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescOnce.Do(func() {
		file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescData)
	})
	return file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDescData
}

var file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_goTypes = []interface{}{
	(GlobalWindowsPayload_Enum)(0),  // 0: org.apache.beam.model.pipeline.v1.GlobalWindowsPayload.Enum
	(FixedWindowsPayload_Enum)(0),   // 1: org.apache.beam.model.pipeline.v1.FixedWindowsPayload.Enum
	(SlidingWindowsPayload_Enum)(0), // 2: org.apache.beam.model.pipeline.v1.SlidingWindowsPayload.Enum
	(SessionWindowsPayload_Enum)(0), // 3: org.apache.beam.model.pipeline.v1.SessionWindowsPayload.Enum
	(*GlobalWindowsPayload)(nil),    // 4: org.apache.beam.model.pipeline.v1.GlobalWindowsPayload
	(*FixedWindowsPayload)(nil),     // 5: org.apache.beam.model.pipeline.v1.FixedWindowsPayload
	(*SlidingWindowsPayload)(nil),   // 6: org.apache.beam.model.pipeline.v1.SlidingWindowsPayload
	(*SessionWindowsPayload)(nil),   // 7: org.apache.beam.model.pipeline.v1.SessionWindowsPayload
	(*duration.Duration)(nil),       // 8: google.protobuf.Duration
	(*timestamp.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_depIdxs = []int32{
	8, // 0: org.apache.beam.model.pipeline.v1.FixedWindowsPayload.size:type_name -> google.protobuf.Duration
	9, // 1: org.apache.beam.model.pipeline.v1.FixedWindowsPayload.offset:type_name -> google.protobuf.Timestamp
	8, // 2: org.apache.beam.model.pipeline.v1.SlidingWindowsPayload.size:type_name -> google.protobuf.Duration
	9, // 3: org.apache.beam.model.pipeline.v1.SlidingWindowsPayload.offset:type_name -> google.protobuf.Timestamp
	8, // 4: org.apache.beam.model.pipeline.v1.SlidingWindowsPayload.period:type_name -> google.protobuf.Duration
	8, // 5: org.apache.beam.model.pipeline.v1.SessionWindowsPayload.gap_size:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_init() }
func file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_init() {
	if File_org_apache_beam_model_pipeline_v1_standard_window_fns_proto != nil {
		return
	}
	file_org_apache_beam_model_pipeline_v1_beam_runner_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalWindowsPayload); i {
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
		file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixedWindowsPayload); i {
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
		file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlidingWindowsPayload); i {
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
		file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionWindowsPayload); i {
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
			RawDescriptor: file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_goTypes,
		DependencyIndexes: file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_depIdxs,
		EnumInfos:         file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_enumTypes,
		MessageInfos:      file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_msgTypes,
	}.Build()
	File_org_apache_beam_model_pipeline_v1_standard_window_fns_proto = out.File
	file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_rawDesc = nil
	file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_goTypes = nil
	file_org_apache_beam_model_pipeline_v1_standard_window_fns_proto_depIdxs = nil
}
