package model

import (
	"time"

	"github.com/bearllflee/scholar-track/pkg/global"
)

type AchieveBasic struct {
	global.StModel
	AchieveCode string    `json:"achieveCode" form:"achieveCode" gorm:"column:achieve_code;comment:成果编号"`
	CategoryId  int32     `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:成果类别"`
	AwardLevel  string    `json:"awardLevel" form:"awardLevel" gorm:"column:award_level;comment:获奖级别"`
	AwardRank   string    `json:"awardRank" form:"awardRank" gorm:"column:award_rank;comment:获奖排名"`
	AwardTime   time.Time `json:"awardTime" form:"awardTime" gorm:"column:award_time;comment:获奖时间"`
	Status      int32     `json:"status" form:"status" gorm:"column:status;comment:成果状态"`
	Name        string    `json:"name" form:"name" gorm:"column:name;comment:成果名称"`
	Description string    `json:"description" form:"description" gorm:"column:description;comment:成果描述"`
	OtherInfo   []PropertyValue `json:"otherInfo" form:"otherInfo" gorm:"foreignKey:BasicId;"`
}

func (AchieveBasic) TableName() string {
	return "st_achieve_basic"
}
