// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user_service.proto

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

type MeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeRequest) Reset()         { *m = MeRequest{} }
func (m *MeRequest) String() string { return proto.CompactTextString(m) }
func (*MeRequest) ProtoMessage()    {}
func (*MeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_292f630cd9eb4c90, []int{0}
}
func (m *MeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeRequest.Unmarshal(m, b)
}
func (m *MeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeRequest.Marshal(b, m, deterministic)
}
func (m *MeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeRequest.Merge(m, src)
}
func (m *MeRequest) XXX_Size() int {
	return xxx_messageInfo_MeRequest.Size(m)
}
func (m *MeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeRequest proto.InternalMessageInfo

type MeResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeResponse) Reset()         { *m = MeResponse{} }
func (m *MeResponse) String() string { return proto.CompactTextString(m) }
func (*MeResponse) ProtoMessage()    {}
func (*MeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_292f630cd9eb4c90, []int{1}
}
func (m *MeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeResponse.Unmarshal(m, b)
}
func (m *MeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeResponse.Marshal(b, m, deterministic)
}
func (m *MeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeResponse.Merge(m, src)
}
func (m *MeResponse) XXX_Size() int {
	return xxx_messageInfo_MeResponse.Size(m)
}
func (m *MeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeResponse proto.InternalMessageInfo

func (m *MeResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MeResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MeResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_292f630cd9eb4c90, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*MeRequest)(nil), "isjyi.demo.MeRequest")
	proto.RegisterType((*MeResponse)(nil), "isjyi.demo.MeResponse")
	proto.RegisterType((*User)(nil), "isjyi.demo.User")
}

func init() { proto.RegisterFile("user_service.proto", fileDescriptor_292f630cd9eb4c90) }

var fileDescriptor_292f630cd9eb4c90 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x49, 0x2e, 0x1e, 0x66, 0x82, 0x70, 0xac, 0x28, 0x21, 0x58, 0x84, 0x60, 0x91, 0xc6,
	0x5d, 0x3c, 0x4b, 0x3b, 0x2b, 0x9b, 0x6b, 0x22, 0xa2, 0xd8, 0x48, 0x92, 0x1b, 0xf7, 0x56, 0x4c,
	0x26, 0x66, 0x37, 0x07, 0xb6, 0xbe, 0x82, 0x8f, 0xe6, 0x2b, 0xf8, 0x20, 0xb2, 0xbb, 0x70, 0xca,
	0x75, 0xff, 0x37, 0xf3, 0x33, 0x7c, 0x03, 0x6c, 0xd2, 0x38, 0x3e, 0x6b, 0x1c, 0xb7, 0xaa, 0x45,
	0x3e, 0x8c, 0x64, 0x88, 0x81, 0xd2, 0xaf, 0x1f, 0x8a, 0xaf, 0xb1, 0xa3, 0xec, 0x4c, 0x12, 0xc9,
	0x37, 0x14, 0xf5, 0xa0, 0x44, 0xdd, 0xf7, 0x64, 0x6a, 0xa3, 0xa8, 0xd7, 0xbe, 0x99, 0x5d, 0x48,
	0x65, 0x36, 0x53, 0xc3, 0x5b, 0xea, 0x84, 0x24, 0x49, 0xc2, 0x8d, 0x9b, 0xe9, 0xc5, 0x91, 0x03,
	0x97, 0x7c, 0xbd, 0x48, 0x20, 0x5e, 0x61, 0x85, 0xef, 0x13, 0x6a, 0x53, 0x3c, 0x02, 0x58, 0xd0,
	0x03, 0xf5, 0x1a, 0x19, 0x83, 0xa8, 0xa5, 0x35, 0xa6, 0x41, 0x1e, 0x94, 0x07, 0x95, 0xcb, 0x6c,
	0x01, 0xb3, 0x4e, 0xcb, 0x34, 0xcc, 0x83, 0x32, 0xae, 0x6c, 0x64, 0xe7, 0x10, 0x59, 0xdf, 0x74,
	0x96, 0x07, 0x65, 0xb2, 0x5c, 0xf0, 0x3f, 0x51, 0x7e, 0xaf, 0x71, 0xac, 0xdc, 0xb6, 0x28, 0x20,
	0xb2, 0xc4, 0x32, 0x38, 0xb4, 0xdc, 0xd7, 0x9d, 0xbf, 0x1b, 0x57, 0x3b, 0x5e, 0x3e, 0x40, 0x62,
	0x3b, 0x77, 0xfe, 0x71, 0x76, 0x0b, 0xe1, 0x0a, 0xd9, 0xc9, 0xff, 0x83, 0x3b, 0xd3, 0xec, 0x74,
	0x7f, 0xec, 0x9d, 0x8b, 0xe3, 0xcf, 0xef, 0x9f, 0xaf, 0xf0, 0x88, 0x25, 0x62, 0x7b, 0x29, 0xea,
	0xc9, 0x6c, 0x44, 0x87, 0x37, 0xf3, 0xa7, 0x88, 0x5f, 0x0f, 0x4d, 0x33, 0x77, 0x2f, 0x5f, 0xfd,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x22, 0x1d, 0x8a, 0x7e, 0x61, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Me(ctx context.Context, in *MeRequest, opts ...grpc.CallOption) (*MeResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Me(ctx context.Context, in *MeRequest, opts ...grpc.CallOption) (*MeResponse, error) {
	out := new(MeResponse)
	err := c.cc.Invoke(ctx, "/isjyi.demo.UserService/Me", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Me(context.Context, *MeRequest) (*MeResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Me(ctx context.Context, req *MeRequest) (*MeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/isjyi.demo.UserService/Me",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Me(ctx, req.(*MeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "isjyi.demo.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Me",
			Handler:    _UserService_Me_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}