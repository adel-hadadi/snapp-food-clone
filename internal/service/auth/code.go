package authservice

import (
	"context"
	"snapp-food/internal/dto"
	otpservice "snapp-food/internal/service/otp"
)

func (s Service) SendCode(ctx context.Context, phone string) error {
	return s.otpSvc.Send(ctx, otpservice.OTPSendReq{Phone: phone, Prefix: "user"})
}

func (s Service) SellerSendCode(ctx context.Context, req dto.AuthSendSellerOTPReq) error {
	return s.otpSvc.Send(ctx, otpservice.OTPSendReq{Phone: req.Phone, Prefix: "seller"})
}
