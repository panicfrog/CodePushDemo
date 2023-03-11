package svc

import (
	"code-push-server/bundle/internal/config"
	"code-push-server/bundle/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	FileModel model.FileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		FileModel: model.NewFileModel(conn, c.CacheRedis),
	}
}
