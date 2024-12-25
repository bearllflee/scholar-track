package svc

import (
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/config"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/initialize"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/service"
)

type ServiceContext struct {
	Config config.Config
	CategoryService *service.CategoryService
	// PropertyService service.PropertyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initialize.MustNewGrom(c.DataSource)
	return &ServiceContext{
		Config: c,
		CategoryService: &service.CategoryService{
			DB: db,
		},
	}
}