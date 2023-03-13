package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/model"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(model.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt >= 1 {
		err = errors.New("该邮箱已被注册")
		return nil, err
	}
	//该邮箱在数据库中不存在
	code := helper.RandCode()
	err = l.svcCtx.RDB.Set(l.ctx, req.Email, code, define.TimeExpired).Err()
	if err != nil {
		return nil, err
	}
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return

}
