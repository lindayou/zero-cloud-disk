package logic

import (
	"cloud-disk/model"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetialsLogic {
	return &UserDetialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetialsLogic) UserDetials(req *types.UserDetialRequest) (resp *types.UserDetialReply, err error) {
	// todo: add your logic here and delete this line
	uc := new(model.UserBasic)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(uc)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("查询用户不存在")
	}
	//resp = new(types.UserDetialReply)
	resp = &types.UserDetialReply{
		Name:  uc.Name,
		Email: uc.Email,
	}
	//resp.Name = uc.Name
	//resp.Email = uc.Email
	return
}
