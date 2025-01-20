package tokenservice

import (
	"context"
	"snapp-food/pkg/apperr"

	"github.com/golang-jwt/jwt/v5"
)

func (s Service) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.refreshSecret), nil
	})
	if err != nil || !token.Valid {
		return "", apperr.New(apperr.Unauthorized)
	}

	refreshHash := hashToken(refreshToken)

	const retriveTokenFromDatabaseSysMSG = "token service retrive token from database"
	_, err = s.repo.Get(ctx, claims.User.ID, refreshHash)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return "", apperr.New(apperr.Unauthorized)
		}

		return "", apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(retriveTokenFromDatabaseSysMSG)
	}

	// TODO: check expire at
	// if time.Now().After(expiresAt) {
	// 	return "", errors.New("refresh token expired")
	// }

	return s.GenerateAccessToken(claims.User)
}
