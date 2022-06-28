package data

import (
	shoppingcarts "altaproject/features/shoppingCarts"

	"gorm.io/gorm"
)

type mysqlShoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepository(conn *gorm.DB) shoppingcarts.Data {
	return &mysqlShoppingCartRepository{
		db: conn,
	}
}

func (repo *mysqlShoppingCartRepository) SelectAllOrder(id int, limit int, offset int) ([]shoppingcarts.Core, error) {
	var dataShoppingCart []ShoppingCart
	result := repo.db.Preload("User").Where("user_id = ? ", id).Not("status = ?", "wish list").Limit(limit).Offset(offset).Find(&dataShoppingCart)
	if result.Error != nil {
		return []shoppingcarts.Core{}, result.Error
	}
	return toCoreList(dataShoppingCart), nil
}
