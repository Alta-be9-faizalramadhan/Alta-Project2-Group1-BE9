package business

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockShoppingCartDetailData struct{}

func (repo mockShoppingCartDetailData) InsertCartDetails(data shoppingcartdetails.Core) (row int, err error) {
	return 1, nil
}
func (repo mockShoppingCartDetailData) SelectAllCartDetails(id, limit, offset int) ([]shoppingcartdetails.Core, error) {
	return []shoppingcartdetails.Core{{ID: 1, Book: shoppingcartdetails.Book{ID: 1, Title: "Harry Potter", Price: 150000}, QuantityBuyBook: 1, TotalPriceBook: 150000, ShoppingCart: shoppingcartdetails.ShoppingCart{ID: 1, UserID: 1}}}, nil
}
func (repo mockShoppingCartDetailData) SelectCartDetail(idCart int, idBook int) (shoppingcartdetails.Core, error) {
	return shoppingcartdetails.Core{ID: 1, Book: shoppingcartdetails.Book{ID: 1, Title: "Harry Potter", Price: 150000}, QuantityBuyBook: 1, TotalPriceBook: 150000, ShoppingCart: shoppingcartdetails.ShoppingCart{ID: 1, UserID: 1}}, nil
}
func (repo mockShoppingCartDetailData) DeleteCartDetails(idCart int, idBook int) (row int, err error) {
	return 1, nil
}
func (repo mockShoppingCartDetailData) PutCartDetails(idCart int, idBook int, input shoppingcartdetails.Core) (int, error) {
	return 1, nil
}
func (repo mockShoppingCartDetailData) IsBookNotInCartDetail(idBook int, idCart int) (bool, shoppingcartdetails.Core) {
	return false, shoppingcartdetails.Core{ID: 1, Book: shoppingcartdetails.Book{ID: 1, Title: "Harry Potter", Price: 150000}, QuantityBuyBook: 1, TotalPriceBook: 150000, ShoppingCart: shoppingcartdetails.ShoppingCart{ID: 1, UserID: 1}}
}

type mockShoppingCartDetailDataFailed struct{}

func (repo mockShoppingCartDetailDataFailed) InsertCartDetails(data shoppingcartdetails.Core) (row int, err error) {
	return 0, fmt.Errorf("Failed to insert data")
}
func (repo mockShoppingCartDetailDataFailed) SelectAllCartDetails(id, limit, offset int) ([]shoppingcartdetails.Core, error) {
	return nil, fmt.Errorf("Failed to select data")
}
func (repo mockShoppingCartDetailDataFailed) SelectCartDetail(idCart int, idBook int) (shoppingcartdetails.Core, error) {
	return shoppingcartdetails.Core{}, fmt.Errorf("Failed to select data")
}
func (repo mockShoppingCartDetailDataFailed) DeleteCartDetails(idCart int, idBook int) (row int, err error) {
	return 0, fmt.Errorf("Failed to delete data")
}
func (repo mockShoppingCartDetailDataFailed) PutCartDetails(idCart int, idBook int, input shoppingcartdetails.Core) (int, error) {
	return 0, fmt.Errorf("Failed to updated data")
}
func (repo mockShoppingCartDetailDataFailed) IsBookNotInCartDetail(idBook int, idCart int) (bool, shoppingcartdetails.Core) {
	return true, shoppingcartdetails.Core{}
}

func TestGetAllCartDetails(t *testing.T) {
	t.Run("Test Get All Cart Details Success", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailData{})
		result, err := shoppingCartDetailBusiness.GetAllCartDetails(1, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, result[0].ID)
	})
	t.Run("Test Get All Cart Details Failed", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		result, err := shoppingCartDetailBusiness.GetAllCartDetails(1, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestInsertCartDetails(t *testing.T) {
	t.Run("Test Insert Cart Details Success", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailData{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.InsertCartDetails(cart)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Insert Cart Details Failed", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.InsertCartDetails(cart)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("Test Insert Cart Details Failed when Quantity is not filled", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 0,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.InsertCartDetails(cart)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Insert Cart Details Failed when Price is not filled", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  0,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.InsertCartDetails(cart)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Insert Cart Details Failed when Book ID is not filled", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    0,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.InsertCartDetails(cart)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Insert Cart Details Failed when Shopping Cart is not filled", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     0,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.InsertCartDetails(cart)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}

func TestDeleteCartDetails(t *testing.T) {
	t.Run("Test Delete Details Success", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailData{})
		result, err := shoppingCartDetailBusiness.DeleteCartDetails(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Delete Details Failed", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		result, err := shoppingCartDetailBusiness.DeleteCartDetails(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestUpdateCartDetails(t *testing.T) {
	t.Run("Test Update Cart Details Success", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailData{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.UpdateCartDetails(1, 1, cart)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Update Cart Details Failed", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.UpdateCartDetails(1, 1, cart)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("Test Update Cart Details Failed when Quantity is not filled", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 0,
			TotalPriceBook:  150000,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.UpdateCartDetails(1, 1, cart)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
	t.Run("Test Update Cart Details Failed when Price is not filled", func(t *testing.T) {
		shoppingCartDetailBusiness := NewShoppingCartDetailBusiness(mockShoppingCartDetailDataFailed{})
		var cart = shoppingcartdetails.Core{
			QuantityBuyBook: 1,
			TotalPriceBook:  0,
			ShoppingCart: shoppingcartdetails.ShoppingCart{
				ID:     1,
				UserID: 1,
			},
			Book: shoppingcartdetails.Book{
				ID:    1,
				Title: "Buku Dongeng Nusantara",
				Price: 15000,
			},
		}
		result, err := shoppingCartDetailBusiness.UpdateCartDetails(1, 1, cart)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})
}
