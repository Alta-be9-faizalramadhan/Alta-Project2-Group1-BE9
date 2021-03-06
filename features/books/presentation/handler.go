package presentation

import (
	"altaproject/features/books"
	_requestBook "altaproject/features/books/presentation/request"
	_responseBook "altaproject/features/books/presentation/response"
	"altaproject/middlewares"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
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

	var storageClient *storage.Client
	bucket := "bucket-project-2"
	ctx := appengine.NewContext(c.Request())
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "misssing credentials file",
		})
	}
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	if file.Size > 1024*1024 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "The uploaded image is too big. Please use an image less than 1MB in size",
		})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
		if file.Filename[len(file.Filename)-4:] != "jpeg" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "The provided file format is not allowed. Please upload a JPG or JPEG or PNG image",
			})
		}
	}

	sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return err
	}

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
		UserId:      idToken,
		ImageURL:    u.String(),
	}

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

func (h *BookHandler) GetBookById(c echo.Context) error {
	id := c.Param("id")
	idBook, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.bookBusiness.GetBookById(idBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FromCore(result),
	})
}

func (h *BookHandler) UpdatedBook(c echo.Context) error {
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

	id := c.Param("id")
	idBook, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}

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
	sold := c.FormValue("sold")
	soldInt, _ := strconv.Atoi(sold)

	var storageClient *storage.Client
	bucket := "bucket-project-2"
	ctx := appengine.NewContext(c.Request())
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "misssing credentials file",
		})
	}
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	if file.Size > 1024*1024 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "The uploaded image is too big. Please use an image less than 1MB in size",
		})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
		if file.Filename[len(file.Filename)-4:] != "jpeg" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "The provided file format is not allowed. Please upload a JPG or JPEG or PNG image",
			})
		}
	}

	sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return err
	}

	var book = _requestBook.Book{
		Title:       title,
		Author:      author,
		Publisher:   publisher,
		ISBN:        isbn,
		Category:    category,
		Description: description,
		Price:       uint(priceInt),
		Stock:       uint(stockInt),
		Sold:        uint(soldInt),
		BookPage:    bookPage,
		UserId:      idToken,
		ImageURL:    u.String(),
	}

	result, err := h.bookBusiness.UpdatedBook(idBook, _requestBook.ToCore(book))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
	})
}

func (h *BookHandler) DeleteBookById(c echo.Context) error {
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
	id := c.Param("id")
	idBook, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.bookBusiness.SoftDeleteBook(idBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete data",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}

func (h *BookHandler) GetBookByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	result, err := h.bookBusiness.SelectBookByCategory(category)
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

func (h *BookHandler) GetBookByUserId(c echo.Context) error {
	idToken, _ := middlewares.ExtractToken(c)
	user_id := idToken
	result, err := h.bookBusiness.SelectBookByUserId(user_id)
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
