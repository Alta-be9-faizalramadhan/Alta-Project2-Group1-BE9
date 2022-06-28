package request

import (
	shoppingcarts "altaproject/features/shoppingCarts"
)

type ShoppingCart struct {
	TotalQuantity uint   `json:"total_quantity" form:"total_quantity"`
	TotalPrice    uint   `json:"total_price" form:"total_price"`
	Status        string `json:"status" form:"status"`
	UserID        int    `json:"user_id" form:"user_id"`
}

func ToCore(req ShoppingCart) shoppingcarts.Core {
	return shoppingcarts.Core{
		TotalQuantity: req.TotalQuantity,
		TotalPrice:    req.TotalPrice,
		Status:        req.Status,
		User: shoppingcarts.User{
			ID: req.UserID,
		},
	}
}
