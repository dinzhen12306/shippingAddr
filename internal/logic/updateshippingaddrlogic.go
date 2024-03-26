package logic

import (
	"context"
	"github.com/dinzhen12306/shippingAddr/models/mysql"

	"github.com/dinzhen12306/shippingAddr/internal/svc"
	"github.com/dinzhen12306/shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShippingAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateShippingAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShippingAddrLogic {
	return &UpdateShippingAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateShippingAddrLogic) UpdateShippingAddr(in *shippingAddrs.UpdateShippingAddrReq) (*shippingAddrs.UpdateShippingAddrResp, error) {
	address := mysql.PbToSql(in.Data)
	err := address.Update(l.svcCtx)
	if err != nil {
		return nil, err
	}
	return &shippingAddrs.UpdateShippingAddrResp{
		ID: int64(address.ID),
	}, nil
}
