package middleware

import (
	"go-boilerplate/models"
	"go-boilerplate/shared/context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := c.(*context.ApplicationContext)
		authorizationHeader := c.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return ct.ResponseJSON(400, nil, "Token is missing")
		}

		token := strings.Split(authorizationHeader, " ")
		if len(token) != 2 {
			return ct.ResponseJSON(400, nil, "Token is missing")
		}

		res := verifyToken(token[1])

		if res == http.StatusUnauthorized {
			return ct.ResponseJSON(400, nil, "Unauthorized")
		}

		if res == http.StatusBadRequest {
			return ct.ResponseJSON(400, nil, "Token error")
		}

		return next(c)
	}
}

func verifyToken(tokenString string) int {
	jwtKey := []byte(viper.GetString("jwt.secretKey"))
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized
		}
		return http.StatusBadRequest
	}
	if !tkn.Valid {
		return http.StatusUnauthorized
	}

	return 200
}
