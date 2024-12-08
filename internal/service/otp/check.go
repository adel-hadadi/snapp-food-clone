package otpservice

import (
	"context"
	"snapp-food/pkg/apperr"
)

type OTPCheckReq struct {
	Phone string
	Code  int
}

const OTPIsInvalid = "کد وارد شده اشتباه می‌باشد"

func (s Service) Check(ctx context.Context, req OTPCheckReq) error {
	const checkOTPSysMsg = "otp service check otp code"

	otp, err := s.otpRepo.Check(ctx, req.Phone, req.Code)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return apperr.New(apperr.Invalid).WithMsg(OTPIsInvalid)
		}

		return apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(checkOTPSysMsg)
	}

	const updateOTPStatus = "otp service update otp status to used"
	if err := s.otpRepo.Used(ctx, otp.ID); err != nil {
		return apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(updateOTPStatus)
	}

	return nil
}
