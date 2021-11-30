// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: intertx/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/ibc-go/v2/modules/apps/27-interchain-accounts/types"
	_ "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryIBCAccountFromAddressRequest struct {
	Address                  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	ConnectionId             string                                        `protobuf:"bytes,2,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	CounterpartyConnectionId string                                        `protobuf:"bytes,3,opt,name=counterpartyConnectionId,proto3" json:"counterpartyConnectionId,omitempty"`
}

func (m *QueryIBCAccountFromAddressRequest) Reset()         { *m = QueryIBCAccountFromAddressRequest{} }
func (m *QueryIBCAccountFromAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountFromAddressRequest) ProtoMessage()    {}
func (*QueryIBCAccountFromAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b2be0eec8e13d, []int{0}
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountFromAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountFromAddressRequest.Merge(m, src)
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountFromAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountFromAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountFromAddressRequest proto.InternalMessageInfo

type QueryIBCAccountFromAddressResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryIBCAccountFromAddressResponse) Reset()         { *m = QueryIBCAccountFromAddressResponse{} }
func (m *QueryIBCAccountFromAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountFromAddressResponse) ProtoMessage()    {}
func (*QueryIBCAccountFromAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b2be0eec8e13d, []int{1}
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountFromAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountFromAddressResponse.Merge(m, src)
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountFromAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountFromAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountFromAddressResponse proto.InternalMessageInfo

func (m *QueryIBCAccountFromAddressResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryIBCAccountFromAddressRequest)(nil), "intertx.QueryIBCAccountFromAddressRequest")
	proto.RegisterType((*QueryIBCAccountFromAddressResponse)(nil), "intertx.QueryIBCAccountFromAddressResponse")
}

func init() { proto.RegisterFile("intertx/query.proto", fileDescriptor_5c0b2be0eec8e13d) }

var fileDescriptor_5c0b2be0eec8e13d = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0x4e, 0x14, 0xad, 0x1d, 0x7a, 0x8a, 0x0a, 0xa1, 0x48, 0x52, 0x73, 0x2a, 0x4a, 0x32, 0x54,
	0xd1, 0x43, 0x0f, 0x42, 0x5b, 0x10, 0x8a, 0x17, 0xed, 0xd1, 0x8b, 0x4c, 0x27, 0x43, 0x3a, 0xd8,
	0xce, 0x3b, 0x9d, 0x99, 0x48, 0xf3, 0x0f, 0x3c, 0xfa, 0x13, 0xfa, 0x73, 0x3c, 0xf6, 0x28, 0x1e,
	0x96, 0xa5, 0xbd, 0xec, 0x6f, 0xd8, 0xd3, 0x92, 0x4c, 0xda, 0xdd, 0xb2, 0x9f, 0xa7, 0xbc, 0x3c,
	0x1f, 0x6f, 0x9e, 0x79, 0x1f, 0xf4, 0x9c, 0x0b, 0xc3, 0x94, 0x59, 0xe1, 0x65, 0xce, 0x54, 0x91,
	0x48, 0x05, 0x06, 0xbc, 0x46, 0x0d, 0xb6, 0x5f, 0x64, 0x90, 0x41, 0x85, 0xe1, 0x72, 0xb2, 0x74,
	0xfb, 0x55, 0x06, 0x90, 0xcd, 0x19, 0x26, 0x92, 0x63, 0x22, 0x04, 0x18, 0x62, 0x38, 0x08, 0x5d,
	0xb3, 0x21, 0x9f, 0x52, 0x4c, 0x41, 0x31, 0x4c, 0xe7, 0x9c, 0x09, 0x83, 0x7f, 0xf5, 0xea, 0xa9,
	0x16, 0x7c, 0x28, 0x05, 0x44, 0xca, 0x39, 0xa7, 0xd6, 0x88, 0xab, 0xdf, 0xd1, 0x19, 0xe1, 0xe2,
	0x07, 0xa1, 0x14, 0x72, 0x61, 0x74, 0xe9, 0xaa, 0x67, 0x6b, 0x8b, 0xfe, 0xbb, 0xe8, 0xf5, 0xb7,
	0x32, 0xe4, 0x78, 0x38, 0x1a, 0x58, 0xe6, 0xb3, 0x82, 0xc5, 0x20, 0x4d, 0x15, 0xd3, 0x7a, 0xc2,
	0x96, 0x39, 0xd3, 0xc6, 0xfb, 0x82, 0x1a, 0xc4, 0x22, 0xbe, 0xdb, 0x71, 0xbb, 0xad, 0x61, 0xef,
	0xfc, 0x24, 0x8c, 0x33, 0x6e, 0x66, 0xf9, 0x34, 0xa1, 0xb0, 0xc0, 0x14, 0xf4, 0x02, 0x74, 0xfd,
	0x89, 0x75, 0xfa, 0x13, 0x9b, 0x42, 0x32, 0x9d, 0x0c, 0x28, 0xdd, 0xaf, 0xda, 0x6f, 0xf0, 0x22,
	0xd4, 0xa2, 0x20, 0x04, 0xa3, 0x65, 0xcc, 0x71, 0xea, 0x3f, 0xea, 0xb8, 0xdd, 0xe6, 0xe4, 0x08,
	0xf3, 0xfa, 0xc8, 0xaf, 0xb2, 0x30, 0x25, 0x89, 0x32, 0xc5, 0xe8, 0xaa, 0xfe, 0x71, 0xa5, 0xbf,
	0x95, 0xef, 0x3f, 0xfb, 0xbd, 0x0e, 0x9d, 0xb3, 0x75, 0xe8, 0x44, 0x9f, 0x50, 0x74, 0xd7, 0xdb,
	0xb4, 0x04, 0xa1, 0x99, 0xe7, 0x1f, 0x3f, 0xae, 0x79, 0x48, 0xfa, 0xae, 0x40, 0x4f, 0x2a, 0xbf,
	0x27, 0xd1, 0xcb, 0x1b, 0x77, 0x78, 0x6f, 0x92, 0xba, 0xd4, 0xe4, 0xde, 0x23, 0xb6, 0xdf, 0x3e,
	0x48, 0x6b, 0x43, 0x0d, 0xbf, 0xfe, 0xdd, 0x06, 0xee, 0x66, 0x1b, 0xb8, 0xa7, 0xdb, 0xc0, 0xfd,
	0xb3, 0x0b, 0x9c, 0xcd, 0x2e, 0x70, 0xfe, 0xed, 0x02, 0xe7, 0xfb, 0xc7, 0xeb, 0x67, 0xbf, 0x6c,
	0x3a, 0x3e, 0x34, 0xbd, 0xb2, 0x68, 0x6c, 0x56, 0xb6, 0x8a, 0xe9, 0xd3, 0xaa, 0xf0, 0xf7, 0x17,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x48, 0x78, 0xfa, 0x9c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	IBCAccountFromAddress(ctx context.Context, in *QueryIBCAccountFromAddressRequest, opts ...grpc.CallOption) (*QueryIBCAccountFromAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IBCAccountFromAddress(ctx context.Context, in *QueryIBCAccountFromAddressRequest, opts ...grpc.CallOption) (*QueryIBCAccountFromAddressResponse, error) {
	out := new(QueryIBCAccountFromAddressResponse)
	err := c.cc.Invoke(ctx, "/intertx.Query/IBCAccountFromAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	IBCAccountFromAddress(context.Context, *QueryIBCAccountFromAddressRequest) (*QueryIBCAccountFromAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IBCAccountFromAddress(ctx context.Context, req *QueryIBCAccountFromAddressRequest) (*QueryIBCAccountFromAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCAccountFromAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IBCAccountFromAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIBCAccountFromAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IBCAccountFromAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intertx.Query/IBCAccountFromAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IBCAccountFromAddress(ctx, req.(*QueryIBCAccountFromAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "intertx.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IBCAccountFromAddress",
			Handler:    _Query_IBCAccountFromAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intertx/query.proto",
}

func (m *QueryIBCAccountFromAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountFromAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountFromAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CounterpartyConnectionId) > 0 {
		i -= len(m.CounterpartyConnectionId)
		copy(dAtA[i:], m.CounterpartyConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CounterpartyConnectionId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIBCAccountFromAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountFromAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountFromAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryIBCAccountFromAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.CounterpartyConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIBCAccountFromAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryIBCAccountFromAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIBCAccountFromAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountFromAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryIBCAccountFromAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIBCAccountFromAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountFromAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
