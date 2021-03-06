// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: arith.proto

package arith

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Arith service

type ArithService interface {
	Mul(ctx context.Context, in *Args, opts ...client.CallOption) (*Reply, error)
}

type arithService struct {
	c    client.Client
	name string
}

func NewArithService(name string, c client.Client) ArithService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "arith"
	}
	return &arithService{
		c:    c,
		name: name,
	}
}

func (c *arithService) Mul(ctx context.Context, in *Args, opts ...client.CallOption) (*Reply, error) {
	req := c.c.NewRequest(c.name, "Arith.Mul", in)
	out := new(Reply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Arith service

type ArithHandler interface {
	Mul(context.Context, *Args, *Reply) error
}

func RegisterArithHandler(s server.Server, hdlr ArithHandler, opts ...server.HandlerOption) error {
	type arith interface {
		Mul(ctx context.Context, in *Args, out *Reply) error
	}
	type Arith struct {
		arith
	}
	h := &arithHandler{hdlr}
	return s.Handle(s.NewHandler(&Arith{h}, opts...))
}

type arithHandler struct {
	ArithHandler
}

func (h *arithHandler) Mul(ctx context.Context, in *Args, out *Reply) error {
	return h.ArithHandler.Mul(ctx, in, out)
}
