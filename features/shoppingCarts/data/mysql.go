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
		return nil, result.Error
	}
	return toCoreList(dataShoppingCart), nil
}

func (repo *mysqlShoppingCartRepository) SelectOrder(idUser int) (shoppingcarts.Core, error) {
	var dataShoppingCart ShoppingCart
	result := repo.db.Model(&ShoppingCart{}).Where("user_id = ? AND status = ?", idUser, "wish list").First(&dataShoppingCart)
	if result.Error != nil {
		return shoppingcarts.Core{}, result.Error
	}
	return dataShoppingCart.toCore(), nil
}

func (repo *mysqlShoppingCartRepository) InsertNewCart(data shoppingcarts.Core) (int, int, error) {
	cart := fromCore(data)
	result := repo.db.Create(&cart)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, 0, fmt.Errorf("failed to insert data")
	}
	return int(cart.ID), int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartRepository) UpdatedStatusCart(id int, status string) (int, error) {
	result := repo.db.Model(&ShoppingCart{}).Where("status = ? AND user_id = ?", "Wish List", id).First(&ShoppingCart{}).Update("status", status)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to update status")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartRepository) UpdatedCart(idUser int, data shoppingcarts.Core) (shoppingcarts.Core, int, error) {
	var dataShoppingCart = fromCore(data)
	result := repo.db.Model(&ShoppingCart{}).Where("status = ? AND user_id = ?", "Wish List", idUser).Updates(&dataShoppingCart).First(&dataShoppingCart)
	if result.Error != nil {
		return dataShoppingCart.toCore(), 0, result.Error
	}
	if result.RowsAffected != 1 {
		return dataShoppingCart.toCore(), 0, fmt.Errorf("failed to update cart")
	}
	return dataShoppingCart.toCore(), int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartRepository) IsCartNotExist(id int) (bool, shoppingcarts.Core) {
	var dataShoppingCart ShoppingCart
	result := repo.db.Model(&ShoppingCart{}).Where("status = ? AND user_id = ?", "Wish List", id).First(&dataShoppingCart)
	if result.RowsAffected == 0 {
		return true, shoppingcarts.Core{TotalQuantity: 0, TotalPrice: 0}
	}
	return false, dataShoppingCart.toCore()
}
