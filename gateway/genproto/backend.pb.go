// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type MessageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MessageResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "proto.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "proto.MessageResponse")
}

func init() { proto.RegisterFile("backend.proto", fileDescriptor_5ab9ba5b8d8b2ba5) }

var fileDescriptor_5ab9ba5b8d8b2ba5 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x2a, 0x5c,
	0x7c, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6,
	0x92, 0x36, 0x17, 0x3f, 0x5c, 0x55, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b,
	0x2e, 0x44, 0x08, 0xaa, 0x12, 0xc6, 0x35, 0xf2, 0xe6, 0xe2, 0x75, 0x82, 0x58, 0x15, 0x9c, 0x5a,
	0x54, 0x96, 0x5a, 0x24, 0x64, 0xc5, 0xc5, 0x0e, 0xd5, 0x2d, 0x24, 0x0a, 0xb1, 0x5d, 0x0f, 0xd5,
	0x4e, 0x29, 0x31, 0x74, 0x61, 0x88, 0x25, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x09, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0x3a, 0x22, 0xc4, 0xbe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendServerClient is the client API for BackendServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendServerClient interface {
	Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type backendServerClient struct {
	cc *grpc.ClientConn
}

func NewBackendServerClient(cc *grpc.ClientConn) BackendServerClient {
	return &backendServerClient{cc}
}

func (c *backendServerClient) Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/proto.BackendServer/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServerServer is the server API for BackendServer service.
type BackendServerServer interface {
	Message(context.Context, *MessageRequest) (*MessageResponse, error)
}

func RegisterBackendServerServer(s *grpc.Server, srv BackendServerServer) {
	s.RegisterService(&_BackendServer_serviceDesc, srv)
}

func _BackendServer_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServerServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BackendServer/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServerServer).Message(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackendServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BackendServer",
	HandlerType: (*BackendServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Message",
			Handler:    _BackendServer_Message_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}
