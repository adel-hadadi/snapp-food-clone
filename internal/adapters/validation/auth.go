package validation

import (
	"fmt"
	"snapp-food/internal/dto"

	"github.com/rezakhademix/govalidator/v2"
)

const (
	codeLenSize        = 4
	phoneNumberLenSize = 10

	phoneIsRequiredMsg = "شماره تلفن اجبرای میباشد"
	phoneNumberLenMsg  = "شماره تلفن باید %d رقم باشد"
	codeIsRequiredMsg  = "کد تایید الزامی میباشد"
	codeLenMsg         = "کد تایید باید %v رقم باشد"
)

type AuthValidation struct {
	validator govalidator.Validator
}

func NewAuthValidation(validator govalidator.Validator) AuthValidation {
	return AuthValidation{
		validator: validator,
	}
}

func (v AuthValidation) ValidateLoginRegister(req dto.AuthLoginRegisterReq) (map[string]string, bool) {
	ok := v.validator.RequiredString(req.Phone, "phone", phoneIsRequiredMsg).
		LenString(req.Phone, phoneNumberLenSize, "phone", fmt.Sprintf(phoneNumberLenMsg, phoneNumberLenSize)).
		RequiredInt(req.Code, "code", codeIsRequiredMsg).
		LenInt(req.Code, codeLenSize, "code", fmt.Sprintf(codeLenMsg, codeLenSize)).
		IsPassed()

	return v.validator.Errors(), ok
}
