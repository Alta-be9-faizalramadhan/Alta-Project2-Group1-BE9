package business

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	shoppingcarts "altaproject/features/shoppingCarts"
)

type shoppingCartUsecase struct {
	shoppingCartData       shoppingcarts.Data
	shoppingCartDetailData shoppingcartdetails.Data
}

func NewShoppingCartBusiness(scData shoppingcarts.Data, scdData shoppingcartdetails.Data) shoppingcarts.Business {
	return &shoppingCartUsecase{
		shoppingCartData:       scData,
		shoppingCartDetailData: scdData,
	}
}

func (uc *shoppingCartUsecase) GetHistoryOrder(id int, limit int, offset int) (resp []shoppingcarts.Core, err error) {
	resp, err = uc.shoppingCartData.SelectAllOrder(id, limit, offset)
	return resp, err
}

func (uc *shoppingCartUsecase) CreateCart(idUser int, idBook int, data shoppingcarts.Core) (rowSC int, errSC error) {
	cond, shoppingCart := uc.shoppingCartData.IsCartNotExist(idUser)
	if cond {
		var idShoppingCart int
		data.TotalPrice = data.TotalPrice * data.TotalQuantity
		idShoppingCart, rowSC, errSC = uc.shoppingCartData.InsertNewCart(data)
		if errSC != nil {
			return 0, errSC
		}
		var product = shoppingcartdetails.Core{
			QuantityBuyBook: data.TotalQuantity,
			TotalPriceBook:  data.TotalPrice,
			Book: shoppingcartdetails.Book{
				ID:    idBook,
				Price: (data.TotalPrice / data.TotalQuantity),
			},
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     idShoppingCart,
				UserID: uint(idUser),
			},
		}
		_, err := uc.shoppingCartDetailData.InsertCartDetails(product)
		if err != nil {
			return 0, err
		}
		return rowSC, nil
	} else {
		var quantity = data.TotalQuantity
		var price = data.TotalPrice
		var updates = shoppingcarts.Core{
			TotalQuantity: data.TotalQuantity + shoppingCart.TotalQuantity,
			TotalPrice:    shoppingCart.TotalPrice + (data.TotalPrice * data.TotalQuantity),
		}
		data, rowSC, errSC := uc.shoppingCartData.UpdatedCart(idUser, updates)
		if errSC != nil {
			return 0, errSC
		}
		cond, productDetail := uc.shoppingCartDetailData.IsBookNotInCartDetail(idBook, data.ID)

		if cond {
			var product = shoppingcartdetails.Core{
				QuantityBuyBook: quantity,
				TotalPriceBook:  price,
				Book: shoppingcartdetails.Book{
					ID:    idBook,
					Price: (data.TotalPrice / data.TotalQuantity),
				},
				ShoppingCart: shoppingcartdetails.ShoppingCart{
					ID:     data.ID,
					UserID: uint(idUser),
				},
			}
			_, err := uc.shoppingCartDetailData.InsertCartDetails(product)
			if err != nil {
				return 0, err
			}
		} else {
			var product = shoppingcartdetails.Core{
				QuantityBuyBook: productDetail.QuantityBuyBook + quantity,
				TotalPriceBook:  productDetail.TotalPriceBook + (price * quantity),
			}
			_, err := uc.shoppingCartDetailData.PutCartDetails(data.ID, idBook, product)
			if err != nil {
				return 0, err
			}
		}
		return rowSC, nil
	}
}

func (uc *shoppingCartUsecase) UpdatedStatusCart(id int, status string) (row int, err error) {
	row, err = uc.shoppingCartData.UpdatedStatusCart(id, status)
	return row, err
}
