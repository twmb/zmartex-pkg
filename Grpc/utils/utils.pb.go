// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/utils.proto

package GrpcUtils

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

type Decimal struct {
	Val                  int64    `protobuf:"varint,1,opt,name=Val,proto3" json:"Val,omitempty"`
	Exp                  int32    `protobuf:"varint,2,opt,name=Exp,proto3" json:"Exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decimal) Reset()         { *m = Decimal{} }
func (m *Decimal) String() string { return proto.CompactTextString(m) }
func (*Decimal) ProtoMessage()    {}
func (*Decimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb28cb36e66ac70, []int{0}
}

func (m *Decimal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decimal.Unmarshal(m, b)
}
func (m *Decimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decimal.Marshal(b, m, deterministic)
}
func (m *Decimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decimal.Merge(m, src)
}
func (m *Decimal) XXX_Size() int {
	return xxx_messageInfo_Decimal.Size(m)
}
func (m *Decimal) XXX_DiscardUnknown() {
	xxx_messageInfo_Decimal.DiscardUnknown(m)
}

var xxx_messageInfo_Decimal proto.InternalMessageInfo

func (m *Decimal) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *Decimal) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func init() {
	proto.RegisterType((*Decimal)(nil), "GrpcUtils.Decimal")
}

func init() { proto.RegisterFile("proto/utils.proto", fileDescriptor_9eb28cb36e66ac70) }

var fileDescriptor_9eb28cb36e66ac70 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2d, 0xc9, 0xcc, 0x29, 0xd6, 0x03, 0xb3, 0x85, 0x38, 0xdd, 0x8b, 0x0a, 0x92,
	0x43, 0x41, 0x02, 0x4a, 0xba, 0x5c, 0xec, 0x2e, 0xa9, 0xc9, 0x99, 0xb9, 0x89, 0x39, 0x42, 0x02,
	0x5c, 0xcc, 0x61, 0x89, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20, 0x26, 0x48, 0xc4,
	0xb5, 0xa2, 0x40, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x08, 0xc4, 0x74, 0xd2, 0x8b, 0xd2, 0x49,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xaf, 0x2a, 0xce, 0x4d, 0x2c, 0x2a,
	0x49, 0xad, 0xd0, 0x2f, 0xc8, 0x4e, 0xd7, 0x07, 0x99, 0x09, 0xb1, 0xc5, 0x1a, 0x6e, 0x7c, 0x12,
	0x1b, 0xd8, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x5f, 0x8f, 0xdf, 0x85, 0x00,
	0x00, 0x00,
}