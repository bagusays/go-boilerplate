package service

import (
	"go-boilerplate/domain/authentication"
	"go-boilerplate/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type serviceHandler struct {
	authenticationRepo authentication.Repository
}

func NewAuthenticationService(authenticationRepo authentication.Repository) authentication.Service {
	return &serviceHandler{
		authenticationRepo: authenticationRepo,
	}
}

func (s serviceHandler) GenerateToken(createSess models.CreateSessionRequest) (interface{}, error) {
	res := s.authenticationRepo.Authenticate(createSess)

	if !res {
		return false, nil
	}

	jwtKey := []byte(viper.GetString("jwt.secretKey"))
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := models.Claims{
		Username: "admin",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
