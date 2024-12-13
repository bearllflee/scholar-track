package model

import (
	"github.com/bearllflee/scholar-track/pkg/global"
)

type Role struct {
	global.StModel
	RoleName string `gorm:"type:varchar(255);not null;unique" json:"roleName"`
	ParentId uint64   `gorm:"not null;default:0" json:"parentId"`
}

func (Role) TableName() string {
	return "st_role"
}
