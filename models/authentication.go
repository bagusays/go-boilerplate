package models

type CreateSessionRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
}
