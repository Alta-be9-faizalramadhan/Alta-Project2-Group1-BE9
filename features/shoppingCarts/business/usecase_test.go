package business

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	shoppingcarts "altaproject/features/shoppingCarts"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockShoppingCartData struct{}

func (repo mockShoppingCartData) SelectAllOrder(id int, limit int, offset int) ([]shoppingcarts.Core, error) {
	return []shoppingcarts.Core{{ID: 1, Status: "Done", TotalQuantity: 1, TotalPrice: 150000, User: shoppingcarts.User{ID: 1, UserName: "Aldi"}}}, nil
}
func (repo mockShoppingCartData) SelectOrder(idUser int) (shoppingcarts.Core, error) {
	return shoppingcarts.Core{ID: 1, Status: "Wish List", TotalQuantity: 1, TotalPrice: 150000, User: shoppingcarts.User{ID: 1, UserName: "Aldi"}}, nil
}
func (repo mockShoppingCartData) InsertNewCart(data shoppingcarts.Core) (int, int, error) {
	return 1, 1, nil
}
func (repo mockShoppingCartData) UpdatedStatusCart(id int, status string) (int, error) {
	return 1, nil
}
func (repo mockShoppingCartData) UpdatedCart(idUser int, data shoppingcarts.Core) (shoppingcarts.Core, int, error) {
	return shoppingcarts.Core{ID: 1, Status: "Wish List", TotalQuantity: 1, TotalPrice: 150000, User: shoppingcarts.User{ID: 1, UserName: "Aldi"}}, 1, nil
}
func (repo mockShoppingCartData) IsCartNotExist(id int) (bool, shoppingcarts.Core) {
	return false, shoppingcarts.Core{ID: 1, Status: "Wish List", TotalQuantity: 1, TotalPrice: 150000, User: shoppingcarts.User{ID: 1, UserName: "Aldi"}}
}

type mockShoppingCartDataFailed struct{}

func (repo mockShoppingCartDataFailed) SelectAllOrder(id int, limit int, offset int) ([]shoppingcarts.Core, error) {
	return nil, fmt.Errorf("Failed to select data")
}
func (repo mockShoppingCartDataFailed) SelectOrder(idUser int) (shoppingcarts.Core, error) {
	return shoppingcarts.Core{}, fmt.Errorf("Failed to select data")
}
func (repo mockShoppingCartDataFailed) InsertNewCart(data shoppingcarts.Core) (int, int, error) {
	return 0, 0, fmt.Errorf("Failed to insert data")
}
func (repo mockShoppingCartDataFailed) UpdatedStatusCart(id int, status string) (int, error) {
	return 0, fmt.Errorf("Failed to updated data")
}
func (repo mockShoppingCartDataFailed) UpdatedCart(idUser int, data shoppingcarts.Core) (shoppingcarts.Core, int, error) {
	return shoppingcarts.Core{}, 0, fmt.Errorf("Failed to updated data")
}
func (repo mockShoppingCartDataFailed) IsCartNotExist(id int) (bool, shoppingcarts.Core) {
	return true, shoppingcarts.Core{TotalQuantity: 0, TotalPrice: 0}
}

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
func (repo mockShoppingCartDetailData) FindIDCart(idUser int) (int, error) {
	return 1, nil
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
func (repo mockShoppingCartDetailDataFailed) FindIDCart(idUser int) (int, error) {
	return 0, fmt.Errorf("Failed to find ID Cart")
}
func TestGetHistoryOrder(t *testing.T) {
	t.Run("Test Get History Order Success", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartData{}, mockShoppingCartDetailData{})
		result, err := shoppingCartBusiness.GetHistoryOrder(1, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, "Done", result[0].Status)
	})

	t.Run("Test Get History Order Failed", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartDataFailed{}, mockShoppingCartDetailDataFailed{})
		result, err := shoppingCartBusiness.GetHistoryOrder(1, 0, 0)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdateStatusCart(t *testing.T) {
	t.Run("Test Update Status Success", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartData{}, mockShoppingCartDetailData{})
		result, err := shoppingCartBusiness.UpdatedStatusCart(1, "Done")
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Update Status Failed", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartDataFailed{}, mockShoppingCartDetailDataFailed{})
		result, err := shoppingCartBusiness.UpdatedStatusCart(1, "Done")
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestCreateCart(t *testing.T) {
	t.Run("Test Create Cart When Cart is Not Empty Success", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartData{}, mockShoppingCartDetailData{})
		var product = shoppingcarts.Core{}
		idCart, rowSC, err := shoppingCartBusiness.CreateCart(1, 1, product)
		assert.Nil(t, err)
		assert.Equal(t, 1, rowSC)
		assert.Equal(t, 1, idCart)
	})
	t.Run("Test Create Cart When Cart is Empty Success", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartData{}, mockShoppingCartDetailData{})
		var product = shoppingcarts.Core{}
		idCart, rowSC, err := shoppingCartBusiness.CreateCart(1, 1, product)
		assert.Nil(t, err)
		assert.Equal(t, 1, rowSC)
		assert.Equal(t, 1, idCart)
	})
	t.Run("Test Create Cart When Cart is Not Empty Failed", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartData{}, mockShoppingCartDetailDataFailed{})
		var product = shoppingcarts.Core{}
		idCart, rowSC, err := shoppingCartBusiness.CreateCart(1, 1, product)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rowSC)
		assert.Equal(t, 0, idCart)
	})
	t.Run("Test Create Cart When Cart is Empty Failed", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartDataFailed{}, mockShoppingCartDetailDataFailed{})
		var product = shoppingcarts.Core{}
		idCart, rowSC, err := shoppingCartBusiness.CreateCart(1, 1, product)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rowSC)
		assert.Equal(t, 0, idCart)
	})
}

func TestUpdatedCart(t *testing.T) {
	t.Run("Test Update Cart Success", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartData{}, mockShoppingCartDetailData{})
		rowSC, err := shoppingCartBusiness.UpdatedCart(1, 1, 1, 1, 150000)
		assert.Nil(t, err)
		assert.Equal(t, 1, rowSC)
	})
	t.Run("Test Update Cart Failed", func(t *testing.T) {
		shoppingCartBusiness := NewShoppingCartBusiness(mockShoppingCartDataFailed{}, mockShoppingCartDetailDataFailed{})
		rowSC, err := shoppingCartBusiness.UpdatedCart(1, 1, 1, 1, 150000)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rowSC)
	})
}
