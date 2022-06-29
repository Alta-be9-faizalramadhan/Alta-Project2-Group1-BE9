package presentation

import (
	category "altaproject/features/categorys"
	_responseCategory "altaproject/features/categorys/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	catagoryBusiness category.Business
}

func NewCategoryHandler(business category.Business) *CategoryHandler {
	return &CategoryHandler{
		catagoryBusiness: business,
	}
}

func (h *CategoryHandler) GetAllCategory(c echo.Context) error {
	result, err := h.catagoryBusiness.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all category",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseCategory.FromCoreList(result),
	})
}
