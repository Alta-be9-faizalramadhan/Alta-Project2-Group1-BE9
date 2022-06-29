package response

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	"time"
)

type ShoppingCartDetail struct {
	ID              int          `json:"id" form:"id"`
	QuantityBuyBook uint         `json:"quantity_buy_book" form:"quantity_buy_book"`
	TotalPriceBook  uint         `json:"total_price_book" form:"total_price_book"`
	CreatedAt       time.Time    `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at" form:"updated_at"`
	Book            Book         `json:"book" form:"book"`
	ShoppingCart    ShoppingCart `json:"shoppingcart" form:"shoppingcart"`
}

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}

type ShoppingCart struct {
	ID     int  `json:"id"`
	UserID uint `json:"user_id"`
}

func FromCore(data shoppingcartdetails.Core) ShoppingCartDetail {
	return ShoppingCartDetail{
		ID:              data.ID,
		QuantityBuyBook: data.QuantityBuyBook,
		TotalPriceBook:  data.TotalPriceBook,
		CreatedAt:       data.CreatedAt,
		Book: Book{
			ID:    data.Book.ID,
			Title: data.Book.Title,
			Price: data.Book.Price,
		},
		ShoppingCart: ShoppingCart{
			ID:     data.ShoppingCart.ID,
			UserID: data.ShoppingCart.UserID,
		},
	}
}

func FromCoreList(data []shoppingcartdetails.Core) []ShoppingCartDetail {
	result := []ShoppingCartDetail{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
