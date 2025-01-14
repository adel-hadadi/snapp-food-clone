package authservice

import (
	"context"

	otpservice "snapp-food/internal/service/otp"
	userservice "snapp-food/internal/service/user"
	"snapp-food/pkg/apperr"
	"snapp-food/pkg/convert"
)

type LoginRegisterReq struct {
	Phone string
	Code  int
}

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRegisterRes struct {
	User  userservice.UserRes `json:"user"`
	Token TokenRes            `json:"token"`
}

func (s Service) LoginRegister(ctx context.Context, req LoginRegisterReq) (LoginRegisterRes, error) {
	err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone:  req.Phone,
		Code:   req.Code,
		Prefix: "user",
	})
	if err != nil {
		return LoginRegisterRes{}, err
	}

	const getUserSysMsg = "auth service get user"
	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			const createUserSysMsg = "auth service create user"
			user, err = s.userRepo.Create(ctx, req.Phone)
			if err != nil {
				return LoginRegisterRes{}, apperr.New(apperr.Unexpected).
					WithErr(err).
					WithSysMsg(createUserSysMsg)
			}
		} else {
			return LoginRegisterRes{}, apperr.New(apperr.Unexpected).
				WithErr(err).
				WithSysMsg(getUserSysMsg)
		}
	}

	act, rft, err := s.tokenSvc.GenerateTokens(ctx, user.ID)
	if err != nil {
		return LoginRegisterRes{}, err
	}

	userRes, _ := convert.ToStruct[userservice.UserRes](user)

	return LoginRegisterRes{
		User:  userRes,
		Token: TokenRes{AccessToken: act, RefreshToken: rft},
	}, nil
}
