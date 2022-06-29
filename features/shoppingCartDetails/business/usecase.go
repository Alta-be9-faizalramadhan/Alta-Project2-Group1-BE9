package business

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
)

type shoppingCartDetailUseCase struct {
	shoppingCartDetailData shoppingcartdetails.Data
}

func NewShoppingCartDetailBusiness(scdData shoppingcartdetails.Data) shoppingcartdetails.Business {
	return &shoppingCartDetailUseCase{
		shoppingCartDetailData: scdData,
	}
}

func (uc *shoppingCartDetailUseCase) GetAllCartDetails(id, limit, offset int) (resp []shoppingcartdetails.Core, err error) {
	resp, err = uc.shoppingCartDetailData.SelectAllCartDetails(id, limit, offset)
	return resp, err
}
