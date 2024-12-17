package storemanagerservice

import (
	"context"

	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/apperr"
)

func (s Service) Login(ctx context.Context, phone string, code int) (tokenservice.TokenRes, error) {
	const storeManagerLoginSysMSG = "store manager service login"

	if err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone:  phone,
		Code:   code,
		Prefix: "store-manager",
	}); err != nil {
		return tokenservice.TokenRes{}, err
	}

	store, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return tokenservice.TokenRes{}, apperr.New(apperr.NotFound)
		}

		return tokenservice.TokenRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(storeManagerLoginSysMSG)
	}

	tokenRes, err := s.tokenSvc.Generate(ctx, tokenservice.GenerateTokenReq{
		Name:   store.ManagerFirstName + " " + store.ManagerLastName,
		Phone:  store.Phone,
		UserID: store.ID,
	})
	if err != nil {
		return tokenservice.TokenRes{}, err
	}

	return tokenRes, nil
}
