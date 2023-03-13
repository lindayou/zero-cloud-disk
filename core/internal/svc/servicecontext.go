package svc

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/model"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.Mysql.DataSource),
		RDB:    model.InitRedis(c.Redis.Addr),
	}
}
