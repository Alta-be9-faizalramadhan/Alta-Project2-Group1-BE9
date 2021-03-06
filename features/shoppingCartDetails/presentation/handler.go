package presentation

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	_requestSCD "altaproject/features/shoppingCartDetails/presentation/request"
	_responseSCD "altaproject/features/shoppingCartDetails/presentation/response"
	"altaproject/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShoppingCartDetailHandler struct {
	shoppingCartDetailBusiness shoppingcartdetails.Business
}

func NewShoppingCartDetailHandler(business shoppingcartdetails.Business) *ShoppingCartDetailHandler {
	return &ShoppingCartDetailHandler{
		shoppingCartDetailBusiness: business,
	}
}

func (h *ShoppingCartDetailHandler) GetAllCartDetails(c echo.Context) error {
	//id := c.Param("idUser")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	//idInt, errid := strconv.Atoi(id)
	// if errid != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "id not recognized",
	// 	})
	// }
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	result, err := h.shoppingCartDetailBusiness.GetAllCartDetails(idToken, limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseSCD.FromCoreList(result),
	})
}

func (h *ShoppingCartDetailHandler) InsertCartDetails(c echo.Context) error {
	quantityBuyBook := c.FormValue("quantity_buy_book")
	quantityBuyBookInt, _ := strconv.Atoi(quantityBuyBook)
	totalPriceBook := c.FormValue("total_price_book")
	totalPriceBookInt, _ := strconv.Atoi(totalPriceBook)
	bookId := c.FormValue("book_id")
	bookIdInt, _ := strconv.Atoi(bookId)
	shoppingcartId := c.FormValue("shopping_cart_id")
	shoppingCartIdInt, _ := strconv.Atoi(shoppingcartId)

	var newShoppingcartdetail = _requestSCD.ShoppingCartDetail{
		QuantityBuyBook: uint(quantityBuyBookInt),
		TotalPriceBook:  uint(totalPriceBookInt),
		BookId:          bookIdInt,
		ShoppingCartId:  shoppingCartIdInt,
	}

	dataBook := _requestSCD.ToCore(newShoppingcartdetail)
	row, err := h.shoppingCartDetailBusiness.InsertCartDetails(dataBook)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to insert data",
	})
}

func (h *ShoppingCartDetailHandler) DeleteCartDetails(c echo.Context) error {
	id := c.Param("idcart")
	idHap, errHap := strconv.Atoi(id)
	if errHap != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id shoppingcart is not recognize",
		})
	}
	idbook := c.Param("idbook")
	idBookInt, errHap := strconv.Atoi(idbook)
	if errHap != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id book is not recognize",
		})
	}
	_, err := h.shoppingCartDetailBusiness.DeleteCartDetails(idHap, idBookInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete shopping cart details",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to delete shopping cart details",
	})
}

func (h *ShoppingCartDetailHandler) UpdateCartDetails(c echo.Context) error {
	idCart := c.QueryParam("idCart")
	idCartInt, errCart := strconv.Atoi(idCart)
	if errCart != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized id cart",
		})
	}
	idBook := c.QueryParam("idBook")
	idBookInt, errBook := strconv.Atoi(idBook)
	if errBook != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized id book",
		})
	}
	quantityBuyBook := c.FormValue("quantity_buy_book")
	quantityBuyBookInt, _ := strconv.Atoi(quantityBuyBook)
	totalPriceBook := c.FormValue("total_price_book")
	totalPriceBookInt, _ := strconv.Atoi(totalPriceBook)
	// bookId := c.FormValue("book_id")
	// bookIdInt, _ := strconv.Atoi(bookId)
	// shoppingCartId := c.FormValue("shopping_cart_id")
	// shoppingCartIdInt, _ := strconv.Atoi(shoppingCartId)
	var cartDetail = _requestSCD.ShoppingCartDetail{
		QuantityBuyBook: uint(quantityBuyBookInt),
		TotalPriceBook:  uint(totalPriceBookInt),
		// BookId:          bookIdInt,
		// ShoppingCartId:  shoppingCartIdInt,
	}
	result, err := h.shoppingCartDetailBusiness.UpdateCartDetails(idCartInt, idBookInt, _requestSCD.ToCore(cartDetail))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update shopping cart details",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update shopping cart details",
		})
	}
	if result == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data shopping cart details",
	})
}
