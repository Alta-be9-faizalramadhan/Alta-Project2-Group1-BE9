package presentation

import (
	"altaproject/features/users"
	_requestUser "altaproject/features/users/presentation/request"
	_responseUser "altaproject/features/users/presentation/response"
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
	//param, query param, binding
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
	//idToken := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	// if idToken != idUser {
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
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
	// idToken, errToken := ExtractToken(c)
	// if errToken != nil {
	// 	c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "invalid token",
	// 	})
	// }
	// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
	// 		"message": "unauthorized",
	// 	})
	// }
	id := c.Param("id")
	idnya, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id not recognize",
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
