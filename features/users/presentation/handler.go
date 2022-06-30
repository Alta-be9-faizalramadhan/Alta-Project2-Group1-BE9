package presentation

import (
	"altaproject/features/users"
	_requestUser "altaproject/features/users/presentation/request"
	_responseUser "altaproject/features/users/presentation/response"
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

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.userBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FromCoreList(result),
	})
}

func (h *UserHandler) AddUser(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	alamat := c.FormValue("alamat")
	notelp := c.FormValue("notelp")

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
	var newUser = _requestUser.User{
		UserName: username,
		Email:    email,
		Password: password,
		Alamat:   alamat,
		NoTelp:   notelp,
		ImageURL: u.String(),
	}

	dataUser := _requestUser.ToCore(newUser)
	row, err := h.userBusiness.InsertData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to insert data",
	})
}

func (h *UserHandler) PutData(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	alamat := c.FormValue("alamat")
	notelp := c.FormValue("notelp")

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
	var user = _requestUser.User{
		UserName: username,
		Email:    email,
		Password: password,
		Alamat:   alamat,
		NoTelp:   notelp,
		ImageURL: u.String(),
	}

	result, err := h.userBusiness.UpdateData(idUser, _requestUser.ToCore(user))
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
	if result == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
	})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}
	id := c.Param("id")
	idnya, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id not recognize",
		})
	}
	if idToken != idnya {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.userBusiness.GetDataUser(idnya)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get data user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"data":    _responseUser.FromCore(result),
	})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	idTok, errDel := middlewares.ExtractToken(c)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid",
		})
	}
	id := c.Param("id")
	idDel, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id not recognize",
		})
	}
	if idTok != idDel {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	_, err := h.userBusiness.DeleteData(idDel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to delete user",
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	//var authData _requestUser.User
	email := c.FormValue("email")
	password := c.FormValue("password")
	// err := c.Bind(&authData)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "failed to bind data, check your input",
	// 	})
	// }
	//dataUser := _requestUser.ToCore(authData)
	token, username, id, e := h.userBusiness.Login(email, password)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "email or password incorrect",
		})
	}
	data := map[string]interface{}{
		"token":     token,
		"user_name": username,
		"user_id":   id,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login succes",
		"data":    data,
	})
}
