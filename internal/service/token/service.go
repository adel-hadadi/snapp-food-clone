package tokenservice

import (
	"context"
	"snapp-food/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpireTime  = 900   // 15m
	RefreshTokenExpireTime = 86400 // 24h

	UserID = "user_id"
	Name   = "name"
	Exp    = "exp"
	Phone  = "phone"
	Status = "status"
)

type Service struct {
	accessSecret  string
	refreshSecret string
	repo          tokenRepository
}

type tokenRepository interface {
	Create(ctx context.Context, userID int, token string, expireAt time.Time) error
	Get(ctx context.Context, userID int, refreshToken string) (entity.Token, error)
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func New(repo tokenRepository, accessSecret, refreshSecret string) Service {
	return Service{
		repo:          repo,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}
