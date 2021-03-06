// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *User) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserId struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Ids struct {
	Ids                  []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ids) Reset()         { *m = Ids{} }
func (m *Ids) String() string { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()    {}
func (*Ids) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Ids) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ids.Unmarshal(m, b)
}
func (m *Ids) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ids.Marshal(b, m, deterministic)
}
func (m *Ids) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ids.Merge(m, src)
}
func (m *Ids) XXX_Size() int {
	return xxx_messageInfo_Ids.Size(m)
}
func (m *Ids) XXX_DiscardUnknown() {
	xxx_messageInfo_Ids.DiscardUnknown(m)
}

var xxx_messageInfo_Ids proto.InternalMessageInfo

func (m *Ids) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Usernames struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Username             []string `protobuf:"bytes,3,rep,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Usernames) Reset()         { *m = Usernames{} }
func (m *Usernames) String() string { return proto.CompactTextString(m) }
func (*Usernames) ProtoMessage()    {}
func (*Usernames) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Usernames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Usernames.Unmarshal(m, b)
}
func (m *Usernames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Usernames.Marshal(b, m, deterministic)
}
func (m *Usernames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Usernames.Merge(m, src)
}
func (m *Usernames) XXX_Size() int {
	return xxx_messageInfo_Usernames.Size(m)
}
func (m *Usernames) XXX_DiscardUnknown() {
	xxx_messageInfo_Usernames.DiscardUnknown(m)
}

var xxx_messageInfo_Usernames proto.InternalMessageInfo

func (m *Usernames) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Usernames) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Usernames) GetUsername() []string {
	if m != nil {
		return m.Username
	}
	return nil
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type IdResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdResponse) Reset()         { *m = IdResponse{} }
func (m *IdResponse) String() string { return proto.CompactTextString(m) }
func (*IdResponse) ProtoMessage()    {}
func (*IdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *IdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdResponse.Unmarshal(m, b)
}
func (m *IdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdResponse.Marshal(b, m, deterministic)
}
func (m *IdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdResponse.Merge(m, src)
}
func (m *IdResponse) XXX_Size() int {
	return xxx_messageInfo_IdResponse.Size(m)
}
func (m *IdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdResponse proto.InternalMessageInfo

func (m *IdResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *IdResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *IdResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserListResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Page                 *PageMap `protobuf:"bytes,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserListResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserListResponse) GetPage() *PageMap {
	if m != nil {
		return m.Page
	}
	return nil
}

type PageNumPageSize struct {
	PageNum              int32    `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageNumPageSize) Reset()         { *m = PageNumPageSize{} }
func (m *PageNumPageSize) String() string { return proto.CompactTextString(m) }
func (*PageNumPageSize) ProtoMessage()    {}
func (*PageNumPageSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *PageNumPageSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageNumPageSize.Unmarshal(m, b)
}
func (m *PageNumPageSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageNumPageSize.Marshal(b, m, deterministic)
}
func (m *PageNumPageSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageNumPageSize.Merge(m, src)
}
func (m *PageNumPageSize) XXX_Size() int {
	return xxx_messageInfo_PageNumPageSize.Size(m)
}
func (m *PageNumPageSize) XXX_DiscardUnknown() {
	xxx_messageInfo_PageNumPageSize.DiscardUnknown(m)
}

var xxx_messageInfo_PageNumPageSize proto.InternalMessageInfo

func (m *PageNumPageSize) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *PageNumPageSize) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type PageMap struct {
	PageNum              int32    `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageMap) Reset()         { *m = PageMap{} }
func (m *PageMap) String() string { return proto.CompactTextString(m) }
func (*PageMap) ProtoMessage()    {}
func (*PageMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *PageMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageMap.Unmarshal(m, b)
}
func (m *PageMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageMap.Marshal(b, m, deterministic)
}
func (m *PageMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageMap.Merge(m, src)
}
func (m *PageMap) XXX_Size() int {
	return xxx_messageInfo_PageMap.Size(m)
}
func (m *PageMap) XXX_DiscardUnknown() {
	xxx_messageInfo_PageMap.DiscardUnknown(m)
}

var xxx_messageInfo_PageMap proto.InternalMessageInfo

func (m *PageMap) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *PageMap) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PageMap) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Token struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Token) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Users)(nil), "go.micro.srv.user.Users")
	proto.RegisterType((*UserId)(nil), "go.micro.srv.user.UserId")
	proto.RegisterType((*Ids)(nil), "go.micro.srv.user.Ids")
	proto.RegisterType((*Usernames)(nil), "go.micro.srv.user.Usernames")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*IdResponse)(nil), "go.micro.srv.user.IdResponse")
	proto.RegisterType((*UserResponse)(nil), "go.micro.srv.user.UserResponse")
	proto.RegisterType((*UserListResponse)(nil), "go.micro.srv.user.UserListResponse")
	proto.RegisterType((*PageNumPageSize)(nil), "go.micro.srv.user.PageNumPageSize")
	proto.RegisterType((*PageMap)(nil), "go.micro.srv.user.PageMap")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0xe3, 0x9f, 0xa4, 0x93, 0x2f, 0x1f, 0xed, 0xaa, 0x50, 0x13, 0x04, 0x44, 0xcb, 0x4d,
	0x24, 0x84, 0x85, 0x82, 0xc4, 0x0d, 0x42, 0xa2, 0x11, 0x10, 0x45, 0x6a, 0x4a, 0xe4, 0x26, 0x48,
	0x5c, 0x1a, 0x7b, 0x15, 0xac, 0x36, 0x5e, 0xcb, 0xbb, 0x0e, 0x0a, 0x2f, 0xc2, 0x3b, 0xf0, 0x56,
	0xbc, 0x09, 0x9a, 0x5d, 0x27, 0x0d, 0x10, 0xe7, 0x87, 0x2b, 0xef, 0xd9, 0x33, 0x7b, 0x3c, 0x73,
	0x66, 0xd6, 0x06, 0xc8, 0x05, 0xcb, 0xbc, 0x34, 0xe3, 0x92, 0x93, 0x93, 0x29, 0xf7, 0x66, 0x71,
	0x98, 0x71, 0x4f, 0x64, 0x73, 0x0f, 0x09, 0xfa, 0xc3, 0x00, 0x6b, 0x22, 0x58, 0x46, 0xfe, 0x87,
	0x6a, 0x1c, 0xb9, 0x46, 0xdb, 0xe8, 0x34, 0xfd, 0x6a, 0x1c, 0x91, 0x16, 0xd4, 0x31, 0x20, 0x09,
	0x66, 0xcc, 0xad, 0xb6, 0x8d, 0xce, 0x91, 0xbf, 0xc2, 0xc8, 0x25, 0x71, 0x78, 0xad, 0x38, 0x53,
	0x73, 0x4b, 0x8c, 0x5c, 0x1a, 0x08, 0xf1, 0x95, 0x67, 0x91, 0x6b, 0x69, 0x6e, 0x89, 0xc9, 0x23,
	0x80, 0x30, 0x63, 0x81, 0x64, 0xe3, 0x78, 0xc6, 0x5c, 0x5b, 0xb1, 0x6b, 0x3b, 0xc8, 0xe7, 0x69,
	0xb4, 0xe4, 0x1d, 0xcd, 0xdf, 0xee, 0xd0, 0x97, 0x60, 0x63, 0xae, 0x82, 0x3c, 0x03, 0x1b, 0x93,
	0x11, 0xae, 0xd1, 0x36, 0x3b, 0x8d, 0xee, 0x99, 0xf7, 0x57, 0x61, 0x1e, 0x06, 0xfa, 0x3a, 0x8a,
	0xba, 0xe0, 0x20, 0x1c, 0x44, 0x7f, 0x56, 0x49, 0xcf, 0xc0, 0x1c, 0x44, 0x82, 0x1c, 0x83, 0x19,
	0x47, 0x5a, 0xad, 0xe9, 0xe3, 0x92, 0x0e, 0xe1, 0x68, 0x52, 0x94, 0x2b, 0x08, 0x01, 0x2b, 0xe4,
	0x11, 0x53, 0xe7, 0x6c, 0x5f, 0xad, 0xf1, 0xc8, 0x4c, 0x4c, 0x0b, 0x6b, 0x70, 0xf9, 0x9b, 0x63,
	0x66, 0xdb, 0x5c, 0x77, 0x8c, 0x3e, 0x87, 0xba, 0xcf, 0x44, 0xca, 0x13, 0xc1, 0xf6, 0x53, 0xa3,
	0x3d, 0x80, 0x41, 0x74, 0xd8, 0x99, 0xa2, 0x3a, 0x73, 0x55, 0x5d, 0x00, 0xff, 0x29, 0x1b, 0x0e,
	0x53, 0x79, 0x0a, 0x16, 0xe6, 0xad, 0x74, 0xb6, 0x78, 0xab, 0x82, 0xe8, 0x77, 0x03, 0x8e, 0x11,
	0x5e, 0xc4, 0x42, 0x1e, 0xf8, 0x9e, 0x55, 0x13, 0xcd, 0x7d, 0x9a, 0x48, 0x3c, 0xb0, 0xd2, 0x60,
	0xca, 0xd4, 0x50, 0x35, 0xba, 0xad, 0x0d, 0xd1, 0xa3, 0x60, 0xca, 0x86, 0x41, 0xea, 0xab, 0x38,
	0xda, 0x87, 0x3b, 0xb8, 0x71, 0x99, 0xcf, 0xf0, 0x71, 0x15, 0x7f, 0x63, 0xc4, 0x85, 0x5a, 0xaa,
	0xb7, 0x8a, 0xd4, 0x96, 0x50, 0x4f, 0xad, 0x8e, 0x52, 0x29, 0xda, 0xfe, 0x0a, 0xd3, 0x09, 0xd4,
	0x0a, 0xe5, 0x7f, 0x13, 0x20, 0xa7, 0x60, 0x87, 0x3c, 0x4f, 0xa4, 0x72, 0xd4, 0xf6, 0x35, 0xa0,
	0x9f, 0xc0, 0x1e, 0xf3, 0x6b, 0x96, 0xec, 0xe9, 0xd6, 0x29, 0xd8, 0x12, 0xc3, 0x8b, 0x0b, 0xa7,
	0x01, 0xee, 0xce, 0x83, 0x9b, 0x58, 0x5f, 0xb5, 0xba, 0xaf, 0x41, 0xf7, 0xa7, 0x0d, 0x0d, 0xb4,
	0xee, 0x8a, 0x65, 0xf3, 0x38, 0x64, 0xe4, 0x3d, 0x38, 0x7d, 0x26, 0x3f, 0x24, 0x8c, 0xdc, 0x2f,
	0x31, 0x79, 0x10, 0xb5, 0x1e, 0x97, 0xf9, 0x5f, 0x74, 0x95, 0x56, 0xc8, 0x18, 0x6a, 0x7d, 0x26,
	0x87, 0x41, 0xb2, 0x20, 0xb4, 0xc4, 0xff, 0x35, 0xbb, 0x5b, 0x4f, 0x4a, 0x14, 0xd7, 0x67, 0x85,
	0x56, 0xc8, 0x1b, 0xb0, 0x46, 0x5c, 0x48, 0x52, 0x36, 0x00, 0xad, 0x87, 0x1b, 0x88, 0xdb, 0xbb,
	0x41, 0x2b, 0xe4, 0x35, 0x98, 0xa3, 0x7c, 0x8b, 0xc0, 0x83, 0x0d, 0xc4, 0xda, 0xf1, 0x1e, 0x38,
	0x6f, 0xd9, 0x0d, 0x93, 0x5b, 0xed, 0xd9, 0xa1, 0xf1, 0x0a, 0xac, 0xf3, 0x5c, 0x7e, 0x29, 0xcf,
	0xc1, 0xdd, 0x40, 0xa8, 0xfe, 0xd3, 0x0a, 0x39, 0x07, 0xe7, 0x82, 0x4f, 0x79, 0x2e, 0x49, 0x69,
	0xd4, 0xae, 0xf7, 0xbf, 0x83, 0xe6, 0x47, 0xec, 0x3d, 0x7e, 0x2a, 0xd5, 0x64, 0x94, 0x2b, 0x6d,
	0xcb, 0x64, 0x0c, 0x77, 0xfb, 0x4c, 0xea, 0xd9, 0x11, 0x22, 0xe6, 0x49, 0x6f, 0xb1, 0x4b, 0x6e,
	0x8f, 0xb9, 0xf1, 0xe1, 0xa4, 0xcf, 0xe4, 0x65, 0xf1, 0x8b, 0x10, 0xbd, 0x05, 0x7e, 0x73, 0xef,
	0x6d, 0xec, 0xaa, 0xd8, 0x73, 0x6a, 0x3e, 0x3b, 0xea, 0x97, 0xf6, 0xe2, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x5f, 0xb4, 0x72, 0xe0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	//    rpc Create (User) returns (Response) {
	//    }
	GetOne(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserResponse, error)
	GetMany(ctx context.Context, in *PageNumPageSize, opts ...client.CallOption) (*UserListResponse, error)
	Post(ctx context.Context, in *User, opts ...client.CallOption) (*IdResponse, error)
	Put(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *UserId, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	Logout(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	GetUserSessionByToken(ctx context.Context, in *Token, opts ...client.CallOption) (*UserResponse, error)
	GetNicknamesByIds(ctx context.Context, in *Ids, opts ...client.CallOption) (*UserListResponse, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) GetOne(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetOne", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMany(ctx context.Context, in *PageNumPageSize, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetMany", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Post(ctx context.Context, in *User, opts ...client.CallOption) (*IdResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Post", in)
	out := new(IdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Put(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Put", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Logout", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserSessionByToken(ctx context.Context, in *Token, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetUserSessionByToken", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetNicknamesByIds(ctx context.Context, in *Ids, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetNicknamesByIds", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//    rpc Create (User) returns (Response) {
	//    }
	GetOne(context.Context, *UserId, *UserResponse) error
	GetMany(context.Context, *PageNumPageSize, *UserListResponse) error
	Post(context.Context, *User, *IdResponse) error
	Put(context.Context, *User, *Response) error
	Delete(context.Context, *UserId, *Response) error
	Auth(context.Context, *User, *Token) error
	Logout(context.Context, *Token, *Response) error
	ValidateToken(context.Context, *Token, *Token) error
	GetUserSessionByToken(context.Context, *Token, *UserResponse) error
	GetNicknamesByIds(context.Context, *Ids, *UserListResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) GetOne(ctx context.Context, in *UserId, out *UserResponse) error {
	return h.UserServiceHandler.GetOne(ctx, in, out)
}

func (h *UserService) GetMany(ctx context.Context, in *PageNumPageSize, out *UserListResponse) error {
	return h.UserServiceHandler.GetMany(ctx, in, out)
}

func (h *UserService) Post(ctx context.Context, in *User, out *IdResponse) error {
	return h.UserServiceHandler.Post(ctx, in, out)
}

func (h *UserService) Put(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Put(ctx, in, out)
}

func (h *UserService) Delete(ctx context.Context, in *UserId, out *Response) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) Logout(ctx context.Context, in *Token, out *Response) error {
	return h.UserServiceHandler.Logout(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func (h *UserService) GetUserSessionByToken(ctx context.Context, in *Token, out *UserResponse) error {
	return h.UserServiceHandler.GetUserSessionByToken(ctx, in, out)
}

func (h *UserService) GetNicknamesByIds(ctx context.Context, in *Ids, out *UserListResponse) error {
	return h.UserServiceHandler.GetNicknamesByIds(ctx, in, out)
}
