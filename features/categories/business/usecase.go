package business

import category "altaproject/features/categories"

type categoryUsecase struct {
	categoryData category.Data
}

func NewCategoryBusiness(cateData category.Data) category.Business {
	return &categoryUsecase{
		categoryData: cateData,
	}
}

func (uc *categoryUsecase) GetAllCategory() (resp []category.Core, err error) {
	resp, err = uc.categoryData.SelectAllCategory()
	return resp, err
}
