// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: grpc/adm/administrator.proto

package grpc_adm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AdministratorService service

type AdministratorService interface {
	Login(ctx context.Context, in *AdministratorLoginRequestDto, opts ...client.CallOption) (*AdministratorLoginResponseDto, error)
}

type administratorService struct {
	c    client.Client
	name string
}

func NewAdministratorService(name string, c client.Client) AdministratorService {
	return &administratorService{
		c:    c,
		name: name,
	}
}

func (c *administratorService) Login(ctx context.Context, in *AdministratorLoginRequestDto, opts ...client.CallOption) (*AdministratorLoginResponseDto, error) {
	req := c.c.NewRequest(c.name, "AdministratorService.Login", in)
	out := new(AdministratorLoginResponseDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdministratorService service

type AdministratorServiceHandler interface {
	Login(context.Context, *AdministratorLoginRequestDto, *AdministratorLoginResponseDto) error
}

func RegisterAdministratorServiceHandler(s server.Server, hdlr AdministratorServiceHandler, opts ...server.HandlerOption) error {
	type administratorService interface {
		Login(ctx context.Context, in *AdministratorLoginRequestDto, out *AdministratorLoginResponseDto) error
	}
	type AdministratorService struct {
		administratorService
	}
	h := &administratorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AdministratorService{h}, opts...))
}

type administratorServiceHandler struct {
	AdministratorServiceHandler
}

func (h *administratorServiceHandler) Login(ctx context.Context, in *AdministratorLoginRequestDto, out *AdministratorLoginResponseDto) error {
	return h.AdministratorServiceHandler.Login(ctx, in, out)
}
