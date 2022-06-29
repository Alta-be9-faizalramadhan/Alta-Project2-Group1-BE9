package data

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"

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
	result := repo.db.Model(&ShoppingCartDetail{}).Preload("Book").Preload("ShoppingCart").Find(&dataShoppingCartDetails, "shoppingcart_id = ?", id)
	if result.Error != nil {
		return []shoppingcartdetails.Core{}, result.Error
	}
	return toCoreList(dataShoppingCartDetails), nil
}
