package presentation

import (
	"altaproject/features/books"
	_requestBook "altaproject/features/books/presentation/request"
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

func (h *BookHandler) PostNewBook(c echo.Context) error {
	title := c.FormValue("title")
	author := c.FormValue("author")
	publisher := c.FormValue("publisher")
	isbn := c.FormValue("isbn")
	category := c.FormValue("category")
	description := c.FormValue("description")
	price := c.FormValue("price")
	priceInt, _ := strconv.Atoi(price)
	stock := c.FormValue("stock")
	stockInt, _ := strconv.Atoi(stock)
	bookPage := c.FormValue("book_page")
	userId := c.FormValue("user_id")
	userIdInt, _ := strconv.Atoi(userId)

	var newBook = _requestBook.Book{
		Title:       title,
		Author:      author,
		Publisher:   publisher,
		ISBN:        isbn,
		Category:    category,
		Description: description,
		Price:       uint(priceInt),
		Stock:       uint(stockInt),
		Sold:        0,
		BookPage:    bookPage,
		UserId:      userIdInt,
	}

	// errBind := c.Bind(&newBook)
	// if errBind != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "failed to bind data, check your input",
	// 	})
	// }

	dataUser := _requestBook.ToCore(newBook)
	row, err := h.bookBusiness.CreateNewBook(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert book",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to insert book",
	})
}
