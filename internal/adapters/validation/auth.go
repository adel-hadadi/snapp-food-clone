package validation

import (
	"snapp-food/internal/dto"

	"github.com/rezakhademix/govalidator/v2"
)

const (
	codeSize           = 4
	phoneNumberLenSize = 10
)

type AuthValidation struct {
	v govalidator.Validator
}

func NewAuthValidation(v govalidator.Validator) AuthValidation {
	return AuthValidation{
		v: v,
	}
}

func (v AuthValidation) ValidateLoginRegister(req dto.AuthLoginRegisterReq) (map[string]string, bool) {
	ok := v.v.RequiredString(req.Phone, "phone", "").
		LenString(req.Phone, phoneNumberLenSize, "phone", "").
		RequiredInt(req.Code, "phone", "").
		LenInt(req.Code, codeSize, "code", "").
		IsPassed()

	return v.v.Errors(), ok
}
