package service

import (
	"gorm.io/gorm"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
)

type DictionaryService struct {
	db *gorm.DB
}
func NewDictionaryService(db *gorm.DB) *DictionaryService {
	return &DictionaryService{db: db}
}

func (s *DictionaryService) CreateDictionary(dictionary *model.Dictionary) error {
	return s.db.Create(dictionary).Error
}

func (s *DictionaryService) DeleteDictionary(id uint64) error {
	return s.db.Delete(&model.Dictionary{}, id).Error
}

func (s *DictionaryService) UpdateDictionary(dictionary *model.Dictionary) (*model.Dictionary, error) {
	err := s.db.Save(dictionary).Error
	if err != nil {
		return nil, err
	}
	return dictionary, nil
}

func (s *DictionaryService) QueryDictionaryDetail(id uint64) (*model.Dictionary, error) {
	var dictionary model.Dictionary
	err := s.db.Preload("DictionaryDetails").First(&dictionary, id).Error
	if err != nil {
		return nil, err
	}
	return &dictionary, nil
}

func (s *DictionaryService) QueryAllDictionary() ([]*model.Dictionary, error) {
	var dictionaries []*model.Dictionary
	err := s.db.Preload("DictionaryDetails").Find(&dictionaries).Error
	if err != nil {
		return nil, err
	}
	return dictionaries, nil
}
