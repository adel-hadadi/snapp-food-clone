package storemanagerservice

import (
	"context"

	"snapp-food/internal/entity"
	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
)

type Service struct {
	otpSvc   otpService
	tokenSvc tokenService
	repo     storeRepository
}

type otpService interface {
	Check(ctx context.Context, req otpservice.OTPCheckReq) error
}

type tokenService interface {
	Generate(ctx context.Context, user tokenservice.GenerateTokenReq) (tokenservice.TokenRes, error)
}

type storeRepository interface {
	FindByPhone(ctx context.Context, phone string) (entity.Store, error)
}

func New(otpSvc otpService, tokenSvc tokenService, repo storeRepository) Service {
	return Service{
		otpSvc:   otpSvc,
		tokenSvc: tokenSvc,
		repo:     repo,
	}
}
