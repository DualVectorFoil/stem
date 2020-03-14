// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package pb

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

type Profile struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	PhoneNum             *string  `protobuf:"bytes,2,req,name=phoneNum" json:"phoneNum,omitempty"`
	AvatarUrl            *string  `protobuf:"bytes,3,req,name=avatarUrl" json:"avatarUrl,omitempty"`
	UserName             *string  `protobuf:"bytes,4,req,name=userName" json:"userName,omitempty"`
	Locale               *string  `protobuf:"bytes,5,req,name=locale" json:"locale,omitempty"`
	Bio                  *string  `protobuf:"bytes,6,req,name=bio" json:"bio,omitempty"`
	Followers            *int32   `protobuf:"varint,7,req,name=followers" json:"followers,omitempty"`
	Following            *int32   `protobuf:"varint,8,req,name=following" json:"following,omitempty"`
	ArtworkCount         *int32   `protobuf:"varint,9,req,name=artworkCount" json:"artworkCount,omitempty"`
	Pwd                  *string  `protobuf:"bytes,10,req,name=pwd" json:"pwd,omitempty"`
	RegisteredAt         *int64   `protobuf:"varint,11,req,name=registeredAt" json:"registeredAt,omitempty"`
	LastLoginAt          *int64   `protobuf:"varint,12,req,name=lastLoginAt" json:"lastLoginAt,omitempty"`
	Token                *string  `protobuf:"bytes,13,req,name=token" json:"token,omitempty"`
	Email                *string  `protobuf:"bytes,14,req,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Profile) GetPhoneNum() string {
	if m != nil && m.PhoneNum != nil {
		return *m.PhoneNum
	}
	return ""
}

func (m *Profile) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *Profile) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *Profile) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *Profile) GetBio() string {
	if m != nil && m.Bio != nil {
		return *m.Bio
	}
	return ""
}

func (m *Profile) GetFollowers() int32 {
	if m != nil && m.Followers != nil {
		return *m.Followers
	}
	return 0
}

func (m *Profile) GetFollowing() int32 {
	if m != nil && m.Following != nil {
		return *m.Following
	}
	return 0
}

func (m *Profile) GetArtworkCount() int32 {
	if m != nil && m.ArtworkCount != nil {
		return *m.ArtworkCount
	}
	return 0
}

func (m *Profile) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *Profile) GetRegisteredAt() int64 {
	if m != nil && m.RegisteredAt != nil {
		return *m.RegisteredAt
	}
	return 0
}

func (m *Profile) GetLastLoginAt() int64 {
	if m != nil && m.LastLoginAt != nil {
		return *m.LastLoginAt
	}
	return 0
}

func (m *Profile) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Profile) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type RegisterRequest struct {
	PhoneNum             *string  `protobuf:"bytes,1,req,name=phoneNum" json:"phoneNum,omitempty"`
	UserName             *string  `protobuf:"bytes,2,req,name=userName" json:"userName,omitempty"`
	Email                *string  `protobuf:"bytes,3,req,name=email" json:"email,omitempty"`
	Pwd                  *string  `protobuf:"bytes,4,req,name=pwd" json:"pwd,omitempty"`
	PhoneCode            *int32   `protobuf:"varint,5,req,name=phoneCode" json:"phoneCode,omitempty"`
	Id                   *string  `protobuf:"bytes,6,req,name=id" json:"id,omitempty"`
	Token                *string  `protobuf:"bytes,7,req,name=Token" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
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

func (m *RegisterRequest) GetPhoneNum() string {
	if m != nil && m.PhoneNum != nil {
		return *m.PhoneNum
	}
	return ""
}

func (m *RegisterRequest) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *RegisterRequest) GetPhoneCode() int32 {
	if m != nil && m.PhoneCode != nil {
		return *m.PhoneCode
	}
	return 0
}

func (m *RegisterRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *RegisterRequest) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type RegisterResp struct {
	IsSuccess            *bool    `protobuf:"varint,1,req,name=isSuccess" json:"isSuccess,omitempty"`
	ErrMsg               *string  `protobuf:"bytes,2,req,name=errMsg" json:"errMsg,omitempty"`
	Profile              *Profile `protobuf:"bytes,3,req,name=profile" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetIsSuccess() bool {
	if m != nil && m.IsSuccess != nil {
		return *m.IsSuccess
	}
	return false
}

func (m *RegisterResp) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *RegisterResp) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type LoginRequest struct {
	PhoneNum             *string  `protobuf:"bytes,1,req,name=phoneNum" json:"phoneNum,omitempty"`
	UserName             *string  `protobuf:"bytes,2,req,name=userName" json:"userName,omitempty"`
	Email                *string  `protobuf:"bytes,3,req,name=email" json:"email,omitempty"`
	Pwd                  *string  `protobuf:"bytes,4,req,name=pwd" json:"pwd,omitempty"`
	PhoneCode            *int32   `protobuf:"varint,5,req,name=phoneCode" json:"phoneCode,omitempty"`
	Id                   *string  `protobuf:"bytes,6,req,name=id" json:"id,omitempty"`
	Token                *string  `protobuf:"bytes,7,req,name=Token" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
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

func (m *LoginRequest) GetPhoneNum() string {
	if m != nil && m.PhoneNum != nil {
		return *m.PhoneNum
	}
	return ""
}

func (m *LoginRequest) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *LoginRequest) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *LoginRequest) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *LoginRequest) GetPhoneCode() int32 {
	if m != nil && m.PhoneCode != nil {
		return *m.PhoneCode
	}
	return 0
}

func (m *LoginRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *LoginRequest) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type LoginResp struct {
	IsSuccess            *bool    `protobuf:"varint,1,req,name=isSuccess" json:"isSuccess,omitempty"`
	ErrMsg               *string  `protobuf:"bytes,2,req,name=errMsg" json:"errMsg,omitempty"`
	Profile              *Profile `protobuf:"bytes,3,req,name=profile" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetIsSuccess() bool {
	if m != nil && m.IsSuccess != nil {
		return *m.IsSuccess
	}
	return false
}

func (m *LoginResp) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *LoginResp) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "pb.Profile")
	proto.RegisterType((*RegisterRequest)(nil), "pb.RegisterRequest")
	proto.RegisterType((*RegisterResp)(nil), "pb.RegisterResp")
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResp)(nil), "pb.LoginResp")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc7, 0x89, 0x52, 0x27, 0xf1, 0xc9, 0xc7, 0x8a, 0x56, 0x86, 0x08, 0xbd, 0x08, 0x86, 0x41,
	0xae, 0x02, 0xeb, 0x1b, 0x84, 0xde, 0x6e, 0x65, 0xb8, 0xdb, 0x03, 0x28, 0xf6, 0xa9, 0x2b, 0xe2,
	0x58, 0x9a, 0x24, 0x37, 0xef, 0x35, 0xd8, 0xe3, 0x0d, 0x86, 0x8e, 0x9d, 0xd8, 0xc9, 0xfd, 0x2e,
	0x7a, 0x97, 0xff, 0xef, 0xf8, 0x7c, 0xfd, 0x75, 0x02, 0x73, 0x99, 0x65, 0xba, 0xae, 0xfc, 0xc6,
	0x58, 0xed, 0x35, 0x67, 0x66, 0x97, 0xfc, 0x65, 0x30, 0xfe, 0x6e, 0xf5, 0x8b, 0x2a, 0x91, 0x2f,
	0x80, 0xa9, 0x5c, 0x0c, 0x56, 0x6c, 0x1d, 0xa7, 0x4c, 0xe5, 0x7c, 0x09, 0x13, 0xf3, 0xaa, 0x2b,
	0x7c, 0xaa, 0x0f, 0x82, 0x11, 0x3d, 0x6b, 0x7e, 0x0f, 0xb1, 0x7c, 0x93, 0x5e, 0xda, 0x9f, 0xb6,
	0x14, 0x43, 0x0a, 0x76, 0x20, 0x64, 0xd6, 0x0e, 0xed, 0x93, 0x3c, 0xa0, 0xb8, 0x69, 0x32, 0x4f,
	0x9a, 0x7f, 0x82, 0x51, 0xa9, 0x33, 0x59, 0xa2, 0x88, 0x28, 0xd2, 0x2a, 0x7e, 0x0b, 0xc3, 0x9d,
	0xd2, 0x62, 0x44, 0x30, 0xfc, 0x0c, 0x3d, 0x5e, 0x74, 0x59, 0xea, 0x23, 0x5a, 0x27, 0xc6, 0x2b,
	0xb6, 0x8e, 0xd2, 0x0e, 0x74, 0x51, 0x55, 0x15, 0x62, 0xd2, 0x8f, 0xaa, 0xaa, 0xe0, 0x09, 0xcc,
	0xa4, 0xf5, 0x47, 0x6d, 0xf7, 0x8f, 0x61, 0x63, 0x11, 0xd3, 0x07, 0x17, 0x2c, 0x74, 0x34, 0xc7,
	0x5c, 0x40, 0xd3, 0xd1, 0x1c, 0xf3, 0x90, 0x65, 0xb1, 0x50, 0xce, 0xa3, 0xc5, 0x7c, 0xeb, 0xc5,
	0x74, 0xc5, 0xd6, 0xc3, 0xf4, 0x82, 0xf1, 0x15, 0x4c, 0x4b, 0xe9, 0xfc, 0x57, 0x5d, 0xa8, 0x6a,
	0xeb, 0xc5, 0x8c, 0x3e, 0xe9, 0x23, 0x7e, 0x07, 0x91, 0xd7, 0x7b, 0xac, 0xc4, 0x9c, 0x2a, 0x37,
	0x22, 0x50, 0x3c, 0x48, 0x55, 0x8a, 0x45, 0x43, 0x49, 0x24, 0x7f, 0x06, 0xf0, 0x21, 0x6d, 0xcb,
	0xa7, 0xf8, 0xab, 0x46, 0xe7, 0x2f, 0x7c, 0x1f, 0x5c, 0xf9, 0xde, 0x77, 0x96, 0x5d, 0x39, 0x7b,
	0xee, 0x30, 0xec, 0x75, 0x38, 0x6d, 0x79, 0xd3, 0x6d, 0x79, 0x0f, 0x31, 0xd5, 0x7b, 0xd4, 0x79,
	0xf3, 0x08, 0x51, 0xda, 0x81, 0xf6, 0x0a, 0x46, 0xe7, 0x2b, 0xb8, 0x83, 0xe8, 0x07, 0x6d, 0x33,
	0x6e, 0xaa, 0x92, 0x48, 0xf6, 0x30, 0xeb, 0xc6, 0x76, 0x26, 0xd4, 0x54, 0xee, 0xb9, 0xce, 0x32,
	0x74, 0x8e, 0x86, 0x9e, 0xa4, 0x1d, 0x08, 0x6f, 0x8e, 0xd6, 0x7e, 0x73, 0x45, 0x3b, 0x73, 0xab,
	0xf8, 0x67, 0x18, 0x9b, 0xe6, 0xf8, 0x68, 0xe6, 0xe9, 0xc3, 0x74, 0x63, 0x76, 0x9b, 0xf6, 0x1e,
	0xd3, 0x53, 0x2c, 0xf9, 0x3d, 0x80, 0x19, 0x99, 0xfb, 0x9e, 0x1c, 0x7a, 0x85, 0xb8, 0x9d, 0xf9,
	0x3f, 0xdb, 0xf3, 0x70, 0x80, 0xc5, 0xb6, 0xf9, 0x63, 0x3f, 0xa3, 0x7d, 0x53, 0x19, 0xf2, 0x2f,
	0x30, 0x39, 0xbd, 0x0e, 0xff, 0x18, 0x72, 0xae, 0x4e, 0x6c, 0x79, 0x7b, 0x09, 0x9d, 0xe1, 0x6b,
	0x88, 0x68, 0x5c, 0x4e, 0xa1, 0xbe, 0xdb, 0xcb, 0x79, 0x8f, 0x38, 0xf3, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x0b, 0x85, 0x93, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResp, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/pb.AccountService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/pb.AccountService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResp, error)
	Login(context.Context, *LoginRequest) (*LoginResp, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAccountServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccountService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AccountService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AccountService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}