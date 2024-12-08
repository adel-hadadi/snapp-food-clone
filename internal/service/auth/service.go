package authservice

import (
	"context"
	"snapp-food/internal/entity"
	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
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
	Generate(ctx context.Context, user entity.User) (tokenservice.TokenRes, error)
}

func New(otpSvc otpService, userRepo userRepository, tokenSvc tokenService) Service {
	return Service{
		otpSvc:   otpSvc,
		userRepo: userRepo,
		tokenSvc: tokenSvc,
	}
}
