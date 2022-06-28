package business

import (
	shoppingcarts "altaproject/features/shoppingCarts"
)

type shoppingCartUsecase struct {
	shoppingCartData shoppingcarts.Data
}

func NewShoppingCartBusiness(scData shoppingcarts.Data) shoppingcarts.Business {
	return &shoppingCartUsecase{
		shoppingCartData: scData,
	}
}

func (uc *shoppingCartUsecase) GetHistoryOrder(id int, limit int, offset int) (resp []shoppingcarts.Core, err error) {
	resp, err = uc.shoppingCartData.SelectAllOrder(id, limit, offset)
	return resp, err
}

func (uc *shoppingCartUsecase) CreateNewCart(data shoppingcarts.Core) (row int, err error) {
	row, err = uc.shoppingCartData.InsertNewCart(data)
	return row, err
}

func (uc *shoppingCartUsecase) UpdatedStatusCart(id int, status string) (row int, err error) {
	row, err = uc.shoppingCartData.UpdatedStatusCart(id, status)
	return row, err
}
