package service

import (
	"context"
	"errors"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/model"
	"gorm.io/gorm"
)

type AchieveService struct {
	db *gorm.DB
}

func NewAchieveService(db *gorm.DB) *AchieveService {
	return &AchieveService{db: db}
}

func (s *AchieveService) UploadAchieve(ctx context.Context, achieve *model.AchieveBasic) (*model.AchieveBasic, error) {
	db := s.db.WithContext(ctx).Model(&model.AchieveBasic{})
	var achieveBasic = &model.AchieveBasic{}
	err := db.Where("code = ?", achieve.Code).First(achieveBasic).Error
	// 如果存在，则返回错误
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// 创建新记录
	if err = db.Create(achieve).Error; err != nil {
		return nil, err
	}
	return achieve, nil
}

func (s *AchieveService) DeleteAchieve(ctx context.Context, id uint64) error {
	db := s.db.WithContext(ctx).Model(&model.AchieveBasic{})
	// 物理删除
	if err := db.Unscoped().Delete(&model.AchieveBasic{}, id).Error; err != nil {
		return err
	}
	return nil
}
