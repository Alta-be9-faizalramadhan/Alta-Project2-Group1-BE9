package presentation

import (
	shoppingcarts "altaproject/features/shoppingCarts"
	_requestShoppingCart "altaproject/features/shoppingCarts/presentation/request"
	_responseShoppingCart "altaproject/features/shoppingCarts/presentation/response"
	"altaproject/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShoppingCartHandler struct {
	shoppingCartBusiness shoppingcarts.Business
}

func NewShoppingCartHandler(business shoppingcarts.Business) *ShoppingCartHandler {
	return &ShoppingCartHandler{
		shoppingCartBusiness: business,
	}
}

func (h *ShoppingCartHandler) GetAllHistoryOrder(c echo.Context) error {
	id := c.Param("idUser")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	idInt, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.shoppingCartBusiness.GetHistoryOrder(idInt, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	if result == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseShoppingCart.FromCoreList(result),
	})
}

func (h *ShoppingCartHandler) AddCart(c echo.Context) error {
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
	id := c.Param("idBook")
	idBookInt, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized Book ID",
		})
	}
	quantity := c.FormValue("quantity")
	quantitiyInt, _ := strconv.Atoi(quantity)
	price := c.FormValue("price")
	priceInt, _ := strconv.Atoi(price)
	var cart = _requestShoppingCart.ShoppingCart{
		TotalQuantity: uint(quantitiyInt),
		TotalPrice:    uint(priceInt),
		Status:        "Wish List",
		UserID:        idToken,
	}
	idCart, result, err := h.shoppingCartBusiness.CreateCart(idToken, idBookInt, _requestShoppingCart.ToCore(cart))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed insert to cart",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed insert to cart",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert to cart",
		"id_cart": idCart,
	})
}

func (h *ShoppingCartHandler) UpdatedStatusCart(c echo.Context) error {
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
	status := c.FormValue("status")
	result, err := h.shoppingCartBusiness.UpdatedStatusCart(idToken, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to updated status",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to updated status",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updated status",
	})
}

func (h *ShoppingCartHandler) UpdatedCart(c echo.Context) error {
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
	quantity := c.FormValue("quantity")
	quantitiyInt, _ := strconv.Atoi(quantity)
	price := c.FormValue("price")
	priceInt, _ := strconv.Atoi(price)
	result, err := h.shoppingCartBusiness.UpdatedCart(idCartInt, idToken, idBookInt, quantitiyInt, priceInt)
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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data shopping cart details",
	})
}
