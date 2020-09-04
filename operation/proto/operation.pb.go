// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operation/proto/operation.proto

package proto

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

type Operation struct {
	Number1              float64  `protobuf:"fixed64,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              float64  `protobuf:"fixed64,2,opt,name=number2,proto3" json:"number2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetNumber1() float64 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Operation) GetNumber2() float64 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type SumRequest struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{1}
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

func (m *SumRequest) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type SumResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{2}
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

func (m *SumResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MultiplicationRequest struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MultiplicationRequest) Reset()         { *m = MultiplicationRequest{} }
func (m *MultiplicationRequest) String() string { return proto.CompactTextString(m) }
func (*MultiplicationRequest) ProtoMessage()    {}
func (*MultiplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{3}
}

func (m *MultiplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplicationRequest.Unmarshal(m, b)
}
func (m *MultiplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplicationRequest.Marshal(b, m, deterministic)
}
func (m *MultiplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplicationRequest.Merge(m, src)
}
func (m *MultiplicationRequest) XXX_Size() int {
	return xxx_messageInfo_MultiplicationRequest.Size(m)
}
func (m *MultiplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplicationRequest proto.InternalMessageInfo

func (m *MultiplicationRequest) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type MultiplicationResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplicationResponse) Reset()         { *m = MultiplicationResponse{} }
func (m *MultiplicationResponse) String() string { return proto.CompactTextString(m) }
func (*MultiplicationResponse) ProtoMessage()    {}
func (*MultiplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{4}
}

func (m *MultiplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplicationResponse.Unmarshal(m, b)
}
func (m *MultiplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplicationResponse.Marshal(b, m, deterministic)
}
func (m *MultiplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplicationResponse.Merge(m, src)
}
func (m *MultiplicationResponse) XXX_Size() int {
	return xxx_messageInfo_MultiplicationResponse.Size(m)
}
func (m *MultiplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplicationResponse proto.InternalMessageInfo

func (m *MultiplicationResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type DivisionRequest struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DivisionRequest) Reset()         { *m = DivisionRequest{} }
func (m *DivisionRequest) String() string { return proto.CompactTextString(m) }
func (*DivisionRequest) ProtoMessage()    {}
func (*DivisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{5}
}

func (m *DivisionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivisionRequest.Unmarshal(m, b)
}
func (m *DivisionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivisionRequest.Marshal(b, m, deterministic)
}
func (m *DivisionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivisionRequest.Merge(m, src)
}
func (m *DivisionRequest) XXX_Size() int {
	return xxx_messageInfo_DivisionRequest.Size(m)
}
func (m *DivisionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DivisionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DivisionRequest proto.InternalMessageInfo

func (m *DivisionRequest) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type DivisionResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DivisionResponse) Reset()         { *m = DivisionResponse{} }
func (m *DivisionResponse) String() string { return proto.CompactTextString(m) }
func (*DivisionResponse) ProtoMessage()    {}
func (*DivisionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{6}
}

func (m *DivisionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivisionResponse.Unmarshal(m, b)
}
func (m *DivisionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivisionResponse.Marshal(b, m, deterministic)
}
func (m *DivisionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivisionResponse.Merge(m, src)
}
func (m *DivisionResponse) XXX_Size() int {
	return xxx_messageInfo_DivisionResponse.Size(m)
}
func (m *DivisionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DivisionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DivisionResponse proto.InternalMessageInfo

func (m *DivisionResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type SubtractionRequest struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SubtractionRequest) Reset()         { *m = SubtractionRequest{} }
func (m *SubtractionRequest) String() string { return proto.CompactTextString(m) }
func (*SubtractionRequest) ProtoMessage()    {}
func (*SubtractionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{7}
}

func (m *SubtractionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractionRequest.Unmarshal(m, b)
}
func (m *SubtractionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractionRequest.Marshal(b, m, deterministic)
}
func (m *SubtractionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractionRequest.Merge(m, src)
}
func (m *SubtractionRequest) XXX_Size() int {
	return xxx_messageInfo_SubtractionRequest.Size(m)
}
func (m *SubtractionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractionRequest proto.InternalMessageInfo

func (m *SubtractionRequest) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type SubtractionResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractionResponse) Reset()         { *m = SubtractionResponse{} }
func (m *SubtractionResponse) String() string { return proto.CompactTextString(m) }
func (*SubtractionResponse) ProtoMessage()    {}
func (*SubtractionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c3607c964f93ad, []int{8}
}

func (m *SubtractionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractionResponse.Unmarshal(m, b)
}
func (m *SubtractionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractionResponse.Marshal(b, m, deterministic)
}
func (m *SubtractionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractionResponse.Merge(m, src)
}
func (m *SubtractionResponse) XXX_Size() int {
	return xxx_messageInfo_SubtractionResponse.Size(m)
}
func (m *SubtractionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractionResponse proto.InternalMessageInfo

func (m *SubtractionResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Operation)(nil), "Operation")
	proto.RegisterType((*SumRequest)(nil), "SumRequest")
	proto.RegisterType((*SumResponse)(nil), "SumResponse")
	proto.RegisterType((*MultiplicationRequest)(nil), "MultiplicationRequest")
	proto.RegisterType((*MultiplicationResponse)(nil), "MultiplicationResponse")
	proto.RegisterType((*DivisionRequest)(nil), "DivisionRequest")
	proto.RegisterType((*DivisionResponse)(nil), "DivisionResponse")
	proto.RegisterType((*SubtractionRequest)(nil), "SubtractionRequest")
	proto.RegisterType((*SubtractionResponse)(nil), "SubtractionResponse")
}

func init() {
	proto.RegisterFile("operation/proto/operation.proto", fileDescriptor_07c3607c964f93ad)
}

var fileDescriptor_07c3607c964f93ad = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x62, 0x6b, 0x27, 0xa2, 0x71, 0xab, 0x31, 0xe4, 0x62, 0x09, 0x08, 0x45, 0x70,
	0xa2, 0x11, 0x44, 0x10, 0x94, 0x8a, 0x57, 0x11, 0x92, 0x9b, 0xb7, 0x26, 0xec, 0x61, 0xa1, 0xcd,
	0xc6, 0xfd, 0xd3, 0x0f, 0xec, 0x27, 0xb1, 0xae, 0xcd, 0xa6, 0xff, 0xb4, 0xd8, 0xd3, 0x32, 0x3b,
	0xcc, 0xfb, 0x31, 0xef, 0x0d, 0x9c, 0xf3, 0x8a, 0x8a, 0xa1, 0x62, 0xbc, 0x8c, 0x2b, 0xc1, 0x15,
	0x8f, 0x6d, 0x8d, 0xa6, 0x8e, 0x9e, 0xa0, 0xf3, 0x56, 0x7f, 0x91, 0x00, 0xda, 0xa5, 0x1e, 0xe7,
	0x54, 0xdc, 0x04, 0x4e, 0xcf, 0xe9, 0x3b, 0x69, 0x5d, 0x36, 0x9d, 0x24, 0xd8, 0x99, 0xef, 0x24,
	0xd1, 0x1d, 0x40, 0xa6, 0xc7, 0x29, 0xfd, 0xd0, 0x54, 0x2a, 0xd2, 0x87, 0x8e, 0x25, 0x18, 0x0d,
	0x37, 0x01, 0xb4, 0x80, 0xb4, 0x69, 0x46, 0x17, 0xe0, 0x9a, 0x39, 0x59, 0xf1, 0x52, 0x52, 0xe2,
	0x43, 0x4b, 0x50, 0xa9, 0x47, 0x6a, 0x46, 0x9e, 0x55, 0xd1, 0x00, 0x4e, 0x5f, 0xa7, 0x2f, 0xab,
	0x46, 0xac, 0xf8, 0xd1, 0xf8, 0x37, 0xe9, 0x1a, 0xfc, 0x65, 0x89, 0x0d, 0xd0, 0x07, 0x38, 0x7a,
	0x61, 0x13, 0x26, 0xb7, 0xc2, 0x5d, 0x82, 0xd7, 0x0c, 0x6f, 0x00, 0x3d, 0x02, 0xc9, 0x74, 0xae,
	0xc4, 0xb0, 0xd8, 0x6e, 0xb5, 0x2b, 0xe8, 0x2e, 0xcc, 0xff, 0x8d, 0x4b, 0x3e, 0x1d, 0xf0, 0xac,
	0x4e, 0x46, 0xc5, 0x84, 0x15, 0x94, 0xf4, 0x60, 0x77, 0x1a, 0x04, 0x71, 0xb1, 0x89, 0x31, 0x3c,
	0xc0, 0xf9, 0x6c, 0x06, 0x70, 0xb8, 0x68, 0x20, 0xf1, 0x71, 0x6d, 0x28, 0xe1, 0x19, 0xfe, 0xe2,
	0x74, 0x0c, 0xfb, 0xb5, 0x29, 0xc4, 0xc3, 0x25, 0x73, 0xc3, 0x63, 0x5c, 0x71, 0xec, 0xfe, 0xfb,
	0x3c, 0xec, 0x66, 0xa4, 0x8b, 0xab, 0x3e, 0x85, 0x27, 0xb8, 0x66, 0xf9, 0xe7, 0xf6, 0xfb, 0x9e,
	0x39, 0xed, 0xbc, 0x65, 0x9e, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x19, 0xbb, 0x36,
	0x04, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperationServiceClient is the client API for OperationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Multiplication(ctx context.Context, in *MultiplicationRequest, opts ...grpc.CallOption) (*MultiplicationResponse, error)
	Division(ctx context.Context, in *DivisionRequest, opts ...grpc.CallOption) (*DivisionResponse, error)
	Subtraction(ctx context.Context, in *SubtractionRequest, opts ...grpc.CallOption) (*SubtractionResponse, error)
}

type operationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationServiceClient(cc grpc.ClientConnInterface) OperationServiceClient {
	return &operationServiceClient{cc}
}

func (c *operationServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/OperationService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) Multiplication(ctx context.Context, in *MultiplicationRequest, opts ...grpc.CallOption) (*MultiplicationResponse, error) {
	out := new(MultiplicationResponse)
	err := c.cc.Invoke(ctx, "/OperationService/Multiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) Division(ctx context.Context, in *DivisionRequest, opts ...grpc.CallOption) (*DivisionResponse, error) {
	out := new(DivisionResponse)
	err := c.cc.Invoke(ctx, "/OperationService/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) Subtraction(ctx context.Context, in *SubtractionRequest, opts ...grpc.CallOption) (*SubtractionResponse, error) {
	out := new(SubtractionResponse)
	err := c.cc.Invoke(ctx, "/OperationService/Subtraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServiceServer is the server API for OperationService service.
type OperationServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Multiplication(context.Context, *MultiplicationRequest) (*MultiplicationResponse, error)
	Division(context.Context, *DivisionRequest) (*DivisionResponse, error)
	Subtraction(context.Context, *SubtractionRequest) (*SubtractionResponse, error)
}

// UnimplementedOperationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperationServiceServer struct {
}

func (*UnimplementedOperationServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedOperationServiceServer) Multiplication(ctx context.Context, req *MultiplicationRequest) (*MultiplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiplication not implemented")
}
func (*UnimplementedOperationServiceServer) Division(ctx context.Context, req *DivisionRequest) (*DivisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Division not implemented")
}
func (*UnimplementedOperationServiceServer) Subtraction(ctx context.Context, req *SubtractionRequest) (*SubtractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtraction not implemented")
}

func RegisterOperationServiceServer(s *grpc.Server, srv OperationServiceServer) {
	s.RegisterService(&_OperationService_serviceDesc, srv)
}

func _OperationService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OperationService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_Multiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Multiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OperationService/Multiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Multiplication(ctx, req.(*MultiplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OperationService/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Division(ctx, req.(*DivisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_Subtraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Subtraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OperationService/Subtraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Subtraction(ctx, req.(*SubtractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OperationService",
	HandlerType: (*OperationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _OperationService_Sum_Handler,
		},
		{
			MethodName: "Multiplication",
			Handler:    _OperationService_Multiplication_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _OperationService_Division_Handler,
		},
		{
			MethodName: "Subtraction",
			Handler:    _OperationService_Subtraction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operation/proto/operation.proto",
}
