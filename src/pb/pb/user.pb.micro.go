// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
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

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetOne(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetOne", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetMany(ctx context.Context, in *PageNumPageSize, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetMany", in)
	out := new(UserListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Post(ctx context.Context, in *User, opts ...client.CallOption) (*IdResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Post", in)
	out := new(IdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Put(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Put", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *UserId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Logout(ctx context.Context, in *Token, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Logout", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserSessionByToken(ctx context.Context, in *Token, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUserSessionByToken", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetNicknamesByIds(ctx context.Context, in *Ids, opts ...client.CallOption) (*UserListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetNicknamesByIds", in)
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

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		GetOne(ctx context.Context, in *UserId, out *UserResponse) error
		GetMany(ctx context.Context, in *PageNumPageSize, out *UserListResponse) error
		Post(ctx context.Context, in *User, out *IdResponse) error
		Put(ctx context.Context, in *User, out *Response) error
		Delete(ctx context.Context, in *UserId, out *Response) error
		Auth(ctx context.Context, in *User, out *Token) error
		Logout(ctx context.Context, in *Token, out *Response) error
		ValidateToken(ctx context.Context, in *Token, out *Token) error
		GetUserSessionByToken(ctx context.Context, in *Token, out *UserResponse) error
		GetNicknamesByIds(ctx context.Context, in *Ids, out *UserListResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetOne(ctx context.Context, in *UserId, out *UserResponse) error {
	return h.UserServiceHandler.GetOne(ctx, in, out)
}

func (h *userServiceHandler) GetMany(ctx context.Context, in *PageNumPageSize, out *UserListResponse) error {
	return h.UserServiceHandler.GetMany(ctx, in, out)
}

func (h *userServiceHandler) Post(ctx context.Context, in *User, out *IdResponse) error {
	return h.UserServiceHandler.Post(ctx, in, out)
}

func (h *userServiceHandler) Put(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Put(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *UserId, out *Response) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *userServiceHandler) Logout(ctx context.Context, in *Token, out *Response) error {
	return h.UserServiceHandler.Logout(ctx, in, out)
}

func (h *userServiceHandler) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func (h *userServiceHandler) GetUserSessionByToken(ctx context.Context, in *Token, out *UserResponse) error {
	return h.UserServiceHandler.GetUserSessionByToken(ctx, in, out)
}

func (h *userServiceHandler) GetNicknamesByIds(ctx context.Context, in *Ids, out *UserListResponse) error {
	return h.UserServiceHandler.GetNicknamesByIds(ctx, in, out)
}
