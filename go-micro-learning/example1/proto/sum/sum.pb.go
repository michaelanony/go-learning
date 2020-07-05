// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum/sum.services

package sum

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
// is compatible with the services package it is being compiled against.
// A compilation error at this line likely means your copy of the
// services package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the services package

type SumRequest struct {
	Input                int64    `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeb37e63dbd73f33, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetInput() int64 {
	if m != nil {
		return m.Input
	}
	return 0
}

type SumResponse struct {
	Output               int64    `protobuf:"varint,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeb37e63dbd73f33, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetOutput() int64 {
	if m != nil {
		return m.Output
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "SumRequest")
	proto.RegisterType((*SumResponse)(nil), "SumResponse")
}

func init() { proto.RegisterFile("sum/sum.services", fileDescriptor_aeb37e63dbd73f33) }

var fileDescriptor_aeb37e63dbd73f33 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xcd, 0xd5,
	0x2f, 0x2e, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe2, 0xe2, 0x0a, 0x2e, 0xcd,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x2b, 0x28, 0x2d,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70, 0x94, 0x54, 0xb9, 0xb8, 0xc1, 0x6a, 0x8a,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0xf2, 0x4b, 0x4b, 0x10, 0xaa, 0xa0, 0x3c,
	0x23, 0x1d, 0x2e, 0xe6, 0xe0, 0xd2, 0x5c, 0x21, 0x55, 0x2e, 0x36, 0xf7, 0xd4, 0x12, 0x10, 0x8b,
	0x5b, 0x0f, 0x61, 0xb4, 0x14, 0x8f, 0x1e, 0x92, 0x19, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xfb, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x8e, 0x8b, 0x60, 0x90, 0x00, 0x00, 0x00,
}
