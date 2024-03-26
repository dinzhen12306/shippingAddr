package logic

import (
	"context"
	"github.com/dinzhen12306/shippingAddr/models/mysql"

	"github.com/dinzhen12306/shippingAddr/internal/svc"
	"github.com/dinzhen12306/shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShippingAddrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateShippingAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShippingAddrLogic {
	return &CreateShippingAddrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateShippingAddrLogic) CreateShippingAddr(in *shippingAddrs.CreateShippingAddrReq) (*shippingAddrs.CreateShippingAddrResp, error) {
	address := mysql.PbToSql(in.Data)
	err := address.Create(l.svcCtx)
	if err != nil {
		return nil, err
	}
	return &shippingAddrs.CreateShippingAddrResp{
		ID: int64(address.ID),
	}, nil
}
