// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_messages.proto

package common_messages

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e686a2dd39c7bca4, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ErrorReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,json=error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorReply) Reset()         { *m = ErrorReply{} }
func (m *ErrorReply) String() string { return proto.CompactTextString(m) }
func (*ErrorReply) ProtoMessage()    {}
func (*ErrorReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e686a2dd39c7bca4, []int{1}
}

func (m *ErrorReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorReply.Unmarshal(m, b)
}
func (m *ErrorReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorReply.Marshal(b, m, deterministic)
}
func (m *ErrorReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorReply.Merge(m, src)
}
func (m *ErrorReply) XXX_Size() int {
	return xxx_messageInfo_ErrorReply.Size(m)
}
func (m *ErrorReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorReply.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorReply proto.InternalMessageInfo

func (m *ErrorReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ErrorReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "common_messages.ID")
	proto.RegisterType((*ErrorReply)(nil), "common_messages.ErrorReply")
}

func init() { proto.RegisterFile("common_messages.proto", fileDescriptor_e686a2dd39c7bca4) }

var fileDescriptor_e686a2dd39c7bca4 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x56, 0xd2, 0xe3, 0x62, 0xf2, 0x74, 0x11, 0xe2, 0xe3, 0x62, 0xf2,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe0, 0x62, 0x0f,
	0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x02, 0x0b, 0xb2, 0x97, 0x41, 0xb8, 0x4a, 0x56,
	0x5c, 0x5c, 0xae, 0x45, 0x45, 0xf9, 0x45, 0x41, 0xa9, 0x05, 0x39, 0x95, 0x42, 0x62, 0x5c, 0x6c,
	0xc1, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x60, 0xbd, 0x1c, 0x41, 0x6c, 0xc5, 0x60, 0x9e, 0x90, 0x08,
	0x17, 0x2b, 0x58, 0x15, 0x54, 0x37, 0x6b, 0x2a, 0x88, 0x93, 0xc4, 0x06, 0x76, 0x83, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0x25, 0x70, 0xf8, 0x9c, 0x00, 0x00, 0x00,
}