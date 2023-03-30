package middleware

import (
	result_dto "hallo-corona/dto/result"
	jwtToken "hallo-corona/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type Result struct {
	Status  int
	Data    interface{}
	Message string
}

// Create Auth function here ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, result_dto.ErrorResult{Code: http.StatusUnauthorized, Message: "unauthorized", Error: "unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, result_dto.ErrorResult{Code: http.StatusUnauthorized, Message: "unathorized", Error: err.Error()})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
