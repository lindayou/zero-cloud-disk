package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/model"
	"context"
	"errors"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// todo: add your logic here and delete this line
	//验证验证码
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取到该邮箱的验证码")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return nil, err
	}
	//判断用户名是否存在
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(model.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt >= 1 {
		err = errors.New("用户名已存在")
		return nil, err
	}
	//插入数据返回结果
	user := &model.UserBasic{
		Name:     req.Name,
		Email:    req.Email,
		Password: helper.Md5(req.Password),
		Identity: helper.GenerateUuid(),
	}
	insert, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("infect :", insert)
	return
}
