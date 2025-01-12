package authservice

import (
	"context"

	otpservice "snapp-food/internal/service/otp"
	"snapp-food/pkg/apperr"
)

type LoginRegisterReq struct {
	Phone string
	Code  int
}

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) LoginRegister(ctx context.Context, req LoginRegisterReq) (TokenRes, error) {
	err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone:  req.Phone,
		Code:   req.Code,
		Prefix: "user",
	})
	if err != nil {
		return TokenRes{}, err
	}

	const getUserSysMsg = "auth service get user"
	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			const createUserSysMsg = "auth service create user"
			user, err = s.userRepo.Create(ctx, req.Phone)
			if err != nil {
				return TokenRes{}, apperr.New(apperr.Unexpected).
					WithErr(err).
					WithSysMsg(createUserSysMsg)
			}
		} else {
			return TokenRes{}, apperr.New(apperr.Unexpected).
				WithErr(err).
				WithSysMsg(getUserSysMsg)
		}
	}

	act, rft, err := s.tokenSvc.GenerateTokens(ctx, user.ID)
	if err != nil {
		return TokenRes{}, err
	}

	return TokenRes{AccessToken: act, RefreshToken: rft}, nil
}
