package otpservice

import (
	"context"

	"snapp-food/internal/entity"
)

const (
	otpLength = 4

	MsgUserNotFound    = "کاربری با این شماره موبایل یافت نشد"
	OTPMessageTemplate = `
    اسنپ فود
    کد احراز هویت شما:
    %v
    `
)

type (
	Service struct {
		sender  NotificationService
		otpRepo OTPRepository
	}

	NotificationService interface {
		Send(ctx context.Context, phone, text string) error
	}

	OTPRepository interface {
		Create(ctx context.Context, phone string, code int, prefix string) error
		Check(ctx context.Context, phone string, code int, prefix string) (entity.OTP, error)
		Used(ctx context.Context, id int) error
	}
)

func New(sender NotificationService, otpRepo OTPRepository) Service {
	return Service{
		sender:  sender,
		otpRepo: otpRepo,
	}
}
