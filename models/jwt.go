package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email    string
	Username string
	Role     string

	jwt.StandardClaims
}
