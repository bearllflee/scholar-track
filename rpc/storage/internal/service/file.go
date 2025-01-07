package service

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/storage/internal/model"
	"gorm.io/gorm"
)

type FileService struct {
	db *gorm.DB
}

func (s *FileService) GetFileById(ctx context.Context, id uint64) (*model.File, error) {
	file := &model.File{}
	err := s.db.WithContext(ctx).First(file, id).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{db: db}
}

func (s *FileService) UploadFile(ctx context.Context, file *model.File) (*model.File, error) {
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(file).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return file, nil
}
