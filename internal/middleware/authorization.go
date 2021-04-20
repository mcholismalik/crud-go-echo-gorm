package middleware

import (
	"fmt"
	"os"
	"strings"

	res "crud-go-echo-gorm/pkg/util/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.RespError(c, &res.ErrUnauthorized)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.RespError(c, &res.ErrUnauthorized)
		}
		c.Set("username", token.Claims.(jwt.MapClaims)["username"])

		return next(c)
	}
}
