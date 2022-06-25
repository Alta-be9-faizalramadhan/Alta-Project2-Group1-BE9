package routes

import (
	"altaproject/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.GET("users", presenter.UserPresenter.GetAll)
	e.GET("users/:id", presenter.UserPresenter.GetUser)
	e.POST("users", presenter.UserPresenter.AddUser)

	return e
}
