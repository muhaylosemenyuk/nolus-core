// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nolus/suspend/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSuspend struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	BlockHeight int64  `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *MsgSuspend) Reset()         { *m = MsgSuspend{} }
func (m *MsgSuspend) String() string { return proto.CompactTextString(m) }
func (*MsgSuspend) ProtoMessage()    {}
func (*MsgSuspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda7a982fa57e103, []int{0}
}
func (m *MsgSuspend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSuspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSuspend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSuspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSuspend.Merge(m, src)
}
func (m *MsgSuspend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSuspend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSuspend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSuspend proto.InternalMessageInfo

func (m *MsgSuspend) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSuspend) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type MsgUnsuspend struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgUnsuspend) Reset()         { *m = MsgUnsuspend{} }
func (m *MsgUnsuspend) String() string { return proto.CompactTextString(m) }
func (*MsgUnsuspend) ProtoMessage()    {}
func (*MsgUnsuspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda7a982fa57e103, []int{1}
}
func (m *MsgUnsuspend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnsuspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnsuspend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnsuspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnsuspend.Merge(m, src)
}
func (m *MsgUnsuspend) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnsuspend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnsuspend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnsuspend proto.InternalMessageInfo

func (m *MsgUnsuspend) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

type MsgSuspendResponse struct {
}

func (m *MsgSuspendResponse) Reset()         { *m = MsgSuspendResponse{} }
func (m *MsgSuspendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSuspendResponse) ProtoMessage()    {}
func (*MsgSuspendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda7a982fa57e103, []int{2}
}
func (m *MsgSuspendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSuspendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSuspendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSuspendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSuspendResponse.Merge(m, src)
}
func (m *MsgSuspendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSuspendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSuspendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSuspendResponse proto.InternalMessageInfo

type MsgUnsuspendResponse struct {
}

func (m *MsgUnsuspendResponse) Reset()         { *m = MsgUnsuspendResponse{} }
func (m *MsgUnsuspendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnsuspendResponse) ProtoMessage()    {}
func (*MsgUnsuspendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eda7a982fa57e103, []int{3}
}
func (m *MsgUnsuspendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnsuspendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnsuspendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnsuspendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnsuspendResponse.Merge(m, src)
}
func (m *MsgUnsuspendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnsuspendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnsuspendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnsuspendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSuspend)(nil), "nolus.suspend.v1beta1.MsgSuspend")
	proto.RegisterType((*MsgUnsuspend)(nil), "nolus.suspend.v1beta1.MsgUnsuspend")
	proto.RegisterType((*MsgSuspendResponse)(nil), "nolus.suspend.v1beta1.MsgSuspendResponse")
	proto.RegisterType((*MsgUnsuspendResponse)(nil), "nolus.suspend.v1beta1.MsgUnsuspendResponse")
}

func init() { proto.RegisterFile("nolus/suspend/v1beta1/tx.proto", fileDescriptor_eda7a982fa57e103) }

var fileDescriptor_eda7a982fa57e103 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcb, 0xcf, 0x29,
	0x2d, 0xd6, 0x2f, 0x2e, 0x2d, 0x2e, 0x48, 0xcd, 0x4b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x05, 0xcb, 0xeb, 0x41,
	0xe5, 0xf5, 0xa0, 0xf2, 0x4a, 0x41, 0x5c, 0x5c, 0xbe, 0xc5, 0xe9, 0xc1, 0x10, 0x51, 0x21, 0x45,
	0x2e, 0x9e, 0xb4, 0xa2, 0xfc, 0xdc, 0xf8, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x6e, 0x90, 0x98, 0x23, 0x44, 0x08, 0xa4, 0x24, 0x29, 0x27, 0x3f,
	0x39, 0x3b, 0x3e, 0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x88,
	0x1b, 0x2c, 0xe6, 0x01, 0x16, 0x52, 0x32, 0xe4, 0xe2, 0xf1, 0x2d, 0x4e, 0x0f, 0xcd, 0x2b, 0x26,
	0xda, 0x54, 0x25, 0x11, 0x2e, 0x21, 0x84, 0x33, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x95, 0xc4, 0xb8, 0x44, 0x90, 0x0d, 0x82, 0x89, 0x1b, 0xed, 0x65, 0xe4, 0x62, 0xf6, 0x2d, 0x4e,
	0x17, 0x0a, 0xe7, 0x62, 0x87, 0xbb, 0x5c, 0x0f, 0xab, 0xff, 0xf4, 0x10, 0xa6, 0x4a, 0x69, 0x12,
	0x54, 0x02, 0xb3, 0x40, 0x28, 0x96, 0x8b, 0x13, 0xe1, 0x7c, 0x65, 0xdc, 0xfa, 0xe0, 0x8a, 0xa4,
	0xb4, 0x89, 0x50, 0x04, 0x33, 0xde, 0x29, 0xe8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0x2c, 0xd2, 0x33, 0x4b, 0x72, 0x12, 0x93, 0x74, 0xf3, 0xf2, 0x73, 0xf3, 0xf5, 0x92,
	0x8b, 0x52, 0x53, 0x32, 0x8b, 0x8b, 0x33, 0x73, 0xf3, 0xf5, 0xf2, 0x52, 0x4b, 0xf4, 0x41, 0x62,
	0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0x55, 0xf9, 0x79, 0xa9, 0xfa, 0x15, 0xf0, 0xa8, 0x2e, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x47, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xda, 0x72,
	0x90, 0x08, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Suspend(ctx context.Context, in *MsgSuspend, opts ...grpc.CallOption) (*MsgSuspendResponse, error)
	Unsuspend(ctx context.Context, in *MsgUnsuspend, opts ...grpc.CallOption) (*MsgUnsuspendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Suspend(ctx context.Context, in *MsgSuspend, opts ...grpc.CallOption) (*MsgSuspendResponse, error) {
	out := new(MsgSuspendResponse)
	err := c.cc.Invoke(ctx, "/nolus.suspend.v1beta1.Msg/Suspend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Unsuspend(ctx context.Context, in *MsgUnsuspend, opts ...grpc.CallOption) (*MsgUnsuspendResponse, error) {
	out := new(MsgUnsuspendResponse)
	err := c.cc.Invoke(ctx, "/nolus.suspend.v1beta1.Msg/Unsuspend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Suspend(context.Context, *MsgSuspend) (*MsgSuspendResponse, error)
	Unsuspend(context.Context, *MsgUnsuspend) (*MsgUnsuspendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Suspend(ctx context.Context, req *MsgSuspend) (*MsgSuspendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suspend not implemented")
}
func (*UnimplementedMsgServer) Unsuspend(ctx context.Context, req *MsgUnsuspend) (*MsgUnsuspendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsuspend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Suspend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSuspend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Suspend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nolus.suspend.v1beta1.Msg/Suspend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Suspend(ctx, req.(*MsgSuspend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Unsuspend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnsuspend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unsuspend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nolus.suspend.v1beta1.Msg/Unsuspend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unsuspend(ctx, req.(*MsgUnsuspend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nolus.suspend.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Suspend",
			Handler:    _Msg_Suspend_Handler,
		},
		{
			MethodName: "Unsuspend",
			Handler:    _Msg_Unsuspend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nolus/suspend/v1beta1/tx.proto",
}

func (m *MsgSuspend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSuspend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSuspend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnsuspend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnsuspend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnsuspend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSuspendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSuspendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSuspendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnsuspendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnsuspendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnsuspendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSuspend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTx(uint64(m.BlockHeight))
	}
	return n
}

func (m *MsgUnsuspend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSuspendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnsuspendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSuspend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSuspend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSuspend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnsuspend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnsuspend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnsuspend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSuspendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSuspendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSuspendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnsuspendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnsuspendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnsuspendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
