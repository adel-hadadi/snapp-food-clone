package tokenservice

import (
	"context"
	"crypto/sha256"
	"fmt"
	"snapp-food/pkg/apperr"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	Role        string `json:"role"`
}

func (s Service) GenerateTokens(ctx context.Context, user User) (string, string, error) {
	accessClaims := Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(AccessTokenExpireTime * time.Second),
			),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(s.accessSecret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(RefreshTokenExpireTime * time.Second),
			),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(s.refreshSecret))
	if err != nil {
		return "", "", err
	}

	refreshHash := hashToken(refreshToken)

	const saveRefreshTokenSysMsg = "token service generate method"
	if err := s.repo.Create(
		ctx,
		user.ID,
		refreshHash,
		refreshClaims.ExpiresAt.Time,
	); err != nil {
		return "", "", apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(saveRefreshTokenSysMsg)
	}

	return accessToken, refreshToken, nil
}

func (s Service) GenerateAccessToken(user User) (string, error) {
	claims := Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(AccessTokenExpireTime * time.Second),
			),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.accessSecret))
}

func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return fmt.Sprintf("%x", h.Sum(nil))
}
