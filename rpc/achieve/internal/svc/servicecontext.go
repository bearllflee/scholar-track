package svc

import (
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/config"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/initialize"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     initialize.MustNewGrom(c.DataSource),
	}
}
