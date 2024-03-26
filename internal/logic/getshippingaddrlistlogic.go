package logic

import (
	"context"
	"shippingAddr/internal/svc"
	"shippingAddr/models/mysql"
	"shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShippingAddrListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShippingAddrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShippingAddrListLogic {
	return &GetShippingAddrListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShippingAddrListLogic) GetShippingAddrList(in *shippingAddrs.GetShippingAddrListReq) (*shippingAddrs.GetShippingAddrListResp, error) {
	addr := mysql.NewAddress()
	which := make(map[string]interface{}, 0)
	for k, v := range in.Where {
		which[k] = v
	}
	gets, err := addr.Gets(l.svcCtx, which)
	if err != nil {
		return nil, err
	}
	addrList := make([]*shippingAddrs.ShippingAddress, 0)
	for _, v := range gets {
		addrList = append(addrList, mysql.SqlToPb(v))
	}
	return &shippingAddrs.GetShippingAddrListResp{
		Data: addrList,
	}, nil
}
