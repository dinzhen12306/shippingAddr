package logic

import (
	"context"
	"github.com/dinzhen12306/shippingAddr/models/mysql"

	"github.com/dinzhen12306/shippingAddr/internal/svc"
	"github.com/dinzhen12306/shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShippingAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShippingAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShippingAddrLogic {
	return &GetShippingAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShippingAddrLogic) GetShippingAddr(in *shippingAddrs.GetShippingAddrReq) (*shippingAddrs.GetShippingAddrResp, error) {
	address := mysql.NewAddress()
	which := make(map[string]interface{}, 0)
	for k, v := range in.Where {
		which[k] = v
	}
	err := address.Get(l.svcCtx, which)
	if err != nil {
		return nil, err
	}
	return &shippingAddrs.GetShippingAddrResp{
		Data: mysql.SqlToPb(address),
	}, nil
}
