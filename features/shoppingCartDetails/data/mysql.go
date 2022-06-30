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
func (repo *mysqlShoppingCartDetailRepository) SelectCartDetail(idCart int, idBook int) (shoppingcartdetails.Core, error) {
	var dataShoppingCartDetails ShoppingCartDetail
	result := repo.db.Model(&ShoppingCartDetail{}).Where("shopping_cart_id = ? AND book_id = ?", idCart, idBook).First(&dataShoppingCartDetails)
	if result.Error != nil {
		return shoppingcartdetails.Core{}, result.Error
	}
	return dataShoppingCartDetails.toCore(), nil
}

func (repo *mysqlShoppingCartDetailRepository) InsertCartDetails(data shoppingcartdetails.Core) (row int, err error) {
	shoppingcartdetail := fromCore(data)
	result := repo.db.Create(&shoppingcartdetail)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartDetailRepository) DeleteCartDetails(idCart int, idBook int) (row int, err error) {
	var dataDetailCart ShoppingCartDetail
	result := repo.db.Where("shopping_cart_id = ? AND book_id = ? ", idCart, idBook).Delete(&dataDetailCart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to delete shoppingcartdetails")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartDetailRepository) PutCartDetails(idCart int, idBook int, input shoppingcartdetails.Core) (row int, err error) {
	var putData = fromCore(input)
	result := repo.db.Model(&ShoppingCartDetail{}).Where("shopping_cart_id = ? AND book_id = ?", idCart, idBook).Updates(&putData)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("failed to update shopping cart details")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlShoppingCartDetailRepository) IsBookNotInCartDetail(idBook int, idCart int) (bool, shoppingcartdetails.Core) {
	var dataShoppingDetail ShoppingCartDetail
	result := repo.db.Where("shopping_cart_id = ? AND book_id = ?", idCart, idBook).First(&dataShoppingDetail)
	fmt.Print(result.RowsAffected)
	if result.RowsAffected == 0 {
		return true, shoppingcartdetails.Core{}
	}
	return false, dataShoppingDetail.toCore()
}
