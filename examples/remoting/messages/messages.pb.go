// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Start
	StartRemote
	Ping
	Pong
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/rogeralsing/gam/actor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Start struct {
}

func (m *Start) Reset()                    { *m = Start{} }
func (m *Start) String() string            { return proto.CompactTextString(m) }
func (*Start) ProtoMessage()               {}
func (*Start) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StartRemote struct {
	Sender *actor.PID `protobuf:"bytes,1,opt,name=Sender,json=sender" json:"Sender,omitempty"`
}

func (m *StartRemote) Reset()                    { *m = StartRemote{} }
func (m *StartRemote) String() string            { return proto.CompactTextString(m) }
func (*StartRemote) ProtoMessage()               {}
func (*StartRemote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartRemote) GetSender() *actor.PID {
	if m != nil {
		return m.Sender
	}
	return nil
}

type Ping struct {
	Sender *actor.PID `protobuf:"bytes,1,opt,name=Sender,json=sender" json:"Sender,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ping) GetSender() *actor.PID {
	if m != nil {
		return m.Sender
	}
	return nil
}

type Pong struct {
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Start)(nil), "messages.Start")
	proto.RegisterType((*StartRemote)(nil), "messages.StartRemote")
	proto.RegisterType((*Ping)(nil), "messages.Ping")
	proto.RegisterType((*Pong)(nil), "messages.Pong")
}

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0xf2, 0xd3, 0x53, 0x8b,
	0x12, 0x73, 0x8a, 0x33, 0xf3, 0xd2, 0xf5, 0xd3, 0x13, 0x73, 0xf5, 0x13, 0x93, 0x4b, 0xf2, 0x8b,
	0x20, 0x24, 0x44, 0x9f, 0x12, 0x3b, 0x17, 0x6b, 0x70, 0x49, 0x62, 0x51, 0x89, 0x92, 0x21, 0x17,
	0x37, 0x98, 0x11, 0x94, 0x9a, 0x9b, 0x5f, 0x92, 0x2a, 0xa4, 0xc4, 0xc5, 0x16, 0x9c, 0x9a, 0x97,
	0x92, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xa5, 0x07, 0xd1, 0x15, 0xe0, 0xe9,
	0x12, 0xc4, 0x56, 0x0c, 0x96, 0x51, 0xd2, 0xe2, 0x62, 0x09, 0x00, 0x9a, 0x4c, 0x94, 0x5a, 0x36,
	0xa0, 0xda, 0xfc, 0xbc, 0xf4, 0x24, 0x36, 0xb0, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0x57, 0x76, 0x99, 0xc0, 0x00, 0x00, 0x00,
}