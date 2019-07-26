// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discount.proto

package protocols

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VerifyDiscountForm struct {
	Userid               string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Prodid               string   `protobuf:"bytes,2,opt,name=prodid,proto3" json:"prodid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyDiscountForm) Reset()         { *m = VerifyDiscountForm{} }
func (m *VerifyDiscountForm) String() string { return proto.CompactTextString(m) }
func (*VerifyDiscountForm) ProtoMessage()    {}
func (*VerifyDiscountForm) Descriptor() ([]byte, []int) {
	return fileDescriptor_78985be6a51b316f, []int{0}
}

func (m *VerifyDiscountForm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyDiscountForm.Unmarshal(m, b)
}
func (m *VerifyDiscountForm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyDiscountForm.Marshal(b, m, deterministic)
}
func (m *VerifyDiscountForm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyDiscountForm.Merge(m, src)
}
func (m *VerifyDiscountForm) XXX_Size() int {
	return xxx_messageInfo_VerifyDiscountForm.Size(m)
}
func (m *VerifyDiscountForm) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyDiscountForm.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyDiscountForm proto.InternalMessageInfo

func (m *VerifyDiscountForm) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *VerifyDiscountForm) GetProdid() string {
	if m != nil {
		return m.Prodid
	}
	return ""
}

type DiscountInfo struct {
	Pct                  float32  `protobuf:"fixed32,1,opt,name=pct,proto3" json:"pct,omitempty"`
	ValueInCents         int64    `protobuf:"varint,2,opt,name=value_in_cents,json=valueInCents,proto3" json:"value_in_cents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscountInfo) Reset()         { *m = DiscountInfo{} }
func (m *DiscountInfo) String() string { return proto.CompactTextString(m) }
func (*DiscountInfo) ProtoMessage()    {}
func (*DiscountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_78985be6a51b316f, []int{1}
}

func (m *DiscountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscountInfo.Unmarshal(m, b)
}
func (m *DiscountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscountInfo.Marshal(b, m, deterministic)
}
func (m *DiscountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscountInfo.Merge(m, src)
}
func (m *DiscountInfo) XXX_Size() int {
	return xxx_messageInfo_DiscountInfo.Size(m)
}
func (m *DiscountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DiscountInfo proto.InternalMessageInfo

func (m *DiscountInfo) GetPct() float32 {
	if m != nil {
		return m.Pct
	}
	return 0
}

func (m *DiscountInfo) GetValueInCents() int64 {
	if m != nil {
		return m.ValueInCents
	}
	return 0
}

func init() {
	proto.RegisterType((*VerifyDiscountForm)(nil), "protocols.VerifyDiscountForm")
	proto.RegisterType((*DiscountInfo)(nil), "protocols.DiscountInfo")
}

func init() { proto.RegisterFile("discount.proto", fileDescriptor_78985be6a51b316f) }

var fileDescriptor_78985be6a51b316f = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2c, 0x4e,
	0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0xc9, 0xf9,
	0x39, 0xc5, 0x4a, 0x2e, 0x5c, 0x42, 0x61, 0xa9, 0x45, 0x99, 0x69, 0x95, 0x2e, 0x50, 0x25, 0x6e,
	0xf9, 0x45, 0xb9, 0x42, 0x62, 0x5c, 0x6c, 0xa5, 0xc5, 0xa9, 0x45, 0x99, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x1e, 0x48, 0xbc, 0xa0, 0x28, 0x3f, 0x25, 0x33, 0x45, 0x82, 0x09,
	0x22, 0x0e, 0xe1, 0x29, 0xb9, 0x71, 0xf1, 0xc0, 0xf4, 0x7b, 0xe6, 0xa5, 0xe5, 0x0b, 0x09, 0x70,
	0x31, 0x17, 0x24, 0x97, 0x80, 0x35, 0x33, 0x05, 0x81, 0x98, 0x42, 0x2a, 0x5c, 0x7c, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0xf1, 0x99, 0x79, 0xf1, 0xc9, 0xa9, 0x79, 0x25, 0xc5, 0x60, 0x13, 0x98, 0x83,
	0x78, 0xc0, 0xa2, 0x9e, 0x79, 0xce, 0x20, 0x31, 0xa3, 0x30, 0x2e, 0x0e, 0x98, 0x39, 0x42, 0x5e,
	0x5c, 0x7c, 0xa8, 0x2e, 0x13, 0x92, 0xd5, 0x83, 0xbb, 0x5b, 0x0f, 0xd3, 0xd1, 0x52, 0xe2, 0x48,
	0xd2, 0xc8, 0xae, 0x51, 0x62, 0x48, 0x62, 0x03, 0xcb, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x29, 0x7d, 0x2e, 0x3d, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscountClient is the client API for Discount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscountClient interface {
	VerifyDiscount(ctx context.Context, in *VerifyDiscountForm, opts ...grpc.CallOption) (*DiscountInfo, error)
}

type discountClient struct {
	cc *grpc.ClientConn
}

func NewDiscountClient(cc *grpc.ClientConn) DiscountClient {
	return &discountClient{cc}
}

func (c *discountClient) VerifyDiscount(ctx context.Context, in *VerifyDiscountForm, opts ...grpc.CallOption) (*DiscountInfo, error) {
	out := new(DiscountInfo)
	err := c.cc.Invoke(ctx, "/protocols.Discount/VerifyDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountServer is the server API for Discount service.
type DiscountServer interface {
	VerifyDiscount(context.Context, *VerifyDiscountForm) (*DiscountInfo, error)
}

// UnimplementedDiscountServer can be embedded to have forward compatible implementations.
type UnimplementedDiscountServer struct {
}

func (*UnimplementedDiscountServer) VerifyDiscount(ctx context.Context, req *VerifyDiscountForm) (*DiscountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyDiscount not implemented")
}

func RegisterDiscountServer(s *grpc.Server, srv DiscountServer) {
	s.RegisterService(&_Discount_serviceDesc, srv)
}

func _Discount_VerifyDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyDiscountForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServer).VerifyDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.Discount/VerifyDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServer).VerifyDiscount(ctx, req.(*VerifyDiscountForm))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discount_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocols.Discount",
	HandlerType: (*DiscountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyDiscount",
			Handler:    _Discount_VerifyDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discount.proto",
}
