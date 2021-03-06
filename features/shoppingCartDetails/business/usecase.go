package business

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	"errors"
)

type shoppingCartDetailUseCase struct {
	shoppingCartDetailData shoppingcartdetails.Data
}

func NewShoppingCartDetailBusiness(scdData shoppingcartdetails.Data) shoppingcartdetails.Business {
	return &shoppingCartDetailUseCase{
		shoppingCartDetailData: scdData,
	}
}

func (uc *shoppingCartDetailUseCase) GetAllCartDetails(idUser, limit, offset int) (resp []shoppingcartdetails.Core, err error) {
	idCart, _ := uc.shoppingCartDetailData.FindIDCart(idUser)
	if idCart == 0 {

	}
	resp, err = uc.shoppingCartDetailData.SelectAllCartDetails(idCart, limit, offset)
	return resp, err
}

func (uc *shoppingCartDetailUseCase) InsertCartDetails(input shoppingcartdetails.Core) (row int, err error) {
	if input.QuantityBuyBook == 0 || input.TotalPriceBook == 0 || input.Book.ID == 0 || input.ShoppingCart.ID == 0 {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.shoppingCartDetailData.InsertCartDetails(input)
	return row, err
}

func (uc *shoppingCartDetailUseCase) DeleteCartDetails(idCart int, idBook int) (row int, err error) {
	row, errDel := uc.shoppingCartDetailData.DeleteCartDetails(idCart, idBook)
	return row, errDel
}

func (uc *shoppingCartDetailUseCase) UpdateCartDetails(idCart int, idBook int, data shoppingcartdetails.Core) (row int, err error) {
	if data.QuantityBuyBook == 0 || data.TotalPriceBook == 0 {
		return -1, errors.New("all input data must be filled")
	}
	row, err1 := uc.shoppingCartDetailData.PutCartDetails(idCart, idBook, data)
	return row, err1
}
