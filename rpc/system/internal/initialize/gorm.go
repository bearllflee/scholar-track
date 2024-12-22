package initialize

import (
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MustNewGrom(dataSource string) {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.DB = db
	AutoMigrate(db)
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{}, &model.Api{})
}
