// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProdService.proto

package Services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProdRequest struct {
	// @inject_tag: json:"size" form:"size"
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size" form:"size"`
	// @inject_tag: json:"pid" form:"pid" uri:"pid"
	ProdId               int32    `protobuf:"varint,2,opt,name=prod_id,json=prodId,proto3" json:"pid" form:"pid" uri:"pid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRequest) Reset()         { *m = ProdRequest{} }
func (m *ProdRequest) String() string { return proto.CompactTextString(m) }
func (*ProdRequest) ProtoMessage()    {}
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ProdService_5640c0dfd9de6d95, []int{0}
}
func (m *ProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRequest.Unmarshal(m, b)
}
func (m *ProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRequest.Marshal(b, m, deterministic)
}
func (dst *ProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRequest.Merge(dst, src)
}
func (m *ProdRequest) XXX_Size() int {
	return xxx_messageInfo_ProdRequest.Size(m)
}
func (m *ProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRequest proto.InternalMessageInfo

func (m *ProdRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ProdRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdListResponse struct {
	Data                 []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProdListResponse) Reset()         { *m = ProdListResponse{} }
func (m *ProdListResponse) String() string { return proto.CompactTextString(m) }
func (*ProdListResponse) ProtoMessage()    {}
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ProdService_5640c0dfd9de6d95, []int{1}
}
func (m *ProdListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdListResponse.Unmarshal(m, b)
}
func (m *ProdListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdListResponse.Marshal(b, m, deterministic)
}
func (dst *ProdListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdListResponse.Merge(dst, src)
}
func (m *ProdListResponse) XXX_Size() int {
	return xxx_messageInfo_ProdListResponse.Size(m)
}
func (m *ProdListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdListResponse proto.InternalMessageInfo

func (m *ProdListResponse) GetData() []*ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProdDeatilResponse struct {
	Data                 *ProdModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProdDeatilResponse) Reset()         { *m = ProdDeatilResponse{} }
func (m *ProdDeatilResponse) String() string { return proto.CompactTextString(m) }
func (*ProdDeatilResponse) ProtoMessage()    {}
func (*ProdDeatilResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ProdService_5640c0dfd9de6d95, []int{2}
}
func (m *ProdDeatilResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdDeatilResponse.Unmarshal(m, b)
}
func (m *ProdDeatilResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdDeatilResponse.Marshal(b, m, deterministic)
}
func (dst *ProdDeatilResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdDeatilResponse.Merge(dst, src)
}
func (m *ProdDeatilResponse) XXX_Size() int {
	return xxx_messageInfo_ProdDeatilResponse.Size(m)
}
func (m *ProdDeatilResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdDeatilResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdDeatilResponse proto.InternalMessageInfo

func (m *ProdDeatilResponse) GetData() *ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdRequest)(nil), "Services.ProdRequest")
	proto.RegisterType((*ProdListResponse)(nil), "Services.ProdListResponse")
	proto.RegisterType((*ProdDeatilResponse)(nil), "Services.ProdDeatilResponse")
}

func init() { proto.RegisterFile("ProdService.proto", fileDescriptor_ProdService_5640c0dfd9de6d95) }

var fileDescriptor_ProdService_5640c0dfd9de6d95 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x28, 0xca, 0x4f,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80,
	0x72, 0x8b, 0xa5, 0x78, 0x7c, 0xf3, 0x53, 0x52, 0x73, 0x8a, 0x21, 0xe2, 0x4a, 0x56, 0x5c, 0xdc,
	0x20, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xc5, 0x99, 0x55,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0xb6, 0x90, 0x38, 0x17, 0x7b, 0x41, 0x51,
	0x7e, 0x4a, 0x7c, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x98, 0x0d, 0xc4, 0xf5, 0x4c, 0x51, 0xb2, 0xe6,
	0x12, 0x00, 0xe9, 0xf5, 0xc9, 0x2c, 0x2e, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x52, 0xe7, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd6,
	0x83, 0x59, 0xab, 0x07, 0x52, 0x09, 0xb6, 0x39, 0x08, 0xac, 0x40, 0xc9, 0x96, 0x4b, 0x08, 0x24,
	0xe4, 0x92, 0x9a, 0x58, 0x92, 0x99, 0x83, 0x45, 0x3b, 0x23, 0x5e, 0xed, 0x46, 0x53, 0x19, 0x21,
	0x0e, 0x87, 0x2a, 0x10, 0x72, 0xe0, 0xe2, 0x76, 0x4f, 0x2d, 0x81, 0x39, 0x47, 0x48, 0x14, 0x55,
	0x27, 0xd4, 0x7b, 0x52, 0x52, 0xa8, 0xc2, 0x28, 0x2e, 0x77, 0xe1, 0xe2, 0x85, 0x9a, 0x00, 0x71,
	0x13, 0x2e, 0x33, 0x64, 0x50, 0x85, 0x51, 0x3d, 0x90, 0xc4, 0x06, 0x0e, 0x56, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x17, 0x6d, 0xed, 0x83, 0x01, 0x00, 0x00,
}