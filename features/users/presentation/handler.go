package presentation

import (
	"altaproject/features/users"
	_requestUser "altaproject/features/users/presentation/request"
	_responseUser "altaproject/features/users/presentation/response"
	"altaproject/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
	var newUser _requestUser.User
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
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
	var user _requestUser.User
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data",
		})
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
