package validator

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator ...
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}

// DefaultValidator ...
func DefaultValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}
