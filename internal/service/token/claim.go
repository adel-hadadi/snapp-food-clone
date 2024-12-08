package tokenservice

import (
	"context"
	"snapp-food/pkg/apperr"

	"github.com/golang-jwt/jwt/v5"
)

func (s Service) Claim(ctx context.Context, tokenString string) (map[string]any, error) {
	const tokenClaimSysMsg = "token service claim token"

	t, err := s.Validate(tokenString)
	if err != nil {
		return nil, err
	}

	claims := make(map[string]any)
	tokenClaims, ok := t.Claims.(jwt.MapClaims)

	if ok && t.Valid {
		for k, v := range tokenClaims {
			claims[k] = v
		}

		return claims, nil
	}

	return nil, apperr.New(apperr.Unexpected).WithErr(err).
		WithSysMsg(tokenClaimSysMsg)
}
