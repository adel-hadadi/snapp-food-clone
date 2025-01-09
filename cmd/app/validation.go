package app

import (
	"snapp-food/internal/adapters/validation"

	"github.com/rezakhademix/govalidator/v2"
)

type Validations struct {
	Auth validation.AuthValidation
}

func (a *Application) setupValidations(v govalidator.Validator) Validations {
	return Validations{
		Auth: validation.NewAuthValidation(v),
	}
}
