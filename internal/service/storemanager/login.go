package storemanagerservice

import (
	"context"

	authservice "snapp-food/internal/service/auth"
	otpservice "snapp-food/internal/service/otp"
	"snapp-food/pkg/apperr"
)

// TODO: centralize every login
func (s Service) Login(ctx context.Context, phone string, code int) (authservice.TokenRes, error) {
	const storeManagerLoginSysMSG = "store manager service login"

	if err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone:  phone,
		Code:   code,
		Prefix: "store-manager",
	}); err != nil {
		return authservice.TokenRes{}, err
	}

	store, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return authservice.TokenRes{}, apperr.New(apperr.NotFound)
		}

		return authservice.TokenRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(storeManagerLoginSysMSG)
	}

	// TODO: fix this mother fucker
	accessToken, refreshToken, err := s.tokenSvc.GenerateTokens(ctx, store.ID)
	if err != nil {
		return authservice.TokenRes{}, err
	}

	return authservice.TokenRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
