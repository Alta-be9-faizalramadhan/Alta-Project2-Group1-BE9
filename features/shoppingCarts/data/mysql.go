package data

import (
	shoppingcarts "altaproject/features/shoppingCarts"
	"fmt"

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

func (repo *mysqlShoppingCartRepository) InsertNewCart(data shoppingcarts.Core) (int, error) {
	cart := fromCore(data)
	result := repo.db.Create(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(result.RowsAffected), nil
}
