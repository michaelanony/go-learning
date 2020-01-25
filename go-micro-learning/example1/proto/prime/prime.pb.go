// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prime/prime.proto

package prime

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

type PrimeRequest struct {
	Input                int64    `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeRequest) Reset()         { *m = PrimeRequest{} }
func (m *PrimeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeRequest) ProtoMessage()    {}
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308b80f54e3a336, []int{0}
}

func (m *PrimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeRequest.Unmarshal(m, b)
}
func (m *PrimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeRequest.Merge(m, src)
}
func (m *PrimeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeRequest.Size(m)
}
func (m *PrimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeRequest proto.InternalMessageInfo

func (m *PrimeRequest) GetInput() int64 {
	if m != nil {
		return m.Input
	}
	return 0
}

type PrimeResponse struct {
	Output               int64    `protobuf:"varint,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeResponse) Reset()         { *m = PrimeResponse{} }
func (m *PrimeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeResponse) ProtoMessage()    {}
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308b80f54e3a336, []int{1}
}

func (m *PrimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeResponse.Unmarshal(m, b)
}
func (m *PrimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeResponse.Merge(m, src)
}
func (m *PrimeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeResponse.Size(m)
}
func (m *PrimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeResponse proto.InternalMessageInfo

func (m *PrimeResponse) GetOutput() int64 {
	if m != nil {
		return m.Output
	}
	return 0
}

func init() {
	proto.RegisterType((*PrimeRequest)(nil), "PrimeRequest")
	proto.RegisterType((*PrimeResponse)(nil), "PrimeResponse")
}

func init() { proto.RegisterFile("prime/prime.proto", fileDescriptor_f308b80f54e3a336) }

var fileDescriptor_f308b80f54e3a336 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0xcc,
	0x4d, 0xd5, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0x2a, 0x5c, 0x3c, 0x01, 0x20,
	0x6e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x41, 0x69,
	0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0xa3, 0xa4, 0xce, 0xc5, 0x0b, 0x55, 0x55,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x96, 0x5f, 0x5a, 0x82, 0x50, 0x07, 0xe5,
	0x19, 0x99, 0x70, 0xb1, 0x82, 0x15, 0x0a, 0x69, 0x73, 0x71, 0xb8, 0xa7, 0x96, 0x40, 0xd8, 0xbc,
	0x7a, 0xc8, 0x56, 0x48, 0xf1, 0xe9, 0xa1, 0x98, 0xa5, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0x8b, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xf9, 0x2f, 0x85, 0xa0, 0x00, 0x00, 0x00,
}