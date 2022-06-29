package data

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"

	"gorm.io/gorm"
)

type ShoppingCartDetail struct {
	gorm.Model
	Book            Book
	ShoppingCart    ShoppingCart
	QuantityBuyBook uint `json:"quantity_buy_book" form:"quantity_buy_book"`
	TotalPriceBook  uint `json:"total_price_book" form:"total_price_book"`
	BookID          int  `json:"book_id" form:"book_id"`
	ShoppingCartID  int  `json:"shopping_cart_id" form:"shopping_cart_id"`
}

type Book struct {
	gorm.Model
	Title              string `form:"title"`
	Price              int    `form:"price"`
	ShoppingCartDetail []ShoppingCartDetail
}

type ShoppingCart struct {
	gorm.Model
	UserID             uint `form:"user_id"`
	ShoppingCartDetail []ShoppingCartDetail
}

func (data *ShoppingCartDetail) toCore() shoppingcartdetails.Core {
	return shoppingcartdetails.Core{
		ID:              int(data.ID),
		QuantityBuyBook: data.QuantityBuyBook,
		TotalPriceBook:  data.TotalPriceBook,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
		Book: shoppingcartdetails.Book{
			ID:    int(data.Book.ID),
			Title: data.Book.Title,
			Price: uint(data.Book.Price),
		},
		ShoppingCart: shoppingcartdetails.ShoppingCart{
			ID:     int(data.ShoppingCart.ID),
			UserID: data.ShoppingCart.UserID,
		},
	}
}

func toCoreList(data []ShoppingCartDetail) []shoppingcartdetails.Core {
	result := []shoppingcartdetails.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core shoppingcartdetails.Core) ShoppingCartDetail {
	return ShoppingCartDetail{
		QuantityBuyBook: core.QuantityBuyBook,
		TotalPriceBook:  core.TotalPriceBook,
		BookID:          int(core.Book.ID),
		ShoppingCartID:  int(core.ShoppingCart.ID),
	}
}
