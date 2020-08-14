package repository

import (
	"go-boilerplate/domain/authentication"
	"go-boilerplate/models"
)

type repoHandler struct {
}

// NewVoucherRepository ....
func NewAuthenticationRepository() authentication.Repository {
	return &repoHandler{}
}

func (r *repoHandler) Authenticate(createSess models.CreateSessionRequest) bool {
	if createSess.Password != "123" {
		return false
	}

	return true
}
