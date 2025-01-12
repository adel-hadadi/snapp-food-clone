package tokenservice

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func (s Service) GetClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.accessSecret), nil
	})
	if err != nil {
		return nil, errors.New("unable to parse access token")
	}
	return claims, nil
}
