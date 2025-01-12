package authservice

import (
	"context"

	"snapp-food/internal/entity"
	otpservice "snapp-food/internal/service/otp"
)

type Service struct {
	otpSvc   otpService
	userRepo userRepository
	tokenSvc tokenService
}

type otpService interface {
	Check(ctx context.Context, req otpservice.OTPCheckReq) error
}

type userRepository interface {
	GetByPhone(ctx context.Context, phone string) (entity.User, error)
	Create(ctx context.Context, phone string) (entity.User, error)
}

type tokenService interface {
	GenerateTokens(ctx context.Context, userID int) (string, string, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
}

func New(otpSvc otpService, userRepo userRepository, tokenSvc tokenService) Service {
	return Service{
		otpSvc:   otpSvc,
		userRepo: userRepo,
		tokenSvc: tokenSvc,
	}
}
