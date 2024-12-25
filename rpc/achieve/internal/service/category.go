package service

import (
	"context"

	"github.com/bearllflee/scholar-track/rpc/achieve/internal/model"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func (c *CategoryService) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	err := c.DB.Create(category).Error
	if err != nil {
		return nil, err	
	}
	return category, nil
}

func (c *CategoryService) QueryCategoryList(ctx context.Context, name string, typeStr string, status int32, page int, pageSize int) (error, int64, []*model.Category) {
	var categories []*model.Category
	var total int64
	if name != "" {	
		err := c.DB.Model(&model.Category{}).Where("name LIKE ?", "%"+name+"%").Where("type = ?", typeStr).Where("status = ?", status).Count(&total).Error
		if err != nil {
			return err, 0, nil
		}
	}
	if typeStr != "" {
		err := c.DB.Model(&model.Category{}).Where("type = ?", typeStr).Where("status = ?", status).Count(&total).Error
		if err != nil {
			return err, 0, nil
		}
	}
	if status != 0 {
		err := c.DB.Model(&model.Category{}).Where("status = ?", status).Count(&total).Error
		if err != nil {
			return err, 0, nil
		}
	}
	// err = c.DB.Model(&model.Category{}).Where("name LIKE ?", "%"+name+"%").Where("type = ?", typeStr).Where("status = ?", status).Offset((page - 1) * pageSize).Limit(pageSize).Find(&categories).Error
	// if err != nil {
	// 	return err, 0, nil
	// }
	return nil, total, categories
}
