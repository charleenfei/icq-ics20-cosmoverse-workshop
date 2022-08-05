// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/v1/tx.proto

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

type MsgFeelingLucky struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Offering string `protobuf:"bytes,2,opt,name=offering,proto3" json:"offering,omitempty"`
}

func (m *MsgFeelingLucky) Reset()         { *m = MsgFeelingLucky{} }
func (m *MsgFeelingLucky) String() string { return proto.CompactTextString(m) }
func (*MsgFeelingLucky) ProtoMessage()    {}
func (*MsgFeelingLucky) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6ea5d56f153e84, []int{0}
}
func (m *MsgFeelingLucky) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFeelingLucky) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFeelingLucky.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFeelingLucky) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFeelingLucky.Merge(m, src)
}
func (m *MsgFeelingLucky) XXX_Size() int {
	return m.Size()
}
func (m *MsgFeelingLucky) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFeelingLucky.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFeelingLucky proto.InternalMessageInfo

func (m *MsgFeelingLucky) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgFeelingLucky) GetOffering() string {
	if m != nil {
		return m.Offering
	}
	return ""
}

type MsgFeelingLuckyResponse struct {
}

func (m *MsgFeelingLuckyResponse) Reset()         { *m = MsgFeelingLuckyResponse{} }
func (m *MsgFeelingLuckyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFeelingLuckyResponse) ProtoMessage()    {}
func (*MsgFeelingLuckyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6ea5d56f153e84, []int{1}
}
func (m *MsgFeelingLuckyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFeelingLuckyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFeelingLuckyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFeelingLuckyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFeelingLuckyResponse.Merge(m, src)
}
func (m *MsgFeelingLuckyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFeelingLuckyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFeelingLuckyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFeelingLuckyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFeelingLucky)(nil), "eightball.v1.MsgFeelingLucky")
	proto.RegisterType((*MsgFeelingLuckyResponse)(nil), "eightball.v1.MsgFeelingLuckyResponse")
}

func init() { proto.RegisterFile("eightball/v1/tx.proto", fileDescriptor_0c6ea5d56f153e84) }

var fileDescriptor_0c6ea5d56f153e84 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x81, 0x0b, 0xeb, 0x95, 0x19, 0x2a, 0xb9, 0x73, 0xf1, 0xfb, 0x16, 0xa7, 0xbb,
	0xa5, 0xa6, 0xe6, 0x64, 0xe6, 0xa5, 0xfb, 0x94, 0x26, 0x67, 0x57, 0x0a, 0x49, 0x70, 0xb1, 0x27,
	0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42,
	0x52, 0x5c, 0x1c, 0xf9, 0x69, 0x69, 0xa9, 0x45, 0x99, 0x79, 0xe9, 0x12, 0x4c, 0x60, 0x29, 0x38,
	0x5f, 0x49, 0x92, 0x4b, 0x1c, 0xcd, 0xa0, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3,
	0x68, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x10, 0x2e, 0x1e, 0x14, 0x7b, 0x64, 0xf5, 0x90, 0x5d,
	0xa2, 0x87, 0xa6, 0x5b, 0x4a, 0x15, 0xaf, 0x34, 0xcc, 0x70, 0xa7, 0xb0, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x4f, 0xce, 0x48, 0x2c, 0xca, 0x49, 0x4d, 0xcd, 0x4b, 0x4b, 0xcd, 0xd4, 0x4f, 0xce,
	0x2f, 0xce, 0xcd, 0x2f, 0x4b, 0x2d, 0x2a, 0x4e, 0xd5, 0x2d, 0xcf, 0x2f, 0xca, 0x2e, 0xce, 0xc8,
	0x2f, 0xd0, 0xaf, 0xd0, 0x47, 0x04, 0x56, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xb4,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x62, 0xd2, 0x6c, 0x46, 0x01, 0x00, 0x00,
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
	FeelingLucky(ctx context.Context, in *MsgFeelingLucky, opts ...grpc.CallOption) (*MsgFeelingLuckyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FeelingLucky(ctx context.Context, in *MsgFeelingLucky, opts ...grpc.CallOption) (*MsgFeelingLuckyResponse, error) {
	out := new(MsgFeelingLuckyResponse)
	err := c.cc.Invoke(ctx, "/eightball.v1.Msg/FeelingLucky", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	FeelingLucky(context.Context, *MsgFeelingLucky) (*MsgFeelingLuckyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) FeelingLucky(ctx context.Context, req *MsgFeelingLucky) (*MsgFeelingLuckyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeelingLucky not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_FeelingLucky_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFeelingLucky)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FeelingLucky(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eightball.v1.Msg/FeelingLucky",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FeelingLucky(ctx, req.(*MsgFeelingLucky))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eightball.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FeelingLucky",
			Handler:    _Msg_FeelingLucky_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eightball/v1/tx.proto",
}

func (m *MsgFeelingLucky) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFeelingLucky) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFeelingLucky) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Offering) > 0 {
		i -= len(m.Offering)
		copy(dAtA[i:], m.Offering)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Offering)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFeelingLuckyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFeelingLuckyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFeelingLuckyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgFeelingLucky) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Offering)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFeelingLuckyResponse) Size() (n int) {
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
func (m *MsgFeelingLucky) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFeelingLucky: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFeelingLucky: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offering", wireType)
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
			m.Offering = string(dAtA[iNdEx:postIndex])
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
func (m *MsgFeelingLuckyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFeelingLuckyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFeelingLuckyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
