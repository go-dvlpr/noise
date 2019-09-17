// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testproto.proto

package proto

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

type Msg struct {
	Autor                string   `protobuf:"bytes,1,opt,name=Autor,json=autor,proto3" json:"Autor,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=Text,json=text,proto3" json:"Text,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=Date,json=date,proto3" json:"Date,omitempty"`
	VN                   string   `protobuf:"bytes,4,opt,name=VN,json=vN,proto3" json:"VN,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33b1c8966e93a98a, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetAutor() string {
	if m != nil {
		return m.Autor
	}
	return ""
}

func (m *Msg) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Msg) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Msg) GetVN() string {
	if m != nil {
		return m.VN
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "proto.Msg")
}

func init() { proto.RegisterFile("testproto.proto", fileDescriptor_33b1c8966e93a98a) }

var fileDescriptor_33b1c8966e93a98a = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x49, 0x2d, 0x2e,
	0x29, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x42, 0xac, 0x60, 0x4a, 0x29, 0x98, 0x8b, 0xd9,
	0xb7, 0x38, 0x5d, 0x48, 0x84, 0x8b, 0xd5, 0xb1, 0xb4, 0x24, 0xbf, 0x48, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x88, 0x35, 0x11, 0xc4, 0x11, 0x12, 0xe2, 0x62, 0x09, 0x49, 0xad, 0x28, 0x91, 0x60,
	0x02, 0x0b, 0xb2, 0x94, 0xa4, 0x56, 0x94, 0x80, 0xc4, 0x5c, 0x12, 0x4b, 0x52, 0x25, 0x98, 0x21,
	0x62, 0x29, 0x89, 0x25, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x61, 0x7e, 0x12, 0x2c, 0x60, 0x11, 0xa6,
	0x32, 0xbf, 0x24, 0x36, 0xb0, 0xd9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x3b, 0x0e,
	0x55, 0x75, 0x00, 0x00, 0x00,
}
