package data

import (
	shoppingcarts "altaproject/features/shoppingCarts"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	TotalQuantity uint   `json:"total_quantity" form:"total_quantity"`
	TotalPrice    uint   `json:"total_price" form:"total_price"`
	Status        string `json:"status" form:"status"`
	UserID        uint   `json:"user_id" form:"user_id"`
	User          User
}

type User struct {
	gorm.Model
	UserName     string `json:"user_name" form:"user_name"`
	ShoppingCart []ShoppingCart
}

func (data *ShoppingCart) toCore() shoppingcarts.Core {
	return shoppingcarts.Core{
		ID:            int(data.ID),
		TotalQuantity: data.TotalQuantity,
		TotalPrice:    data.TotalPrice,
		Status:        data.Status,
		User: shoppingcarts.User{
			ID:       int(data.User.ID),
			UserName: data.User.UserName,
		},
	}
}

func toCoreList(data []ShoppingCart) []shoppingcarts.Core {
	result := []shoppingcarts.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core shoppingcarts.Core) ShoppingCart {
	return ShoppingCart{
		TotalQuantity: core.TotalQuantity,
		TotalPrice:    core.TotalPrice,
		Status:        core.Status,
		UserID:        uint(core.User.ID),
	}
}
