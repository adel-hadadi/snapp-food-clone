package adapters

import (
	"context"
	"log"
)

type NotificationSMS struct {
}

func NewNotificationSMS() NotificationSMS {
	return NotificationSMS{}
}

func (s NotificationSMS) Send(ctx context.Context, phone, text string) error {
	// TODO: complete this section
	log.Printf("send %s sms to %s\n", text, phone)

	return nil
}
