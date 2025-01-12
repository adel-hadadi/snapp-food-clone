package tokenservice

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func (s Service) ValidateAccessToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.accessSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired access token")
	}
	return claims, nil
}
