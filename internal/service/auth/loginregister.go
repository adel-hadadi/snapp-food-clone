package authservice

import (
	"context"

	"snapp-food/internal/dto"
	"snapp-food/internal/entity"
	otpservice "snapp-food/internal/service/otp"
	tokenservice "snapp-food/internal/service/token"
	"snapp-food/pkg/apperr"
)

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) LoginRegister(ctx context.Context, req dto.AuthLoginRegisterReq) (dto.LoginRegisterRes, error) {
	exists, err := s.userRepo.ExistsByPhone(ctx, req.Phone)
	if err != nil {
		return dto.LoginRegisterRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg("check-user-exists")
	}

	if err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone:  req.Phone,
		Code:   req.Code,
		Prefix: "user",
	}); err != nil {
		return dto.LoginRegisterRes{}, err
	}

	const getUserSysMsg = "auth service get user"
	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			const createUserSysMsg = "auth service create user"
			user, err = s.userRepo.Create(ctx, entity.User{Phone: req.Phone})
			if err != nil {
				return dto.LoginRegisterRes{}, apperr.New(apperr.Unexpected).
					WithErr(err).
					WithSysMsg(createUserSysMsg)
			}
		} else {
			return dto.LoginRegisterRes{}, apperr.New(apperr.Unexpected).
				WithErr(err).
				WithSysMsg(getUserSysMsg)
		}
	}

	act, rft, err := s.tokenSvc.GenerateTokens(ctx, tokenservice.User{
		ID:          user.ID,
		DisplayName: user.FullName(),
		Role:        "customer",
	})
	if err != nil {
		return dto.LoginRegisterRes{}, err
	}

	return dto.LoginRegisterRes{
		HasAccount: exists,
		Token:      dto.TokenRes{AccessToken: act, RefreshToken: rft},
	}, nil
}

func (s Service) SellerLoginRegister(ctx context.Context, req dto.AuthLoginRegisterReq) (dto.LoginRegisterRes, error) {
	exists, err := s.userRepo.ExistsByPhone(ctx, req.Phone)
	if err != nil {
		return dto.LoginRegisterRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg("check-user-exists")
	}

	if err := s.otpSvc.Check(ctx, otpservice.OTPCheckReq{
		Phone:  req.Phone,
		Code:   req.Code,
		Prefix: "seller",
	}); err != nil {
		return dto.LoginRegisterRes{}, err
	}

	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if apperr.IsSQLNoRows(err) {
		return dto.LoginRegisterRes{}, apperr.New(apperr.Invalid)
	} else if err != nil {
		return dto.LoginRegisterRes{}, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg("get-user-by-phone")
	}

	act, rft, err := s.tokenSvc.GenerateTokens(ctx, tokenservice.User{
		ID:          user.ID,
		DisplayName: user.FullName(),
		Role:        "seller",
	})

	return dto.LoginRegisterRes{
		HasAccount: exists,
		Token: dto.TokenRes{
			AccessToken:  act,
			RefreshToken: rft,
		},
	}, nil
}
