package initialize

import (
	"github.com/bearllflee/scholar-track/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MustNewGrom(dataSource string) {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.DB = db
}
