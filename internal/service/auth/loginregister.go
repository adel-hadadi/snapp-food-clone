package authservice

import (
	"context"
	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/apperr"
)

type LoginRegisterReq struct {
	Phone string
	Code  int
}

func (s Service) LoginRegister(ctx context.Context, req LoginRegisterReq) (tokenservice.TokenRes, error) {
	err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone: req.Phone,
		Code:  req.Code,
	})
	if err != nil {
		return tokenservice.TokenRes{}, err
	}

	const getUserSysMsg = "auth service get user"
	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			const createUserSysMsg = "auth service create user"
			user, err = s.userRepo.Create(ctx, req.Phone)
			if err != nil {
				return tokenservice.TokenRes{}, apperr.New(apperr.Unexpected).
					WithErr(err).
					WithSysMsg(createUserSysMsg)
			}
		} else {
			return tokenservice.TokenRes{}, apperr.New(apperr.Unexpected).
				WithErr(err).
				WithSysMsg(getUserSysMsg)
		}
	}

	token, err := s.tokenSvc.Generate(ctx, user)
	if err != nil {
		return tokenservice.TokenRes{}, err
	}

	return token, nil
}
