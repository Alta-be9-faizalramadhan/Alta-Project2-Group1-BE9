package routes

import (
	"altaproject/factory"
	"altaproject/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("users", presenter.UserPresenter.GetAll, middlewares.JWTMiddleware())
	e.GET("users/:id", presenter.UserPresenter.GetUser, middlewares.JWTMiddleware())
	e.POST("users", presenter.UserPresenter.AddUser)
	e.PUT("users/:id", presenter.UserPresenter.PutData, middlewares.JWTMiddleware())
	e.DELETE("users/:id", presenter.UserPresenter.DeleteUser, middlewares.JWTMiddleware())
	e.POST("login", presenter.UserPresenter.Login)

	e.GET("books", presenter.BookPresenter.GetAllBook)
	e.GET("books/:id", presenter.BookPresenter.GetBookById)
	e.POST("books", presenter.BookPresenter.PostNewBook, middlewares.JWTMiddleware())

	return e
}
