package data

import (
	category "altaproject/features/categorys"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) category.Data {
	return &mysqlCategoryRepository{
		db: conn,
	}
}

func (repo *mysqlCategoryRepository) SelectAllCategory() ([]category.Core, error) {
	var dataCategorys []Category
	result := repo.db.Find(&dataCategorys)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataCategorys), nil
}
