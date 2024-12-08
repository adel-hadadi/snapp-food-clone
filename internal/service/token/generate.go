package tokenservice

import (
	"context"
	"os"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenRes struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpireTime  int64
	RefreshTokenExpireTime int64
}

func (s Service) Generate(ctx context.Context, user entity.User) (TokenRes, error) {
	var t TokenRes

	t.AccessTokenExpireTime = time.Now().Add(AccessTokenExpireTime * time.Second).Unix()
	t.RefreshTokenExpireTime = time.Now().Add(RefreshTokenExpireTime * time.Second).Unix()

	// TODO: fix jwt claims
	atc := jwt.MapClaims{
		UserID:                   user.ID,
		Name:                     user.FirstName,
		NationalID:               user.NationalID,
		Phone:                    user.Phone,
		Status:                   user.Status,
		"accessTokenExpireTime":  t.AccessTokenExpireTime,
		"refreshTokenExpireTime": t.RefreshTokenExpireTime,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error

	const GenerateTokenSysMsg = "jwt service generate jwt token"
	t.AccessToken, err = at.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return TokenRes{}, apperr.New(apperr.Unexpected).WithErr(err).
			WithSysMsg(GenerateTokenSysMsg)
	}

	rtc := jwt.MapClaims{
		UserID:                   user.ID,
		"refreshTokenExpireTime": t.RefreshTokenExpireTime,
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	t.RefreshToken, err = rt.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return TokenRes{}, apperr.New(apperr.Unexpected).WithErr(err).
			WithSysMsg(GenerateTokenSysMsg)
	}

	return t, nil
}
