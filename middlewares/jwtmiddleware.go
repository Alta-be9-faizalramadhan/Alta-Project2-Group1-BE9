package middlewares

import (
	"altaproject/constant"
	"fmt"
	"time"

	jwtv4 "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(constant.SECRET_JWT),
	})
}

func CreateToken(userId int) (string, error) {
	claims := jwtv4.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token expired after 1 hour
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func ExtractToken(e echo.Context) (int, error) {
	user := e.Get("user").(*jwtv4.Token)
	if user.Valid {
		claims := user.Claims.(jwtv4.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId), nil
	}
	return 0, fmt.Errorf("token invalid")
}
