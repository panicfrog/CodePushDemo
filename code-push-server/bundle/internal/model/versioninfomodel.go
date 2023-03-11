package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VersionInfoModel = (*customVersionInfoModel)(nil)

type (
	// VersionInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVersionInfoModel.
	VersionInfoModel interface {
		versionInfoModel
	}

	customVersionInfoModel struct {
		*defaultVersionInfoModel
	}
)

// NewVersionInfoModel returns a model for the database table.
func NewVersionInfoModel(conn sqlx.SqlConn, c cache.CacheConf) VersionInfoModel {
	return &customVersionInfoModel{
		defaultVersionInfoModel: newVersionInfoModel(conn, c),
	}
}
