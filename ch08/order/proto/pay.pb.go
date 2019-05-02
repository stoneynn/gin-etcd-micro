// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pay.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPayInfoReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPayInfoReq) Reset()         { *m = GetPayInfoReq{} }
func (m *GetPayInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetPayInfoReq) ProtoMessage()    {}
func (*GetPayInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5aad0ac58a7583, []int{0}
}

func (m *GetPayInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPayInfoReq.Unmarshal(m, b)
}
func (m *GetPayInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPayInfoReq.Marshal(b, m, deterministic)
}
func (m *GetPayInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPayInfoReq.Merge(m, src)
}
func (m *GetPayInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetPayInfoReq.Size(m)
}
func (m *GetPayInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPayInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPayInfoReq proto.InternalMessageInfo

func (m *GetPayInfoReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPayInfoRes struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// 数组
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPayInfoRes) Reset()         { *m = GetPayInfoRes{} }
func (m *GetPayInfoRes) String() string { return proto.CompactTextString(m) }
func (*GetPayInfoRes) ProtoMessage()    {}
func (*GetPayInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e5aad0ac58a7583, []int{1}
}

func (m *GetPayInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPayInfoRes.Unmarshal(m, b)
}
func (m *GetPayInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPayInfoRes.Marshal(b, m, deterministic)
}
func (m *GetPayInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPayInfoRes.Merge(m, src)
}
func (m *GetPayInfoRes) XXX_Size() int {
	return xxx_messageInfo_GetPayInfoRes.Size(m)
}
func (m *GetPayInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPayInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPayInfoRes proto.InternalMessageInfo

func (m *GetPayInfoRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetPayInfoRes) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*GetPayInfoReq)(nil), "proto.GetPayInfoReq")
	proto.RegisterType((*GetPayInfoRes)(nil), "proto.GetPayInfoRes")
}

func init() { proto.RegisterFile("proto/pay.proto", fileDescriptor_6e5aad0ac58a7583) }

var fileDescriptor_6e5aad0ac58a7583 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x48, 0xac, 0xd4, 0x03, 0xb3, 0x84, 0x58, 0xc1, 0x94, 0x92, 0x3c, 0x17, 0xaf,
	0x7b, 0x6a, 0x49, 0x40, 0x62, 0xa5, 0x67, 0x5e, 0x5a, 0x7e, 0x50, 0x6a, 0xa1, 0x10, 0x1f, 0x17,
	0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0x66, 0x8a, 0x92, 0x25, 0xaa,
	0x82, 0x62, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0xb0, 0x0a, 0xce, 0x20, 0x10, 0x53, 0x48,
	0x8c, 0x8b, 0xad, 0x2c, 0x31, 0xa7, 0x34, 0xb5, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08,
	0xca, 0x33, 0x72, 0xe4, 0x62, 0x0e, 0x48, 0xac, 0x14, 0xb2, 0xe2, 0xe2, 0x42, 0x98, 0x20, 0x24,
	0x02, 0xb1, 0x5f, 0x0f, 0xc5, 0x56, 0x29, 0x6c, 0xa2, 0xc5, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x61,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xe7, 0xbc, 0x24, 0xbf, 0x00, 0x00, 0x00,
}