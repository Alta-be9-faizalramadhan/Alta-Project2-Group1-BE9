package response

import category "altaproject/features/categories"

type Category struct {
	Name string `json:"name" form:"name"`
}

func FromCore(data category.Core) Category {
	return Category{
		Name: data.Name,
	}
}

func FromCoreList(data []category.Core) []Category {
	result := []Category{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
