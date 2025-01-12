package tokenservice

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s Service) GenerateTokens(ctx context.Context, userID int) (string, string, error) {
	accessClaims := Claims{
		UserID: userID,
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
		UserID: userID,
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

	if err := s.repo.Create(
		ctx,
		userID,
		refreshHash,
		refreshClaims.ExpiresAt.Time,
	); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s Service) GenerateAccessToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
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
