package tokenservice

import (
	"os"
	"snapp-food/pkg/apperr"

	"github.com/golang-jwt/jwt/v5"
)

const (
	ErrTokenIsInvalid = "token is invalid"
)

func (s Service) Validate(tokenString string) (jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, apperr.New(apperr.Unexpected)
			}

			return []byte(os.Getenv("SECRET")), nil
		})

	if err != nil {
		return jwt.Token{}, apperr.New(apperr.Invalid).
			WithMsg(ErrTokenIsInvalid)
	}

	return *token, nil
}
