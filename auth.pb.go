// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

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

type CredentialsMessage struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialsMessage) Reset()         { *m = CredentialsMessage{} }
func (m *CredentialsMessage) String() string { return proto.CompactTextString(m) }
func (*CredentialsMessage) ProtoMessage()    {}
func (*CredentialsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *CredentialsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialsMessage.Unmarshal(m, b)
}
func (m *CredentialsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialsMessage.Marshal(b, m, deterministic)
}
func (m *CredentialsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialsMessage.Merge(m, src)
}
func (m *CredentialsMessage) XXX_Size() int {
	return xxx_messageInfo_CredentialsMessage.Size(m)
}
func (m *CredentialsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialsMessage proto.InternalMessageInfo

func (m *CredentialsMessage) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CredentialsMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type TokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CredentialsMessage)(nil), "proto.CredentialsMessage")
	proto.RegisterType((*TokenResponse)(nil), "proto.TokenResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6e, 0x5c, 0x42, 0xce, 0x45,
	0xa9, 0x29, 0xa9, 0x79, 0x25, 0x99, 0x89, 0x39, 0xc5, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9,
	0x42, 0x22, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x10, 0x8e, 0x90, 0x14, 0x17, 0x47, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13,
	0x58, 0x02, 0xce, 0x57, 0x52, 0xe5, 0xe2, 0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x0b, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x06, 0x1b, 0x01, 0x16, 0x80, 0x19, 0x01, 0xe6, 0x18, 0x35, 0x31, 0x72, 0x71,
	0x3b, 0x96, 0x96, 0x64, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59, 0x70, 0xb1, 0xfa,
	0xe4, 0xa7, 0x67, 0xe6, 0x09, 0x49, 0x42, 0x9c, 0xa5, 0x87, 0xe9, 0x18, 0x29, 0x11, 0xa8, 0x14,
	0xaa, 0xf9, 0x96, 0x5c, 0x6c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x24, 0x6b, 0x4d, 0x62, 0x03,
	0x0b, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0xdc, 0xe8, 0xe3, 0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *CredentialsMessage, opts ...grpc.CallOption) (*TokenResponse, error)
	Create(ctx context.Context, in *CredentialsMessage, opts ...grpc.CallOption) (*TokenResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *CredentialsMessage, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Create(ctx context.Context, in *CredentialsMessage, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *CredentialsMessage) (*TokenResponse, error)
	Create(context.Context, *CredentialsMessage) (*TokenResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *CredentialsMessage) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Create(ctx context.Context, req *CredentialsMessage) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*CredentialsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Create(ctx, req.(*CredentialsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AuthService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
