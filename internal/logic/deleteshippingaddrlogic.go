package logic

import (
	"context"
	"github.com/dinzhen12306/shippingAddr/models/mysql"

	"github.com/dinzhen12306/shippingAddr/internal/svc"
	"github.com/dinzhen12306/shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShippingAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteShippingAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShippingAddrLogic {
	return &DeleteShippingAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteShippingAddrLogic) DeleteShippingAddr(in *shippingAddrs.DeleteShippingAddrReq) (*shippingAddrs.DeleteShippingAddrResp, error) {
	address := mysql.NewAddress()
	address.ID = uint(in.ID)
	err := address.Delete(l.svcCtx)
	if err != nil {
		return nil, err
	}
	return &shippingAddrs.DeleteShippingAddrResp{}, nil
}
