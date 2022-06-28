package presentation

import (
	shoppingcarts "altaproject/features/shoppingCarts"
	_requestShoppingCart "altaproject/features/shoppingCarts/presentation/request"
	_responseShoppingCart "altaproject/features/shoppingCarts/presentation/response"
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
	id := c.Param("id")
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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseShoppingCart.FromCoreList(result),
	})
}

func (h *ShoppingCartHandler) AddCart(c echo.Context) error {
	var cart _requestShoppingCart.ShoppingCart
	errBind := c.Bind(&cart)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to bind data",
		})
	}
	result, err := h.shoppingCartBusiness.CreateNewCart(_requestShoppingCart.ToCore(cart))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to insert book",
	})
}
