package entity

import "time"

type OTP struct {
	ID        int       `json:"id" db:"id"`
	Phone     string    `json:"phone" db:"phone"`
	Code      int       `json:"code" db:"code"`
	Prefix    string    `json:"prefix" db:"prefix"`
	Status    int8      `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

const (
	StatusUsed Status = iota
	StatusUnUsed
)
