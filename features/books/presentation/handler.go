package presentation

import (
	"altaproject/features/books"
	_responseBook "altaproject/features/books/presentation/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookBusiness books.Business
}

func NewBookHandler(business books.Business) *BookHandler {
	return &BookHandler{
		bookBusiness: business,
	}
}

func (h *BookHandler) GetAllBook(c echo.Context) error {
	//param, query param, binding
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.bookBusiness.GetAllBook(uint(limitint), uint(offsetint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FromCoreList(result),
	})
}
