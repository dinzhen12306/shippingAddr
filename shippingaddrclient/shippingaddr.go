// Code generated by goctl. DO NOT EDIT.
// Source: shippingAddrs.proto

package shippingaddrclient

import (
	"context"

	"github.com/dinzhen12306/shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateShippingAddrReq   = shippingAddrs.CreateShippingAddrReq
	CreateShippingAddrResp  = shippingAddrs.CreateShippingAddrResp
	DeleteShippingAddrReq   = shippingAddrs.DeleteShippingAddrReq
	DeleteShippingAddrResp  = shippingAddrs.DeleteShippingAddrResp
	GetShippingAddrListReq  = shippingAddrs.GetShippingAddrListReq
	GetShippingAddrListResp = shippingAddrs.GetShippingAddrListResp
	GetShippingAddrReq      = shippingAddrs.GetShippingAddrReq
	GetShippingAddrResp     = shippingAddrs.GetShippingAddrResp
	ShippingAddress         = shippingAddrs.ShippingAddress
	UpdateShippingAddrReq   = shippingAddrs.UpdateShippingAddrReq
	UpdateShippingAddrResp  = shippingAddrs.UpdateShippingAddrResp

	ShippingAddr interface {
		GetShippingAddrList(ctx context.Context, in *GetShippingAddrListReq, opts ...grpc.CallOption) (*GetShippingAddrListResp, error)
		GetShippingAddr(ctx context.Context, in *GetShippingAddrReq, opts ...grpc.CallOption) (*GetShippingAddrResp, error)
		CreateShippingAddr(ctx context.Context, in *CreateShippingAddrReq, opts ...grpc.CallOption) (*CreateShippingAddrResp, error)
		UpdateShippingAddr(ctx context.Context, in *UpdateShippingAddrReq, opts ...grpc.CallOption) (*UpdateShippingAddrResp, error)
		DeleteShippingAddr(ctx context.Context, in *DeleteShippingAddrReq, opts ...grpc.CallOption) (*DeleteShippingAddrResp, error)
	}

	defaultShippingAddr struct {
		cli zrpc.Client
	}
)

func NewShippingAddr(cli zrpc.Client) ShippingAddr {
	return &defaultShippingAddr{
		cli: cli,
	}
}

func (m *defaultShippingAddr) GetShippingAddrList(ctx context.Context, in *GetShippingAddrListReq, opts ...grpc.CallOption) (*GetShippingAddrListResp, error) {
	client := shippingAddrs.NewShippingAddrClient(m.cli.Conn())
	return client.GetShippingAddrList(ctx, in, opts...)
}

func (m *defaultShippingAddr) GetShippingAddr(ctx context.Context, in *GetShippingAddrReq, opts ...grpc.CallOption) (*GetShippingAddrResp, error) {
	client := shippingAddrs.NewShippingAddrClient(m.cli.Conn())
	return client.GetShippingAddr(ctx, in, opts...)
}

func (m *defaultShippingAddr) CreateShippingAddr(ctx context.Context, in *CreateShippingAddrReq, opts ...grpc.CallOption) (*CreateShippingAddrResp, error) {
	client := shippingAddrs.NewShippingAddrClient(m.cli.Conn())
	return client.CreateShippingAddr(ctx, in, opts...)
}

func (m *defaultShippingAddr) UpdateShippingAddr(ctx context.Context, in *UpdateShippingAddrReq, opts ...grpc.CallOption) (*UpdateShippingAddrResp, error) {
	client := shippingAddrs.NewShippingAddrClient(m.cli.Conn())
	return client.UpdateShippingAddr(ctx, in, opts...)
}

func (m *defaultShippingAddr) DeleteShippingAddr(ctx context.Context, in *DeleteShippingAddrReq, opts ...grpc.CallOption) (*DeleteShippingAddrResp, error) {
	client := shippingAddrs.NewShippingAddrClient(m.cli.Conn())
	return client.DeleteShippingAddr(ctx, in, opts...)
}
