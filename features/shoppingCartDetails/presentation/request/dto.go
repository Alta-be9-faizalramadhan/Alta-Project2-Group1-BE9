package request

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
)

type ShoppingCartDetail struct {
	QuantityBuyBook uint `form:"quantity_buy_book"`
	TotalPriceBook  uint `form:"total_price_book"`
	BookId          int  `form:"book_id"`
	ShoppingCartId  int  `form:"shoppingcart_id"`
}

func ToCore(req ShoppingCartDetail) shoppingcartdetails.Core {
	return shoppingcartdetails.Core{
		QuantityBuyBook: req.QuantityBuyBook,
		TotalPriceBook:  req.TotalPriceBook,
		Book: shoppingcartdetails.Book{
			ID: req.BookId,
		},
		ShoppingCart: shoppingcartdetails.ShoppingCart{
			ID: req.ShoppingCartId,
		},
	}
}
