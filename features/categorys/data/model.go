package data

import (
	category "altaproject/features/categorys"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}

func (data *Category) toCore() category.Core {
	return category.Core{
		Name: data.Name,
	}
}

func toCoreList(data []Category) []category.Core {
	result := []category.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
