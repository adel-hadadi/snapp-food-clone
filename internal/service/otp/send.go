package otpservice

import (
	"context"
	"fmt"

	"snapp-food/pkg/apperr"
	"snapp-food/pkg/random"
)

type OTPSendReq struct {
	Phone  string
	Prefix string
}

func (s Service) Send(ctx context.Context, req OTPSendReq) error {
	const sendOTPCodeSysMSG = "otp service send otp code"

	otpCode := random.Num(otpLength)

	if err := s.sender.Send(
		ctx,
		req.Phone,
		fmt.Sprintf(OTPMessageTemplate, otpCode),
	); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(sendOTPCodeSysMSG)
	}

	const saveOTPCodeSysMSG = "otp service save otp code"
	if err := s.otpRepo.Create(ctx, req.Phone, fmt.Sprintf("%v", otpCode), req.Prefix); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(saveOTPCodeSysMSG)
	}

	return nil
}
