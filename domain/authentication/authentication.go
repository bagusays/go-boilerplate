package authentication

import "go-boilerplate/models"

type Service interface {
	GenerateToken(models.CreateSessionRequest) (interface{}, error)
}

type Repository interface {
	Authenticate(models.CreateSessionRequest) bool
}
