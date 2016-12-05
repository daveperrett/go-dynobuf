// Code generated by protoc-gen-go.
// source: LogEntry.proto
// DO NOT EDIT!

/*
Package dynobuf is a generated protocol buffer package.

It is generated from these files:
	LogEntry.proto
	LogEntryList.proto

It has these top-level messages:
	LogEntry
	LogEntryList
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

type LogEntry struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Database  string `protobuf:"bytes,2,opt,name=database" json:"database,omitempty"`
	Process   string `protobuf:"bytes,3,opt,name=process" json:"process,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Data      string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*LogEntry)(nil), "dynobuf.LogEntry")
}

func init() { proto.RegisterFile("LogEntry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xf3, 0xc9, 0x4f, 0x77,
	0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xa9, 0xcc, 0xcb,
	0x4f, 0x2a, 0x4d, 0x53, 0x6a, 0x62, 0xe4, 0xe2, 0x80, 0xc9, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x01, 0x59, 0x42, 0x52, 0x5c, 0x1c, 0x29, 0x89, 0x25,
	0x89, 0x49, 0x89, 0xc5, 0xa9, 0x12, 0x4c, 0x60, 0x51, 0x38, 0x5f, 0x48, 0x82, 0x8b, 0x1d, 0x68,
	0x54, 0x72, 0x6a, 0x71, 0xb1, 0x04, 0x33, 0x58, 0x0a, 0xc6, 0x15, 0x92, 0xe1, 0xe2, 0x2c, 0xc9,
	0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x01, 0xca, 0x31, 0x07, 0x21, 0x04, 0x84,
	0x84, 0xb8, 0x58, 0x40, 0x66, 0x48, 0xb0, 0x82, 0x35, 0x81, 0xd9, 0x4e, 0x8a, 0x5c, 0xa2, 0x99,
	0xf9, 0x7a, 0x20, 0x27, 0xe5, 0x26, 0x96, 0x64, 0x26, 0x43, 0xdc, 0x08, 0x74, 0x9d, 0x13, 0xdc,
	0x69, 0x49, 0x6c, 0x60, 0x31, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xa2, 0x0f, 0x01,
	0xc9, 0x00, 0x00, 0x00,
}
