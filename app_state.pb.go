// Code generated by protoc-gen-go.
// source: app_state.proto
// DO NOT EDIT!

/*
Package dynobuf is a generated protocol buffer package.

It is generated from these files:
	app_state.proto
	log_entry.proto
	log_entry_list.proto
	metric_data_point.proto
	metric_data_point_list.proto
	metric_state.proto

It has these top-level messages:
	AppState
	LogEntry
	LogEntryList
	MetricDataPoint
	MetricDataPointList
	MetricState
*/
package dynobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AppState struct {
	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	Total         int64  `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *AppState) Reset()                    { *m = AppState{} }
func (m *AppState) String() string            { return proto.CompactTextString(m) }
func (*AppState) ProtoMessage()               {}
func (*AppState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*AppState)(nil), "dynobuf.AppState")
}

func init() { proto.RegisterFile("app_state.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x28, 0x88,
	0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xa9, 0xcc,
	0xcb, 0x4f, 0x2a, 0x4d, 0x53, 0x72, 0xe7, 0xe2, 0x70, 0x2c, 0x28, 0x08, 0x06, 0x49, 0x09, 0xa9,
	0x72, 0xf1, 0x01, 0xd5, 0xe5, 0x64, 0x26, 0x27, 0x96, 0x64, 0xe6, 0xe7, 0xc5, 0x67, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xf1, 0x22, 0x89, 0x7a, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96,
	0xe4, 0x97, 0x24, 0xe6, 0x48, 0x30, 0x01, 0x65, 0x99, 0x83, 0x20, 0x1c, 0x27, 0x71, 0x2e, 0xd1,
	0xcc, 0x7c, 0x3d, 0x90, 0xb1, 0xb9, 0x40, 0x85, 0xc9, 0x10, 0x7b, 0x80, 0x36, 0x24, 0xb1, 0x81,
	0x59, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xd7, 0x59, 0x0d, 0x84, 0x00, 0x00, 0x00,
}
