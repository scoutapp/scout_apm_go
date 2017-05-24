// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

/*
Package event is a generated protocol buffer package.

It is generated from these files:
	event.proto

It has these top-level messages:
	Event
*/
package scout_apm_go

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

type Event struct {
	OrgKey      string            `protobuf:"bytes,1,opt,name=orgKey" json:"orgKey,omitempty"`
	Timestamp   string            `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Duration    int64             `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty"`
	Context     map[string]string `protobuf:"bytes,4,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ChildEvents []*Event          `protobuf:"bytes,5,rep,name=child_events,json=childEvents" json:"child_events,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetOrgKey() string {
	if m != nil {
		return m.OrgKey
	}
	return ""
}

func (m *Event) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Event) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Event) GetChildEvents() []*Event {
	if m != nil {
		return m.ChildEvents
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x7a, 0xcf, 0xc8, 0xc5, 0xea, 0x0a, 0xe2, 0x0b,
	0x89, 0x71, 0xb1, 0xe5, 0x17, 0xa5, 0x7b, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x79, 0x42, 0x32, 0x5c, 0x9c, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12,
	0x4c, 0x60, 0x29, 0x84, 0x80, 0x90, 0x14, 0x17, 0x47, 0x4a, 0x69, 0x51, 0x62, 0x49, 0x66, 0x7e,
	0x9e, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x9c, 0x2f, 0xa4, 0xcb, 0xc5, 0x9e, 0x9c, 0x9f,
	0x57, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0xa2, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xac, 0x07, 0xb6, 0x4a,
	0xcf, 0x19, 0x22, 0xea, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0x53, 0x23, 0xa4, 0xc9, 0xc5, 0x93,
	0x9c, 0x91, 0x99, 0x93, 0x12, 0x0f, 0x76, 0x5f, 0xb1, 0x04, 0x2b, 0x58, 0x0f, 0x1b, 0x44, 0x4f,
	0x10, 0x37, 0x58, 0x0e, 0xcc, 0x2e, 0x96, 0xb2, 0xe2, 0xe2, 0x41, 0x36, 0x43, 0x48, 0x80, 0x8b,
	0x39, 0x1b, 0xee, 0x70, 0x10, 0x53, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0xea,
	0x62, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0x31, 0x89, 0x0d, 0xec, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x49, 0xec, 0x3d, 0x07, 0x01, 0x00, 0x00,
}
