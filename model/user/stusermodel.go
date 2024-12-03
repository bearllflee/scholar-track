package user

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StUserModel = (*customStUserModel)(nil)

type (
	// StUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStUserModel.
	StUserModel interface {
		stUserModel
	}

	customStUserModel struct {
		*defaultStUserModel
	}
)

// NewStUserModel returns a model for the database table.
func NewStUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StUserModel {
	return &customStUserModel{
		defaultStUserModel: newStUserModel(conn, c, opts...),
	}
}
