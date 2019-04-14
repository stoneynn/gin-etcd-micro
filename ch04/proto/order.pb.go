// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order.proto

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

type GetOrderInfoReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderInfoReq) Reset()         { *m = GetOrderInfoReq{} }
func (m *GetOrderInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetOrderInfoReq) ProtoMessage()    {}
func (*GetOrderInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{0}
}

func (m *GetOrderInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderInfoReq.Unmarshal(m, b)
}
func (m *GetOrderInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderInfoReq.Marshal(b, m, deterministic)
}
func (m *GetOrderInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderInfoReq.Merge(m, src)
}
func (m *GetOrderInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetOrderInfoReq.Size(m)
}
func (m *GetOrderInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderInfoReq proto.InternalMessageInfo

func (m *GetOrderInfoReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetOrderInfoRes struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderInfoRes) Reset()         { *m = GetOrderInfoRes{} }
func (m *GetOrderInfoRes) String() string { return proto.CompactTextString(m) }
func (*GetOrderInfoRes) ProtoMessage()    {}
func (*GetOrderInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{1}
}

func (m *GetOrderInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderInfoRes.Unmarshal(m, b)
}
func (m *GetOrderInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderInfoRes.Marshal(b, m, deterministic)
}
func (m *GetOrderInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderInfoRes.Merge(m, src)
}
func (m *GetOrderInfoRes) XXX_Size() int {
	return xxx_messageInfo_GetOrderInfoRes.Size(m)
}
func (m *GetOrderInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderInfoRes proto.InternalMessageInfo

func (m *GetOrderInfoRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOrderInfoReq)(nil), "proto.GetOrderInfoReq")
	proto.RegisterType((*GetOrderInfoRes)(nil), "proto.GetOrderInfoRes")
}

func init() { proto.RegisterFile("proto/order.proto", fileDescriptor_f65b0626cc3aada8) }

var fileDescriptor_f65b0626cc3aada8 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x22,
	0x17, 0xbf, 0x7b, 0x6a, 0x89, 0x3f, 0x48, 0xc2, 0x33, 0x2f, 0x2d, 0x3f, 0x28, 0xb5, 0x50, 0x88,
	0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x88, 0x29, 0x33, 0x45, 0x49,
	0x19, 0x5d, 0x49, 0xb1, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x71, 0x3a, 0x58, 0x0d, 0x67, 0x10, 0x88,
	0x69, 0xe4, 0xc9, 0xc5, 0x0a, 0x56, 0x21, 0xe4, 0xc0, 0xc5, 0x83, 0xac, 0x5a, 0x48, 0x0c, 0x62,
	0x9f, 0x1e, 0x9a, 0x2d, 0x52, 0xd8, 0xc5, 0x8b, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x12, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x3c, 0xc2, 0x97, 0xb5, 0x00, 0x00, 0x00,
}