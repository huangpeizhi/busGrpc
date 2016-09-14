// Code generated by protoc-gen-go.
// source: hservice.proto
// DO NOT EDIT!

/*
Package hservice is a generated protocol buffer package.

It is generated from these files:
	hservice.proto

It has these top-level messages:
	MessageRequest
	MessageReply
*/
package hservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the message value
type MessageRequest struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *MessageRequest) Reset() { *m = MessageRequest{} }
func (m *MessageRequest) String() string {
	return proto.MarshalTextString(m)
	//return proto.CompactTextString(m)
}
func (*MessageRequest) ProtoMessage()               {}
func (*MessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing an ok (true or false)
type MessageReply struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *MessageReply) Reset()                    { *m = MessageReply{} }
func (m *MessageReply) String() string            { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()               {}
func (*MessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MessageRequest)(nil), "hservice.MessageRequest")
	proto.RegisterType((*MessageReply)(nil), "hservice.MessageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for HookService service

type HookServiceClient interface {
	// Sends a greeting
	Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
}

type hookServiceClient struct {
	cc *grpc.ClientConn
}

func NewHookServiceClient(cc *grpc.ClientConn) HookServiceClient {
	return &hookServiceClient{cc}
}

func (c *hookServiceClient) Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := grpc.Invoke(ctx, "/hservice.HookService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HookService service

type HookServiceServer interface {
	// Sends a greeting
	Send(context.Context, *MessageRequest) (*MessageReply, error)
}

func RegisterHookServiceServer(s *grpc.Server, srv HookServiceServer) {
	s.RegisterService(&_HookService_serviceDesc, srv)
}

func _HookService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hservice.HookService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).Send(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hservice.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _HookService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("hservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x28, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xd4,
	0xb8, 0xf8, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x20, 0x1c, 0x25, 0x39, 0x2e, 0x1e, 0xb8, 0xba, 0x82, 0x9c, 0x4a, 0x21, 0x3e, 0x2e, 0xa6, 0xfc,
	0x6c, 0xb0, 0x12, 0x8e, 0x20, 0x20, 0xcb, 0xc8, 0x93, 0x8b, 0xdb, 0x23, 0x3f, 0x3f, 0x3b, 0x18,
	0x62, 0xac, 0x90, 0x15, 0x17, 0x4b, 0x70, 0x6a, 0x5e, 0x8a, 0x90, 0x84, 0x1e, 0xdc, 0x66, 0x54,
	0x6b, 0xa4, 0xc4, 0xb0, 0xc8, 0x00, 0x0d, 0x56, 0x62, 0x70, 0xd2, 0xe4, 0x12, 0x4e, 0xce, 0xcf,
	0xd5, 0x2b, 0xc9, 0xcc, 0x49, 0x35, 0xb6, 0x80, 0xab, 0x72, 0x12, 0x40, 0x32, 0x3f, 0x00, 0xe4,
	0x8b, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0x77, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xd0,
	0x2b, 0x13, 0xe0, 0x00, 0x00, 0x00,
}
