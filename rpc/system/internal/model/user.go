package model

import (
	"github.com/bearllflee/scholar-track/pkg/global"
)

type User struct {
	global.StModel
	Username string `gorm:"type:varchar(15);not null;uniqueIndex:idx_unique_username" json:"username"`
	Email    string `gorm:"type:varchar(255);uniqueIndex:idx_unique_email" json:"email"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`
	Role     uint   `gorm:"not null" json:"role"`
	Status   int8   `gorm:"not null;default:1" json:"status"`
	Nickname string `gorm:"type:varchar(255)" json:"nickname"`
	Phone    string `gorm:"type:char(11);uniqueIndex:idx_unique_phone" json:"phone"`
	Gender   int8   `gorm:"not null;default:1" json:"gender"`
	Major    string `gorm:"type:varchar(255);not null" json:"major"`
	College  string `gorm:"type:varchar(255);not null" json:"college"`
	Grade    string `gorm:"type:char(4);not null" json:"grade"`
	Class    string `gorm:"type:varchar(10);not null" json:"class"`
	Realname string `gorm:"type:varchar(255);not null" json:"realname"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

func (User) TableName() string {
	return "st_user"
}
