// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/example.proto

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "tagger"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Example struct {
	WithNewTags          string   `protobuf:"bytes,1,opt,name=with_new_tags,json=withNewTags,proto3" json:"with_new_tags,omitempty" graphql:"withNewTags,optional"`
	WithNewMultiple      string   `protobuf:"bytes,2,opt,name=with_new_multiple,json=withNewMultiple,proto3" json:"with_new_multiple,omitempty" graphql:"withNewTags,optional" xml:"multi,omitempty"`
	ReplaceDefault       string   `protobuf:"bytes,3,opt,name=replace_default,json=replaceDefault,proto3" json:"replacePrevious"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_381813f5f80f4838, []int{0}
}
func (m *Example) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example.Unmarshal(m, b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example.Marshal(b, m, deterministic)
}
func (dst *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(dst, src)
}
func (m *Example) XXX_Size() int {
	return xxx_messageInfo_Example.Size(m)
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetWithNewTags() string {
	if m != nil {
		return m.WithNewTags
	}
	return ""
}

func (m *Example) GetWithNewMultiple() string {
	if m != nil {
		return m.WithNewMultiple
	}
	return ""
}

func (m *Example) GetReplaceDefault() string {
	if m != nil {
		return m.ReplaceDefault
	}
	return ""
}

type SecondMessage struct {
	WithNewTags          string   `protobuf:"bytes,1,opt,name=with_new_tags,json=withNewTags,proto3" json:"with_new_tags,omitempty" graphql:"withNewTags,optional"`
	WithNewMultiple      string   `protobuf:"bytes,2,opt,name=with_new_multiple,json=withNewMultiple,proto3" json:"with_new_multiple,omitempty" graphql:"withNewTags,optional" xml:"multi,omitempty"`
	ReplaceDefault       string   `protobuf:"bytes,3,opt,name=replace_default,json=replaceDefault,proto3" json:"replacePrevious"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecondMessage) Reset()         { *m = SecondMessage{} }
func (m *SecondMessage) String() string { return proto.CompactTextString(m) }
func (*SecondMessage) ProtoMessage()    {}
func (*SecondMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_381813f5f80f4838, []int{1}
}
func (m *SecondMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecondMessage.Unmarshal(m, b)
}
func (m *SecondMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecondMessage.Marshal(b, m, deterministic)
}
func (dst *SecondMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondMessage.Merge(dst, src)
}
func (m *SecondMessage) XXX_Size() int {
	return xxx_messageInfo_SecondMessage.Size(m)
}
func (m *SecondMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SecondMessage proto.InternalMessageInfo

func (m *SecondMessage) GetWithNewTags() string {
	if m != nil {
		return m.WithNewTags
	}
	return ""
}

func (m *SecondMessage) GetWithNewMultiple() string {
	if m != nil {
		return m.WithNewMultiple
	}
	return ""
}

func (m *SecondMessage) GetReplaceDefault() string {
	if m != nil {
		return m.ReplaceDefault
	}
	return ""
}

func init() {
	proto.RegisterType((*Example)(nil), "example.Example")
	proto.RegisterType((*SecondMessage)(nil), "example.SecondMessage")
}

func init() { proto.RegisterFile("example/example.proto", fileDescriptor_example_381813f5f80f4838) }

var fileDescriptor_example_381813f5f80f4838 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x70, 0x49, 0x62, 0x7a, 0x7a, 0x6a, 0x91, 0x3e, 0x84, 0x82, 0xc8, 0x2a, 0xfd, 0x67, 0xe4,
	0x62, 0x77, 0x85, 0x28, 0x10, 0x72, 0xe7, 0xe2, 0x2d, 0xcf, 0x2c, 0xc9, 0x88, 0xcf, 0x4b, 0x2d,
	0x8f, 0x2f, 0x49, 0x4c, 0x2f, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x52, 0x9e, 0xd5, 0x32,
	0x8f, 0x59, 0x2e, 0xbd, 0x28, 0xb1, 0x20, 0xa3, 0x30, 0xc7, 0x4a, 0x09, 0xa4, 0xc4, 0x2f, 0xb5,
	0x3c, 0x24, 0x31, 0xbd, 0x58, 0x27, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0x2f, 0x31, 0x47, 0x29, 0x88,
	0x1b, 0x49, 0x58, 0x28, 0x95, 0x4b, 0x10, 0x6e, 0x50, 0x6e, 0x69, 0x4e, 0x49, 0x66, 0x41, 0x4e,
	0xaa, 0x04, 0x13, 0xd8, 0x30, 0x4b, 0x90, 0x61, 0x26, 0xf8, 0x0d, 0x53, 0xa8, 0xc8, 0xcd, 0xb1,
	0x52, 0x02, 0x6b, 0xd4, 0xc9, 0xcf, 0xcd, 0x2c, 0x49, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0x0a, 0xe2,
	0x87, 0x2a, 0xf6, 0x85, 0x9a, 0x28, 0xe4, 0xc2, 0xc5, 0x5f, 0x94, 0x5a, 0x90, 0x93, 0x98, 0x9c,
	0x1a, 0x9f, 0x92, 0x9a, 0x96, 0x58, 0x9a, 0x53, 0x22, 0xc1, 0x0c, 0xb6, 0x44, 0x1a, 0x64, 0x89,
	0x58, 0x56, 0x71, 0x7e, 0x9e, 0x95, 0x12, 0x54, 0x45, 0x40, 0x51, 0x6a, 0x59, 0x66, 0x7e, 0x69,
	0xb1, 0x52, 0x10, 0x1f, 0x54, 0xc4, 0x05, 0xa2, 0x45, 0xa9, 0x95, 0x89, 0x8b, 0x37, 0x38, 0x35,
	0x39, 0x3f, 0x2f, 0xc5, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x7d, 0x84, 0x86, 0x43, 0x12, 0x1b, 0x38,
	0x41, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x93, 0x02, 0xdb, 0x9b, 0x47, 0x02, 0x00, 0x00,
}
