package data

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	"fmt"

	"gorm.io/gorm"
)

type mysqlShoppingCartDetailRepository struct {
	db *gorm.DB
}

func NewShoppingCartDetailRepository(conn *gorm.DB) shoppingcartdetails.Data {
	return &mysqlShoppingCartDetailRepository{
		db: conn,
	}
}

func (repo *mysqlShoppingCartDetailRepository) SelectAllCartDetails(id, limit, offset int) (response []shoppingcartdetails.Core, err error) {
	var dataShoppingCartDetails []ShoppingCartDetail
	result := repo.db.Model(&ShoppingCartDetail{}).Preload("Book").Preload("ShoppingCart").Find(&dataShoppingCartDetails, "shopping_cart_id = ?", id)
	if result.Error != nil {
		return []shoppingcartdetails.Core{}, result.Error
	}
	return toCoreList(dataShoppingCartDetails), nil
}

func (repo *mysqlShoppingCartDetailRepository) InsertCartDetails(data shoppingcartdetails.Core) (row int, err error) {
	shoppingcartdetail := fromCore(data)
	result := repo.db.Model(&ShoppingCartDetail{}).Create(&shoppingcartdetail)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartDetailRepository) DeleteCartDetails(idCart int) (row int, err error) {
	var dataDetailCart ShoppingCartDetail
	result := repo.db.Delete(&dataDetailCart, "shopping_cart_id = ?", idCart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to delete shoppingcartdetails")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartDetailRepository) PutCartDetails(intCart int, input shoppingcartdetails.Core) (row int, err error) {
	var putData = fromCore(input)
	result := repo.db.Model(&ShoppingCartDetail{}).Where("shopping_cart_id = ?", intCart).Updates(&putData)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("failed to update shopping cart details")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartDetailRepository) IsBookNotInCartDetail(idBook int, idCart int) (cond bool) {
	result := repo.db.Model(&ShoppingCartDetail{}).Where("shopping_cart_id = ? AND book_id = ", idCart, idBook).First(&ShoppingCartDetail{})
	if result.RowsAffected == 0 {
		return true
	}
	return false
}
