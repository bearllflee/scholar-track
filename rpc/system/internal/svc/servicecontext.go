package svc

import (
	"github.com/bearllflee/scholar-track/rpc/system/internal/config"
	"github.com/bearllflee/scholar-track/rpc/system/internal/service"
)

type ServiceContext struct {
	Config        config.Config
	CasbinService *service.CasbinService
	ApiService    *service.ApiService
	RoleService   *service.RoleService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CasbinService: &service.CasbinService{},
		ApiService:    &service.ApiService{},
		RoleService:   &service.RoleService{},
	}
}
