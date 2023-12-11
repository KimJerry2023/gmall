package dao

import (
	"context"
	model2 "gmall/repository/db/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// ListCategory 分类列表
func (dao *CategoryDao) ListCategory() (category []*model2.Category, err error) {
	err = dao.DB.Model(&model2.Category{}).Find(&category).Error
	return
}
