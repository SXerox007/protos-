// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runprotobasic/proto/basic.proto

package basic

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageExample struct {
	MessageString        string   `protobuf:"bytes,1,opt,name=messageString,proto3" json:"messageString,omitempty"`
	Content              int32    `protobuf:"varint,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageExample) Reset()         { *m = MessageExample{} }
func (m *MessageExample) String() string { return proto.CompactTextString(m) }
func (*MessageExample) ProtoMessage()    {}
func (*MessageExample) Descriptor() ([]byte, []int) {
	return fileDescriptor_7070f7cb0ec6af31, []int{0}
}

func (m *MessageExample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageExample.Unmarshal(m, b)
}
func (m *MessageExample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageExample.Marshal(b, m, deterministic)
}
func (m *MessageExample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageExample.Merge(m, src)
}
func (m *MessageExample) XXX_Size() int {
	return xxx_messageInfo_MessageExample.Size(m)
}
func (m *MessageExample) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageExample.DiscardUnknown(m)
}

var xxx_messageInfo_MessageExample proto.InternalMessageInfo

func (m *MessageExample) GetMessageString() string {
	if m != nil {
		return m.MessageString
	}
	return ""
}

func (m *MessageExample) GetContent() int32 {
	if m != nil {
		return m.Content
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageExample)(nil), "MessageExample")
}

func init() { proto.RegisterFile("runprotobasic/proto/basic.proto", fileDescriptor_7070f7cb0ec6af31) }

var fileDescriptor_7070f7cb0ec6af31 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2a, 0xcd, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x4a, 0x2c, 0xce, 0x4c, 0xd6, 0x07, 0x33, 0xf5, 0xc1, 0x6c, 0x3d,
	0x30, 0x5b, 0x29, 0x80, 0x8b, 0xcf, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0xd5, 0xb5, 0x22, 0x31,
	0xb7, 0x20, 0x27, 0x55, 0x48, 0x85, 0x8b, 0x37, 0x17, 0x22, 0x12, 0x5c, 0x52, 0x94, 0x99, 0x97,
	0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2a, 0x28, 0x24, 0xc1, 0xc5, 0x9e, 0x9c, 0x9f,
	0x57, 0x92, 0x9a, 0x57, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe3, 0x26, 0xb1, 0x81,
	0x0d, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x93, 0xcf, 0x8e, 0x63, 0x7b, 0x00, 0x00, 0x00,
}
