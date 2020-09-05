// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/db-operation.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type DbOperation struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number1              float64              `protobuf:"fixed64,2,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              float64              `protobuf:"fixed64,3,opt,name=number2,proto3" json:"number2,omitempty"`
	Result               float64              `protobuf:"fixed64,4,opt,name=result,proto3" json:"result,omitempty"`
	Operation            string               `protobuf:"bytes,5,opt,name=operation,proto3" json:"operation,omitempty"`
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,6,opt,name=DateCreated,proto3" json:"DateCreated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DbOperation) Reset()         { *m = DbOperation{} }
func (m *DbOperation) String() string { return proto.CompactTextString(m) }
func (*DbOperation) ProtoMessage()    {}
func (*DbOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{0}
}

func (m *DbOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DbOperation.Unmarshal(m, b)
}
func (m *DbOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DbOperation.Marshal(b, m, deterministic)
}
func (m *DbOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DbOperation.Merge(m, src)
}
func (m *DbOperation) XXX_Size() int {
	return xxx_messageInfo_DbOperation.Size(m)
}
func (m *DbOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_DbOperation.DiscardUnknown(m)
}

var xxx_messageInfo_DbOperation proto.InternalMessageInfo

func (m *DbOperation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DbOperation) GetNumber1() float64 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *DbOperation) GetNumber2() float64 {
	if m != nil {
		return m.Number2
	}
	return 0
}

func (m *DbOperation) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *DbOperation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *DbOperation) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

type CreateOperationRequest struct {
	DbOperation          *DbOperation `protobuf:"bytes,1,opt,name=dbOperation,proto3" json:"dbOperation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOperationRequest) Reset()         { *m = CreateOperationRequest{} }
func (m *CreateOperationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOperationRequest) ProtoMessage()    {}
func (*CreateOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{1}
}

func (m *CreateOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOperationRequest.Unmarshal(m, b)
}
func (m *CreateOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOperationRequest.Marshal(b, m, deterministic)
}
func (m *CreateOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOperationRequest.Merge(m, src)
}
func (m *CreateOperationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOperationRequest.Size(m)
}
func (m *CreateOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOperationRequest proto.InternalMessageInfo

func (m *CreateOperationRequest) GetDbOperation() *DbOperation {
	if m != nil {
		return m.DbOperation
	}
	return nil
}

type CreateOperationResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOperationResponse) Reset()         { *m = CreateOperationResponse{} }
func (m *CreateOperationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOperationResponse) ProtoMessage()    {}
func (*CreateOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{2}
}

func (m *CreateOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOperationResponse.Unmarshal(m, b)
}
func (m *CreateOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOperationResponse.Marshal(b, m, deterministic)
}
func (m *CreateOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOperationResponse.Merge(m, src)
}
func (m *CreateOperationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOperationResponse.Size(m)
}
func (m *CreateOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOperationResponse proto.InternalMessageInfo

func (m *CreateOperationResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListOperationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationRequest) Reset()         { *m = ListOperationRequest{} }
func (m *ListOperationRequest) String() string { return proto.CompactTextString(m) }
func (*ListOperationRequest) ProtoMessage()    {}
func (*ListOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{3}
}

func (m *ListOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationRequest.Unmarshal(m, b)
}
func (m *ListOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationRequest.Marshal(b, m, deterministic)
}
func (m *ListOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationRequest.Merge(m, src)
}
func (m *ListOperationRequest) XXX_Size() int {
	return xxx_messageInfo_ListOperationRequest.Size(m)
}
func (m *ListOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationRequest proto.InternalMessageInfo

type ListOperationResponse struct {
	DbOperations         []*DbOperation `protobuf:"bytes,1,rep,name=dbOperations,proto3" json:"dbOperations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListOperationResponse) Reset()         { *m = ListOperationResponse{} }
func (m *ListOperationResponse) String() string { return proto.CompactTextString(m) }
func (*ListOperationResponse) ProtoMessage()    {}
func (*ListOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{4}
}

func (m *ListOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationResponse.Unmarshal(m, b)
}
func (m *ListOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationResponse.Marshal(b, m, deterministic)
}
func (m *ListOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationResponse.Merge(m, src)
}
func (m *ListOperationResponse) XXX_Size() int {
	return xxx_messageInfo_ListOperationResponse.Size(m)
}
func (m *ListOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationResponse proto.InternalMessageInfo

func (m *ListOperationResponse) GetDbOperations() []*DbOperation {
	if m != nil {
		return m.DbOperations
	}
	return nil
}

type DeleteOperationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOperationRequest) Reset()         { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()    {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{5}
}

func (m *DeleteOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOperationRequest.Unmarshal(m, b)
}
func (m *DeleteOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOperationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOperationRequest.Merge(m, src)
}
func (m *DeleteOperationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOperationRequest.Size(m)
}
func (m *DeleteOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOperationRequest proto.InternalMessageInfo

func (m *DeleteOperationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteOperationResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOperationResponse) Reset()         { *m = DeleteOperationResponse{} }
func (m *DeleteOperationResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteOperationResponse) ProtoMessage()    {}
func (*DeleteOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{6}
}

func (m *DeleteOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOperationResponse.Unmarshal(m, b)
}
func (m *DeleteOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOperationResponse.Marshal(b, m, deterministic)
}
func (m *DeleteOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOperationResponse.Merge(m, src)
}
func (m *DeleteOperationResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteOperationResponse.Size(m)
}
func (m *DeleteOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOperationResponse proto.InternalMessageInfo

func (m *DeleteOperationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ReadOperationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOperationRequest) Reset()         { *m = ReadOperationRequest{} }
func (m *ReadOperationRequest) String() string { return proto.CompactTextString(m) }
func (*ReadOperationRequest) ProtoMessage()    {}
func (*ReadOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{7}
}

func (m *ReadOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOperationRequest.Unmarshal(m, b)
}
func (m *ReadOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOperationRequest.Marshal(b, m, deterministic)
}
func (m *ReadOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOperationRequest.Merge(m, src)
}
func (m *ReadOperationRequest) XXX_Size() int {
	return xxx_messageInfo_ReadOperationRequest.Size(m)
}
func (m *ReadOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOperationRequest proto.InternalMessageInfo

func (m *ReadOperationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadOperationResponse struct {
	DbOperation          *DbOperation `protobuf:"bytes,1,opt,name=dbOperation,proto3" json:"dbOperation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadOperationResponse) Reset()         { *m = ReadOperationResponse{} }
func (m *ReadOperationResponse) String() string { return proto.CompactTextString(m) }
func (*ReadOperationResponse) ProtoMessage()    {}
func (*ReadOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1733e2bbcb01496d, []int{8}
}

func (m *ReadOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOperationResponse.Unmarshal(m, b)
}
func (m *ReadOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOperationResponse.Marshal(b, m, deterministic)
}
func (m *ReadOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOperationResponse.Merge(m, src)
}
func (m *ReadOperationResponse) XXX_Size() int {
	return xxx_messageInfo_ReadOperationResponse.Size(m)
}
func (m *ReadOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOperationResponse proto.InternalMessageInfo

func (m *ReadOperationResponse) GetDbOperation() *DbOperation {
	if m != nil {
		return m.DbOperation
	}
	return nil
}

func init() {
	proto.RegisterType((*DbOperation)(nil), "DbOperation")
	proto.RegisterType((*CreateOperationRequest)(nil), "CreateOperationRequest")
	proto.RegisterType((*CreateOperationResponse)(nil), "CreateOperationResponse")
	proto.RegisterType((*ListOperationRequest)(nil), "ListOperationRequest")
	proto.RegisterType((*ListOperationResponse)(nil), "ListOperationResponse")
	proto.RegisterType((*DeleteOperationRequest)(nil), "DeleteOperationRequest")
	proto.RegisterType((*DeleteOperationResponse)(nil), "DeleteOperationResponse")
	proto.RegisterType((*ReadOperationRequest)(nil), "ReadOperationRequest")
	proto.RegisterType((*ReadOperationResponse)(nil), "ReadOperationResponse")
}

func init() {
	proto.RegisterFile("proto/db-operation.proto", fileDescriptor_1733e2bbcb01496d)
}

var fileDescriptor_1733e2bbcb01496d = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xf3, 0x50,
	0x10, 0x25, 0xed, 0xd7, 0xf6, 0xeb, 0xa4, 0x2a, 0x5c, 0xda, 0xf4, 0x12, 0x04, 0x25, 0x0b, 0xa9,
	0x0b, 0x6f, 0x35, 0xdd, 0x8a, 0x0b, 0x2d, 0xa8, 0x20, 0x14, 0x82, 0x2b, 0x77, 0x49, 0x33, 0x96,
	0x40, 0xdb, 0xc4, 0xdc, 0xc4, 0x37, 0xf1, 0x91, 0x7c, 0x2f, 0x87, 0xfc, 0xb4, 0x69, 0x7a, 0x8b,
	0xb8, 0x49, 0xe6, 0xe7, 0x64, 0xe6, 0x9c, 0x33, 0x01, 0x1e, 0xc5, 0x61, 0x12, 0x8e, 0x7d, 0xef,
	0x2a, 0x8c, 0x30, 0x76, 0x93, 0x20, 0x5c, 0x8b, 0xac, 0x64, 0x9e, 0x2d, 0xc2, 0x70, 0xb1, 0xc4,
	0x71, 0x96, 0x79, 0xe9, 0xfb, 0x38, 0x09, 0x56, 0x28, 0x13, 0x77, 0x15, 0xe5, 0x00, 0xeb, 0x5b,
	0x03, 0x7d, 0xea, 0xcd, 0xca, 0xcf, 0xd8, 0x31, 0x34, 0x02, 0x9f, 0x6b, 0xe7, 0xda, 0xa8, 0xeb,
	0x50, 0xc4, 0x38, 0x74, 0xd6, 0xe9, 0xca, 0xc3, 0xf8, 0x86, 0x37, 0xa8, 0xa8, 0x39, 0x65, 0xba,
	0xed, 0xd8, 0xbc, 0x59, 0xed, 0xd8, 0xcc, 0x80, 0x76, 0x8c, 0x32, 0x5d, 0x26, 0xfc, 0x5f, 0xd6,
	0x28, 0x32, 0x76, 0x0a, 0xdd, 0x0d, 0x3f, 0xde, 0xca, 0x56, 0x6c, 0x0b, 0xec, 0x96, 0x88, 0xb8,
	0x09, 0x3e, 0xc4, 0x48, 0x4f, 0x9f, 0xb7, 0xa9, 0xaf, 0xdb, 0xa6, 0xc8, 0x05, 0x88, 0x52, 0x80,
	0x78, 0x2d, 0x05, 0x38, 0x55, 0xb8, 0xf5, 0x04, 0x46, 0x1e, 0x6e, 0xa4, 0x38, 0xf8, 0x91, 0x12,
	0x92, 0x09, 0xd0, 0xfd, 0xad, 0xc0, 0x4c, 0x9a, 0x6e, 0xf7, 0x44, 0x45, 0xb4, 0x53, 0x05, 0x58,
	0x97, 0x30, 0xdc, 0x9b, 0x24, 0xa3, 0x70, 0x2d, 0xb1, 0x6e, 0x8e, 0x65, 0x40, 0xff, 0x25, 0x90,
	0x49, 0x7d, 0xa5, 0xf5, 0x0c, 0x83, 0x5a, 0xbd, 0x18, 0x70, 0x0d, 0xbd, 0xca, 0x2a, 0x49, 0xa3,
	0x9a, 0x7b, 0x64, 0x76, 0x10, 0xd6, 0x08, 0x8c, 0x29, 0x2e, 0x51, 0xa1, 0xab, 0x4e, 0x66, 0x02,
	0xc3, 0x3d, 0x64, 0xb1, 0x96, 0x4e, 0x25, 0xd3, 0xf9, 0x1c, 0xa5, 0xcc, 0xf0, 0xff, 0x9d, 0x32,
	0xb5, 0x2e, 0xa0, 0xef, 0xa0, 0xeb, 0xff, 0x3a, 0xfc, 0x11, 0x06, 0x35, 0x5c, 0x31, 0xfa, 0x8f,
	0xee, 0xda, 0x5f, 0x0d, 0xe8, 0xcc, 0x30, 0xfe, 0x0c, 0xe6, 0xc8, 0xa6, 0x70, 0x52, 0x73, 0x9a,
	0x0d, 0x85, 0xfa, 0x8a, 0x26, 0x17, 0x87, 0x8e, 0x72, 0x07, 0x47, 0x3b, 0x66, 0xb3, 0x81, 0x50,
	0x1d, 0xc5, 0x34, 0x84, 0xfa, 0x26, 0xc4, 0xa2, 0xe6, 0x1b, 0xb1, 0x50, 0x7b, 0x4e, 0x2c, 0x0e,
	0x59, 0x4c, 0x2c, 0x76, 0x0c, 0x22, 0x16, 0x2a, 0x63, 0x89, 0x85, 0xd2, 0xc7, 0xfb, 0xce, 0x5b,
	0x2b, 0xff, 0xc5, 0xdb, 0xd9, 0x6b, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x0a, 0x70, 0x3e,
	0xd4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OerviceClient is the client API for Oervice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OerviceClient interface {
	CreateOperation(ctx context.Context, in *CreateOperationRequest, opts ...grpc.CallOption) (*CreateOperationResponse, error)
	ListOperation(ctx context.Context, in *ListOperationRequest, opts ...grpc.CallOption) (*ListOperationResponse, error)
	DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*DeleteOperationResponse, error)
	ReadOperation(ctx context.Context, in *ReadOperationRequest, opts ...grpc.CallOption) (*ReadOperationResponse, error)
}

type oerviceClient struct {
	cc grpc.ClientConnInterface
}

func NewOerviceClient(cc grpc.ClientConnInterface) OerviceClient {
	return &oerviceClient{cc}
}

func (c *oerviceClient) CreateOperation(ctx context.Context, in *CreateOperationRequest, opts ...grpc.CallOption) (*CreateOperationResponse, error) {
	out := new(CreateOperationResponse)
	err := c.cc.Invoke(ctx, "/Oervice/CreateOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oerviceClient) ListOperation(ctx context.Context, in *ListOperationRequest, opts ...grpc.CallOption) (*ListOperationResponse, error) {
	out := new(ListOperationResponse)
	err := c.cc.Invoke(ctx, "/Oervice/ListOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oerviceClient) DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*DeleteOperationResponse, error) {
	out := new(DeleteOperationResponse)
	err := c.cc.Invoke(ctx, "/Oervice/DeleteOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oerviceClient) ReadOperation(ctx context.Context, in *ReadOperationRequest, opts ...grpc.CallOption) (*ReadOperationResponse, error) {
	out := new(ReadOperationResponse)
	err := c.cc.Invoke(ctx, "/Oervice/ReadOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OerviceServer is the server API for Oervice service.
type OerviceServer interface {
	CreateOperation(context.Context, *CreateOperationRequest) (*CreateOperationResponse, error)
	ListOperation(context.Context, *ListOperationRequest) (*ListOperationResponse, error)
	DeleteOperation(context.Context, *DeleteOperationRequest) (*DeleteOperationResponse, error)
	ReadOperation(context.Context, *ReadOperationRequest) (*ReadOperationResponse, error)
}

// UnimplementedOerviceServer can be embedded to have forward compatible implementations.
type UnimplementedOerviceServer struct {
}

func (*UnimplementedOerviceServer) CreateOperation(ctx context.Context, req *CreateOperationRequest) (*CreateOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOperation not implemented")
}
func (*UnimplementedOerviceServer) ListOperation(ctx context.Context, req *ListOperationRequest) (*ListOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperation not implemented")
}
func (*UnimplementedOerviceServer) DeleteOperation(ctx context.Context, req *DeleteOperationRequest) (*DeleteOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperation not implemented")
}
func (*UnimplementedOerviceServer) ReadOperation(ctx context.Context, req *ReadOperationRequest) (*ReadOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOperation not implemented")
}

func RegisterOerviceServer(s *grpc.Server, srv OerviceServer) {
	s.RegisterService(&_Oervice_serviceDesc, srv)
}

func _Oervice_CreateOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OerviceServer).CreateOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Oervice/CreateOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OerviceServer).CreateOperation(ctx, req.(*CreateOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oervice_ListOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OerviceServer).ListOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Oervice/ListOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OerviceServer).ListOperation(ctx, req.(*ListOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oervice_DeleteOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OerviceServer).DeleteOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Oervice/DeleteOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OerviceServer).DeleteOperation(ctx, req.(*DeleteOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oervice_ReadOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OerviceServer).ReadOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Oervice/ReadOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OerviceServer).ReadOperation(ctx, req.(*ReadOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Oervice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Oervice",
	HandlerType: (*OerviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOperation",
			Handler:    _Oervice_CreateOperation_Handler,
		},
		{
			MethodName: "ListOperation",
			Handler:    _Oervice_ListOperation_Handler,
		},
		{
			MethodName: "DeleteOperation",
			Handler:    _Oervice_DeleteOperation_Handler,
		},
		{
			MethodName: "ReadOperation",
			Handler:    _Oervice_ReadOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/db-operation.proto",
}
