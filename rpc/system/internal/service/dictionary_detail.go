package service

import (
	"gorm.io/gorm"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
)

type DictionaryDetailService struct {
	db *gorm.DB
}

func NewDictionaryDetailService(db *gorm.DB) *DictionaryDetailService {
	return &DictionaryDetailService{db: db}
}

func (s *DictionaryDetailService) CreateDictionaryDetail(dictionaryDetail *model.DictionaryDetail) error {
	return s.db.Create(dictionaryDetail).Error
}

func (s *DictionaryDetailService) DeleteDictionaryDetail(id uint64) error {
	return s.db.Delete(&model.DictionaryDetail{}, id).Error
}

func (s *DictionaryDetailService) UpdateDictionaryDetail(dictionaryDetail *model.DictionaryDetail) (*model.DictionaryDetail, error) {
	err := s.db.Save(dictionaryDetail).Error
	if err != nil {
		return nil, err
	}
	return dictionaryDetail, nil
}

func (s *DictionaryDetailService) QueryDictionaryDetailDetail(id uint64) (*model.DictionaryDetail, error) {
	var dictionaryDetail model.DictionaryDetail
	err := s.db.First(&dictionaryDetail, id).Error
	if err != nil {
		return nil, err
	}
	return &dictionaryDetail, nil
}

func (s *DictionaryDetailService) QueryAllDictionaryDetail() ([]*model.DictionaryDetail, error) {
	var dictionaryDetails []*model.DictionaryDetail
	err := s.db.Find(&dictionaryDetails).Error
	if err != nil {
		return nil, err
	}
	return dictionaryDetails, nil
}
