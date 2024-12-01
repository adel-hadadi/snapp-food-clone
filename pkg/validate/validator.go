package validate

import (
	"fmt"

	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
)

type Validator struct {
	uni      *ut.UniversalTranslator
	validate *validator.Validate
}

func New() Validator {
	validate := validator.New()

	fa := fa.New()

	uni := ut.New(fa, fa)

	trans, _ := uni.GetTranslator("fa")
	fa_translations.RegisterDefaultTranslations(validate, trans)

	return Validator{
		uni:      uni,
		validate: validate,
	}
}

var tagToMessage = map[string]string{
	"required": "%s اجباری است",
}

var fieldTrans = map[string]string{
	"Name":  "نام",
	"Phone": "شماره تلفن",
}

func (v Validator) Struct(s any) map[string]string {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errs := make(map[string]string)
		for _, vErr := range validationErrs {
			// TODO: Should use json tag

			field := vErr.Field()
			f := fieldTrans[field]
			msg := tagToMessage[vErr.Tag()]

			errs[field] = fmt.Sprintf(msg, f)
		}

		return errs
	}

	return nil
}
