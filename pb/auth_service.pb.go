// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" validate:"required" label:"用户名"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required" label:"密码"`
	RepeatPassword       string   `protobuf:"bytes,3,opt,name=repeat_password,json=repeatPassword,proto3" json:"repeat_password,omitempty" validate:"required" label:"密码"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{2}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetRepeatPassword() string {
	if m != nil {
		return m.RepeatPassword
	}
	return ""
}

type RegisterResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{3}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "isjyi.demo.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "isjyi.demo.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "isjyi.demo.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "isjyi.demo.RegisterResponse")
}

func init() { proto.RegisterFile("auth_service.proto", fileDescriptor_0f39bb026ca10b68) }

var fileDescriptor_0f39bb026ca10b68 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbf, 0x8b, 0xd4, 0x40,
	0x14, 0x26, 0xf7, 0x8b, 0x73, 0xee, 0xbc, 0x3b, 0x47, 0x90, 0xbd, 0xb8, 0x70, 0xe7, 0x80, 0xa2,
	0x07, 0x26, 0x78, 0x36, 0x12, 0x2b, 0x03, 0x5a, 0x09, 0x4a, 0x14, 0x04, 0x9b, 0x65, 0x92, 0x3c,
	0x67, 0x47, 0x93, 0x4c, 0x76, 0x66, 0xb2, 0x62, 0xa9, 0x85, 0xd8, 0x08, 0x82, 0xfe, 0x01, 0x96,
	0x76, 0x96, 0xfe, 0x11, 0x96, 0xf6, 0x5b, 0xb9, 0x6e, 0x65, 0xb3, 0x7f, 0x81, 0x64, 0x26, 0xd9,
	0x5d, 0x96, 0x45, 0xb8, 0xee, 0xbd, 0xf7, 0xbd, 0xf9, 0xbe, 0xf7, 0xde, 0x7c, 0x08, 0xd3, 0x4a,
	0xf7, 0x7b, 0x0a, 0xe4, 0x90, 0x27, 0xe0, 0x95, 0x52, 0x68, 0x81, 0x11, 0x57, 0x2f, 0xdf, 0x70,
	0x2f, 0x85, 0x5c, 0xb8, 0x5d, 0x26, 0x04, 0xcb, 0xc0, 0xa7, 0x25, 0xf7, 0x69, 0x51, 0x08, 0x4d,
	0x35, 0x17, 0x85, 0xb2, 0x9d, 0xee, 0x4d, 0xc6, 0x75, 0xbf, 0x8a, 0xbd, 0x44, 0xe4, 0x3e, 0x13,
	0x4c, 0xf8, 0xa6, 0x1c, 0x57, 0x2f, 0x4c, 0x66, 0x12, 0x13, 0xd9, 0x76, 0xf2, 0x00, 0xed, 0x3e,
	0x14, 0x8c, 0x17, 0x11, 0x0c, 0x2a, 0x50, 0x1a, 0xbb, 0x68, 0xbb, 0x52, 0x20, 0x0b, 0x9a, 0x43,
	0xc7, 0x39, 0x76, 0xae, 0x9f, 0x8b, 0x66, 0x79, 0x8d, 0x95, 0x54, 0xa9, 0xd7, 0x42, 0xa6, 0x9d,
	0x35, 0x8b, 0xb5, 0x39, 0x39, 0x45, 0xe7, 0x1b, 0x1e, 0x55, 0x8a, 0x42, 0x01, 0xbe, 0x82, 0x76,
	0x69, 0x92, 0x80, 0x52, 0x3d, 0x2d, 0x5e, 0x41, 0xd1, 0x90, 0xed, 0xd8, 0xda, 0xd3, 0xba, 0x44,
	0xfe, 0x3a, 0x68, 0x3f, 0x02, 0xc6, 0x95, 0x06, 0xd9, 0xea, 0xdf, 0x5f, 0xd6, 0x0f, 0x6f, 0x4c,
	0x47, 0x47, 0x57, 0x87, 0x34, 0xe3, 0x29, 0xd5, 0x10, 0x10, 0x09, 0x83, 0x8a, 0x4b, 0x48, 0xc9,
	0x71, 0x46, 0x63, 0xc8, 0x02, 0x32, 0xf9, 0xf2, 0xed, 0xcf, 0x87, 0x1f, 0xe3, 0x4f, 0x1f, 0xc9,
	0xc2, 0xa8, 0xe1, 0xf2, 0xa8, 0xe1, 0xb5, 0xe9, 0xe8, 0x88, 0xfc, 0x87, 0x66, 0xfc, 0xfd, 0xfd,
	0xe4, 0xeb, 0x5b, 0x32, 0x5f, 0x09, 0x3f, 0x42, 0xfb, 0x12, 0x4a, 0xa0, 0xba, 0x37, 0xa3, 0x5a,
	0x3f, 0x13, 0xd5, 0x9e, 0x7d, 0xfe, 0xb8, 0xbd, 0xd1, 0x1d, 0x74, 0x30, 0x5f, 0xb7, 0x39, 0x13,
	0x46, 0x1b, 0x89, 0x48, 0xed, 0xae, 0x9b, 0x91, 0x89, 0xf1, 0x01, 0x5a, 0xcf, 0x15, 0x6b, 0x4e,
	0x5c, 0x87, 0xa7, 0x3f, 0x1d, 0xb4, 0x73, 0xaf, 0xd2, 0xfd, 0x27, 0xd6, 0x14, 0xf8, 0x19, 0xda,
	0x34, 0xd7, 0xc6, 0x1d, 0x6f, 0x6e, 0x0c, 0x6f, 0xf1, 0x23, 0xdd, 0xc3, 0x15, 0x88, 0xd5, 0x24,
	0x87, 0xef, 0x7e, 0xfd, 0xfe, 0xbc, 0x76, 0x91, 0xec, 0xf9, 0xc3, 0x5b, 0x7e, 0x6d, 0x36, 0x3f,
	0xab, 0xf1, 0xc0, 0x39, 0xc1, 0x09, 0xda, 0x6e, 0x47, 0xc4, 0x97, 0x17, 0x19, 0x96, 0xfe, 0xc9,
	0xed, 0xae, 0x06, 0x1b, 0x85, 0xae, 0x51, 0xb8, 0x44, 0x2e, 0xcc, 0x14, 0x64, 0xd3, 0x12, 0x38,
	0x27, 0xe1, 0xd6, 0xf3, 0x0d, 0xef, 0x6e, 0x19, 0xc7, 0x5b, 0xc6, 0x82, 0xb7, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0xc2, 0x58, 0x87, 0xf1, 0x02, 0x00, 0x00,
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
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/isjyi.demo.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/isjyi.demo.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/isjyi.demo.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/isjyi.demo.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "isjyi.demo.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
