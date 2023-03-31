// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/globalfee/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryCodeAuthorizationRequest struct {
	CodeId uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
}

func (m *QueryCodeAuthorizationRequest) Reset()         { *m = QueryCodeAuthorizationRequest{} }
func (m *QueryCodeAuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCodeAuthorizationRequest) ProtoMessage()    {}
func (*QueryCodeAuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d932d0017ba0ee, []int{0}
}
func (m *QueryCodeAuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCodeAuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodeAuthorizationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCodeAuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodeAuthorizationRequest.Merge(m, src)
}
func (m *QueryCodeAuthorizationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCodeAuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodeAuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodeAuthorizationRequest proto.InternalMessageInfo

func (m *QueryCodeAuthorizationRequest) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

type QueryCodeAuthorizationResponse struct {
	Methods []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (m *QueryCodeAuthorizationResponse) Reset()         { *m = QueryCodeAuthorizationResponse{} }
func (m *QueryCodeAuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCodeAuthorizationResponse) ProtoMessage()    {}
func (*QueryCodeAuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d932d0017ba0ee, []int{1}
}
func (m *QueryCodeAuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCodeAuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodeAuthorizationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCodeAuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodeAuthorizationResponse.Merge(m, src)
}
func (m *QueryCodeAuthorizationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCodeAuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodeAuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodeAuthorizationResponse proto.InternalMessageInfo

func (m *QueryCodeAuthorizationResponse) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

type QueryContractAuthorizationRequest struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *QueryContractAuthorizationRequest) Reset()         { *m = QueryContractAuthorizationRequest{} }
func (m *QueryContractAuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractAuthorizationRequest) ProtoMessage()    {}
func (*QueryContractAuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d932d0017ba0ee, []int{2}
}
func (m *QueryContractAuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractAuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractAuthorizationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractAuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractAuthorizationRequest.Merge(m, src)
}
func (m *QueryContractAuthorizationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractAuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractAuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractAuthorizationRequest proto.InternalMessageInfo

func (m *QueryContractAuthorizationRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

type QueryContractAuthorizationResponse struct {
	Methods []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (m *QueryContractAuthorizationResponse) Reset()         { *m = QueryContractAuthorizationResponse{} }
func (m *QueryContractAuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractAuthorizationResponse) ProtoMessage()    {}
func (*QueryContractAuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d932d0017ba0ee, []int{3}
}
func (m *QueryContractAuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractAuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractAuthorizationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractAuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractAuthorizationResponse.Merge(m, src)
}
func (m *QueryContractAuthorizationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractAuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractAuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractAuthorizationResponse proto.InternalMessageInfo

func (m *QueryContractAuthorizationResponse) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCodeAuthorizationRequest)(nil), "publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest")
	proto.RegisterType((*QueryCodeAuthorizationResponse)(nil), "publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse")
	proto.RegisterType((*QueryContractAuthorizationRequest)(nil), "publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest")
	proto.RegisterType((*QueryContractAuthorizationResponse)(nil), "publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse")
}

func init() { proto.RegisterFile("stargaze/globalfee/v1/query.proto", fileDescriptor_47d932d0017ba0ee) }

var fileDescriptor_47d932d0017ba0ee = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2e, 0x49, 0x2c,
	0x4a, 0x4f, 0xac, 0x4a, 0xd5, 0x4f, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x49, 0x4b, 0x4d, 0xd5, 0x2f,
	0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e,
	0x28, 0x4d, 0xca, 0xc9, 0x4c, 0x4e, 0x2c, 0x4f, 0x2d, 0xce, 0xcf, 0x4d, 0xd5, 0x83, 0x69, 0xd0,
	0x83, 0x6b, 0xd0, 0x2b, 0x33, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c,
	0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x18, 0xa1,
	0x64, 0xc1, 0x25, 0x1b, 0x08, 0x32, 0xd1, 0x39, 0x3f, 0x25, 0xd5, 0xb1, 0xb4, 0x24, 0x23, 0xbf,
	0x28, 0xb3, 0x0a, 0xac, 0x20, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x9c, 0x8b, 0x3d,
	0x39, 0x3f, 0x25, 0x35, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x0d, 0xc4,
	0xf5, 0x4c, 0x51, 0xb2, 0xe2, 0x92, 0xc3, 0xa5, 0xb3, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48,
	0x82, 0x8b, 0x3d, 0x37, 0xb5, 0x24, 0x23, 0x3f, 0xa5, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0x33,
	0x08, 0xc6, 0x55, 0xf2, 0xe3, 0x52, 0x84, 0xea, 0xcd, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xc1, 0x6a,
	0xb3, 0x26, 0x97, 0x40, 0x32, 0x54, 0x3e, 0x3e, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x18, 0xec,
	0x04, 0xce, 0x20, 0x7e, 0x98, 0xb8, 0x23, 0x44, 0x58, 0xc9, 0x8e, 0x4b, 0x09, 0x9f, 0x79, 0x84,
	0xdc, 0x63, 0x74, 0x83, 0x99, 0x8b, 0x15, 0x6c, 0x80, 0xd0, 0x4d, 0x46, 0x2e, 0x41, 0x0c, 0x1f,
	0x09, 0x39, 0xe9, 0x11, 0x11, 0xd2, 0x7a, 0x78, 0x03, 0x52, 0xca, 0x99, 0x22, 0x33, 0x20, 0x5e,
	0x50, 0xb2, 0x6e, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xa9, 0x90, 0xb1, 0x3e, 0xf6, 0xd4, 0x01, 0x8e,
	0xaa, 0x44, 0x64, 0xad, 0xfa, 0xd5, 0xd0, 0xe8, 0xab, 0x15, 0xfa, 0xc4, 0xc8, 0x25, 0x8a, 0x35,
	0x84, 0x84, 0xdc, 0x48, 0x71, 0x1b, 0xee, 0x28, 0x93, 0x72, 0xa7, 0xd8, 0x1c, 0xa8, 0x3f, 0x3d,
	0xc0, 0xfe, 0x74, 0x12, 0x72, 0xc0, 0xe9, 0x4f, 0x58, 0xc2, 0x40, 0xf7, 0x2b, 0x6a, 0x82, 0xa9,
	0x75, 0x0a, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xf3, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0xb3, 0x75, 0xa1, 0xee, 0x46, 0x58,
	0x5a, 0x66, 0xa9, 0x5f, 0x81, 0x64, 0x73, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xeb,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30, 0x41, 0x4c, 0x3b, 0xa2, 0x03, 0x00, 0x00,
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
	CodeAuthorization(ctx context.Context, in *QueryCodeAuthorizationRequest, opts ...grpc.CallOption) (*QueryCodeAuthorizationResponse, error)
	ContractAuthorization(ctx context.Context, in *QueryContractAuthorizationRequest, opts ...grpc.CallOption) (*QueryContractAuthorizationResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CodeAuthorization(ctx context.Context, in *QueryCodeAuthorizationRequest, opts ...grpc.CallOption) (*QueryCodeAuthorizationResponse, error) {
	out := new(QueryCodeAuthorizationResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.globalfee.v1.Query/CodeAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractAuthorization(ctx context.Context, in *QueryContractAuthorizationRequest, opts ...grpc.CallOption) (*QueryContractAuthorizationResponse, error) {
	out := new(QueryContractAuthorizationResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.globalfee.v1.Query/ContractAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	CodeAuthorization(context.Context, *QueryCodeAuthorizationRequest) (*QueryCodeAuthorizationResponse, error)
	ContractAuthorization(context.Context, *QueryContractAuthorizationRequest) (*QueryContractAuthorizationResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CodeAuthorization(ctx context.Context, req *QueryCodeAuthorizationRequest) (*QueryCodeAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeAuthorization not implemented")
}
func (*UnimplementedQueryServer) ContractAuthorization(ctx context.Context, req *QueryContractAuthorizationRequest) (*QueryContractAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractAuthorization not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CodeAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCodeAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CodeAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.globalfee.v1.Query/CodeAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CodeAuthorization(ctx, req.(*QueryCodeAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.globalfee.v1.Query/ContractAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractAuthorization(ctx, req.(*QueryContractAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.globalfee.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CodeAuthorization",
			Handler:    _Query_CodeAuthorization_Handler,
		},
		{
			MethodName: "ContractAuthorization",
			Handler:    _Query_ContractAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargaze/globalfee/v1/query.proto",
}

func (m *QueryCodeAuthorizationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodeAuthorizationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodeAuthorizationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodeAuthorizationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodeAuthorizationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodeAuthorizationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Methods[iNdEx])
			copy(dAtA[i:], m.Methods[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Methods[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractAuthorizationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractAuthorizationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractAuthorizationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractAuthorizationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractAuthorizationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractAuthorizationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Methods[iNdEx])
			copy(dAtA[i:], m.Methods[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Methods[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryCodeAuthorizationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovQuery(uint64(m.CodeId))
	}
	return n
}

func (m *QueryCodeAuthorizationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for _, s := range m.Methods {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryContractAuthorizationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractAuthorizationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for _, s := range m.Methods {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCodeAuthorizationRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCodeAuthorizationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodeAuthorizationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryCodeAuthorizationResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCodeAuthorizationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodeAuthorizationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
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
			m.Methods = append(m.Methods, string(dAtA[iNdEx:postIndex]))
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
func (m *QueryContractAuthorizationRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractAuthorizationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractAuthorizationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryContractAuthorizationResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractAuthorizationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractAuthorizationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
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
			m.Methods = append(m.Methods, string(dAtA[iNdEx:postIndex]))
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
