// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{3}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManytimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManytimesResponse) Reset()         { *m = GreetManytimesResponse{} }
func (m *GreetManytimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManytimesResponse) ProtoMessage()    {}
func (*GreetManytimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{4}
}

func (m *GreetManytimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManytimesResponse.Unmarshal(m, b)
}
func (m *GreetManytimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManytimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManytimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManytimesResponse.Merge(m, src)
}
func (m *GreetManytimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManytimesResponse.Size(m)
}
func (m *GreetManytimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManytimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManytimesResponse proto.InternalMessageInfo

func (m *GreetManytimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreetRequest) Reset()         { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()    {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{5}
}

func (m *LongGreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetRequest.Unmarshal(m, b)
}
func (m *LongGreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetRequest.Marshal(b, m, deterministic)
}
func (m *LongGreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetRequest.Merge(m, src)
}
func (m *LongGreetRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreetRequest.Size(m)
}
func (m *LongGreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetRequest proto.InternalMessageInfo

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreetResponse) Reset()         { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()    {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{6}
}

func (m *LongGreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetResponse.Unmarshal(m, b)
}
func (m *LongGreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetResponse.Marshal(b, m, deterministic)
}
func (m *LongGreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetResponse.Merge(m, src)
}
func (m *LongGreetResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreetResponse.Size(m)
}
func (m *LongGreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetResponse proto.InternalMessageInfo

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetManytimesResponse)(nil), "greet.GreetManytimesResponse")
	proto.RegisterType((*LongGreetRequest)(nil), "greet.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "greet.LongGreetResponse")
}

func init() {
	proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871)
}

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x15,
	0xcc, 0x51, 0x72, 0xe3, 0xe2, 0x70, 0x07, 0x31, 0x32, 0xf3, 0xd2, 0x85, 0x64, 0xb9, 0xb8, 0xd2,
	0x32, 0x8b, 0x8a, 0x4b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x38, 0xc1, 0x22, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0xd2, 0x5c, 0x9c, 0x39, 0x89, 0x30, 0x59, 0x26,
	0xb0, 0x2c, 0x07, 0x48, 0x00, 0x24, 0xa9, 0x64, 0xcd, 0xc5, 0x03, 0x36, 0x27, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x9b, 0x8b, 0x23, 0x1d, 0x6a, 0x2e, 0xd8, 0x24, 0x6e, 0x23, 0x7e,
	0x3d, 0x88, 0xf5, 0x30, 0xeb, 0x82, 0xe0, 0x0a, 0x94, 0xd4, 0xb9, 0x78, 0xa1, 0x9a, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0, 0xae,
	0x80, 0xf2, 0x94, 0x5c, 0xb8, 0x44, 0xc1, 0x0a, 0x7d, 0x13, 0xf3, 0x2a, 0x43, 0x32, 0x73, 0x53,
	0x8b, 0xc9, 0xb2, 0xce, 0x80, 0x4b, 0x0c, 0x6e, 0x4a, 0x09, 0xc4, 0x14, 0x02, 0xf6, 0xda, 0x73,
	0x09, 0xf8, 0xe4, 0xe7, 0xa5, 0x93, 0xef, 0x43, 0x6d, 0x2e, 0x41, 0x24, 0x03, 0xf0, 0xdb, 0x66,
	0x74, 0x8b, 0x11, 0x1a, 0x98, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x26, 0x5c, 0xac,
	0x60, 0xbe, 0x90, 0x30, 0xb2, 0x0d, 0x50, 0x87, 0x48, 0x89, 0xa0, 0x0a, 0x42, 0x0c, 0x57, 0x62,
	0x10, 0x0a, 0x80, 0x86, 0x2a, 0x2c, 0xb0, 0x84, 0x64, 0x90, 0x15, 0xa2, 0x07, 0xa1, 0x94, 0x2c,
	0xba, 0x2c, 0x4a, 0xd0, 0x28, 0x31, 0x18, 0x30, 0x0a, 0x39, 0x71, 0x71, 0xc2, 0x7d, 0x21, 0x24,
	0x0e, 0x55, 0x8f, 0x1e, 0x30, 0x52, 0x12, 0x98, 0x12, 0x30, 0x33, 0x34, 0x18, 0x9d, 0x38, 0xa3,
	0xd8, 0xa1, 0xc9, 0x31, 0x89, 0x0d, 0x9c, 0x12, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45,
	0xe6, 0x65, 0xab, 0xa6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetManyTime(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimeClient, error)
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTime(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTime", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimeClient interface {
	Recv() (*GreetManytimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimeClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimeClient) Recv() (*GreetManytimesResponse, error) {
	m := new(GreetManytimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetManyTime(*GreetManyTimesRequest, GreetService_GreetManyTimeServer) error
	LongGreet(GreetService_LongGreetServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManyTime(req *GreetManyTimesRequest, srv GreetService_GreetManyTimeServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTime not implemented")
}
func (*UnimplementedGreetServiceServer) LongGreet(srv GreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTime(m, &greetServiceGreetManyTimeServer{stream})
}

type GreetService_GreetManyTimeServer interface {
	Send(*GreetManytimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimeServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimeServer) Send(m *GreetManytimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTime",
			Handler:       _GreetService_GreetManyTime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
