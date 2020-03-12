// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/sum.proto

package sumpb

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

type SumRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{0}
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

func (m *SumRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{1}
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

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type NumRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumRequest) Reset()         { *m = NumRequest{} }
func (m *NumRequest) String() string { return proto.CompactTextString(m) }
func (*NumRequest) ProtoMessage()    {}
func (*NumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{2}
}

func (m *NumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumRequest.Unmarshal(m, b)
}
func (m *NumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumRequest.Marshal(b, m, deterministic)
}
func (m *NumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumRequest.Merge(m, src)
}
func (m *NumRequest) XXX_Size() int {
	return xxx_messageInfo_NumRequest.Size(m)
}
func (m *NumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumRequest proto.InternalMessageInfo

func (m *NumRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type NumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumResponse) Reset()         { *m = NumResponse{} }
func (m *NumResponse) String() string { return proto.CompactTextString(m) }
func (*NumResponse) ProtoMessage()    {}
func (*NumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{3}
}

func (m *NumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumResponse.Unmarshal(m, b)
}
func (m *NumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumResponse.Marshal(b, m, deterministic)
}
func (m *NumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumResponse.Merge(m, src)
}
func (m *NumResponse) XXX_Size() int {
	return xxx_messageInfo_NumResponse.Size(m)
}
func (m *NumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NumResponse proto.InternalMessageInfo

func (m *NumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AverageRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRequest) Reset()         { *m = AverageRequest{} }
func (m *AverageRequest) String() string { return proto.CompactTextString(m) }
func (*AverageRequest) ProtoMessage()    {}
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{4}
}

func (m *AverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRequest.Unmarshal(m, b)
}
func (m *AverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRequest.Marshal(b, m, deterministic)
}
func (m *AverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRequest.Merge(m, src)
}
func (m *AverageRequest) XXX_Size() int {
	return xxx_messageInfo_AverageRequest.Size(m)
}
func (m *AverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRequest proto.InternalMessageInfo

func (m *AverageRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type AverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{5}
}

func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (m *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(m, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaximumRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FindMaximumResponse struct {
	Max                  int32    `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_723b391600f96c49, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sum.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sum.SumResponse")
	proto.RegisterType((*NumRequest)(nil), "sum.NumRequest")
	proto.RegisterType((*NumResponse)(nil), "sum.NumResponse")
	proto.RegisterType((*AverageRequest)(nil), "sum.AverageRequest")
	proto.RegisterType((*AverageResponse)(nil), "sum.AverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "sum.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "sum.FindMaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "sum.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "sum.SquareRootResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/sum.proto", fileDescriptor_723b391600f96c49)
}

var fileDescriptor_723b391600f96c49 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0x6f, 0x1a, 0xda, 0x0b, 0xa7, 0x97, 0x36, 0x77, 0x2a, 0x6d, 0xc8, 0xc2, 0x8f, 0x80,
	0x5a, 0x15, 0x5a, 0x51, 0x5c, 0x4b, 0xad, 0x74, 0x67, 0x90, 0x64, 0xe7, 0x46, 0x26, 0xf1, 0x20,
	0x81, 0x4e, 0x26, 0x9d, 0xc9, 0x94, 0xfe, 0x47, 0xff, 0x94, 0x24, 0x99, 0x34, 0xfd, 0x50, 0x71,
	0x37, 0xe7, 0xcc, 0xc3, 0xc3, 0x99, 0xf3, 0x0e, 0x9c, 0x44, 0x74, 0x1e, 0xa9, 0x39, 0xcd, 0xb8,
	0x18, 0xd7, 0xc7, 0x34, 0x1c, 0x4b, 0xc5, 0x46, 0xa9, 0xe0, 0x19, 0x27, 0xa6, 0x54, 0xcc, 0x1d,
	0x02, 0x04, 0x8a, 0xf9, 0xb8, 0x50, 0x28, 0x33, 0xf2, 0x0f, 0x0c, 0x6a, 0x1b, 0xc7, 0xc6, 0xb0,
	0xe9, 0x1b, 0x34, 0xaf, 0x42, 0xbb, 0x51, 0x56, 0xa1, 0x7b, 0x0a, 0xed, 0x82, 0x94, 0x29, 0x4f,
	0x24, 0x92, 0x3e, 0xb4, 0x04, 0x4a, 0x35, 0xcf, 0x34, 0xaf, 0x2b, 0xf7, 0x10, 0xc0, 0xab, 0x85,
	0x16, 0x98, 0x89, 0x62, 0x1a, 0xc9, 0x8f, 0xb9, 0xc6, 0xfb, 0x85, 0xc6, 0x85, 0xce, 0x64, 0x89,
	0x82, 0xbe, 0xe3, 0xf7, 0xaa, 0x0b, 0xe8, 0xae, 0x99, 0x2f, 0x75, 0x8d, 0xb5, 0xee, 0x0c, 0xc8,
	0x2c, 0x4e, 0xde, 0x9e, 0xe8, 0x2a, 0x66, 0x3f, 0x4d, 0x77, 0x0e, 0xbd, 0x2d, 0x4e, 0x6b, 0x2d,
	0x30, 0x19, 0x5d, 0x55, 0x20, 0xa3, 0x2b, 0xf7, 0x0a, 0xfe, 0x07, 0x0b, 0x45, 0x05, 0xfa, 0x9c,
	0x67, 0x95, 0xaf, 0x0f, 0xad, 0x44, 0xb1, 0x10, 0x45, 0xf5, 0x98, 0xb2, 0x72, 0xef, 0x80, 0x6c,
	0xc2, 0x5a, 0x7a, 0x04, 0xed, 0xf2, 0xfe, 0x55, 0x70, 0x5e, 0x0e, 0x6c, 0xf8, 0x50, 0xb6, 0x72,
	0xf0, 0xe6, 0xa3, 0x01, 0xd6, 0x54, 0x67, 0x87, 0x01, 0x8a, 0x65, 0x1c, 0x21, 0xb9, 0x04, 0x33,
	0x50, 0x8c, 0x74, 0x47, 0x79, 0x90, 0x75, 0x74, 0x8e, 0x55, 0x37, 0x4a, 0xbf, 0xfb, 0x87, 0x4c,
	0xc0, 0x7e, 0x16, 0x31, 0x43, 0xaf, 0x70, 0x3e, 0x62, 0xc4, 0x59, 0xca, 0x65, 0x9c, 0xc5, 0x3c,
	0xd1, 0x02, 0x6f, 0x57, 0xe0, 0x6d, 0x0a, 0xae, 0x0d, 0x72, 0x0f, 0x9d, 0x29, 0x67, 0xa9, 0xca,
	0x50, 0xaf, 0x9a, 0xf4, 0x0a, 0x6e, 0x3b, 0x1c, 0xe7, 0x60, 0xbb, 0x59, 0x09, 0x86, 0x06, 0x99,
	0x41, 0x7b, 0x63, 0xa3, 0x64, 0x50, 0x80, 0xfb, 0x59, 0x38, 0xf6, 0xfe, 0x45, 0x6d, 0x29, 0x06,
	0x81, 0x7a, 0x87, 0xa4, 0x5f, 0xbe, 0x76, 0x37, 0x01, 0x67, 0xb0, 0xd7, 0xaf, 0x24, 0x0f, 0x7f,
	0x5f, 0x9a, 0x52, 0xb1, 0x34, 0x0c, 0x5b, 0xc5, 0xf7, 0xbf, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xc8, 0x0a, 0x89, 0x23, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *NumRequest, opts ...grpc.CallOption) (CalculateService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculateService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMaximumClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sum.CalculateService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) PrimeNumberDecomposition(ctx context.Context, in *NumRequest, opts ...grpc.CallOption) (CalculateService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[0], "/sum.CalculateService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculateService_PrimeNumberDecompositionClient interface {
	Recv() (*NumResponse, error)
	grpc.ClientStream
}

type calculateServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculateServicePrimeNumberDecompositionClient) Recv() (*NumResponse, error) {
	m := new(NumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculateService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[1], "/sum.CalculateService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceComputeAverageClient{stream}
	return x, nil
}

type CalculateService_ComputeAverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculateServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculateServiceComputeAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceComputeAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[2], "/sum.CalculateService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceFindMaximumClient{stream}
	return x, nil
}

type CalculateService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculateServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculateServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/sum.CalculateService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*NumRequest, CalculateService_PrimeNumberDecompositionServer) error
	ComputeAverage(CalculateService_ComputeAverageServer) error
	FindMaximum(CalculateService_FindMaximumServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculateServiceServer) PrimeNumberDecomposition(req *NumRequest, srv CalculateService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculateServiceServer) ComputeAverage(srv CalculateService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculateServiceServer) FindMaximum(srv CalculateService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (*UnimplementedCalculateServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.CalculateService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NumRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculateServiceServer).PrimeNumberDecomposition(m, &calculateServicePrimeNumberDecompositionServer{stream})
}

type CalculateService_PrimeNumberDecompositionServer interface {
	Send(*NumResponse) error
	grpc.ServerStream
}

type calculateServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculateServicePrimeNumberDecompositionServer) Send(m *NumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculateService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).ComputeAverage(&calculateServiceComputeAverageServer{stream})
}

type CalculateService_ComputeAverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculateServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculateServiceComputeAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceComputeAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).FindMaximum(&calculateServiceFindMaximumServer{stream})
}

type CalculateService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculateServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculateServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.CalculateService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculateService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculateService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculateService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculateService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/sum.proto",
}
