package presentation

import (
	shoppingcartdetails "altaproject/features/shoppingCartDetails"
	_responseSCD "altaproject/features/shoppingCartDetails/presentation/response"
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
	id := c.Param("id")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idInt, errid := strconv.Atoi(id)
	if errid != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "id not recognized",
		})
	}
	result, err := h.shoppingCartDetailBusiness.GetAllCartDetails(idInt, limitint, offsetint)
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
