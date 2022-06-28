package response

import (
	shoppingcarts "altaproject/features/shoppingCarts"
	"time"
)

type ShoppingCart struct {
	ID            int       `json:"id" form:"id"`
	TotalQuantity uint      `json:"total_quantity" form:"total_quantity"`
	TotalPrice    uint      `json:"total_price" form:"total_price"`
	Status        string    `json:"status" form:"status"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
	User          User
}

type User struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
}

func FromCore(data shoppingcarts.Core) ShoppingCart {
	return ShoppingCart{
		ID:            data.ID,
		TotalQuantity: data.TotalQuantity,
		TotalPrice:    data.TotalPrice,
		Status:        data.Status,
		User: User{
			ID:       data.User.ID,
			UserName: data.User.UserName,
		},
	}
}

func FromCoreList(data []shoppingcarts.Core) []ShoppingCart {
	result := []ShoppingCart{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
