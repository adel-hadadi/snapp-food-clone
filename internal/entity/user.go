package entity

import "time"

type User struct {
	ID         int
	Name       *string   `json:"name" db:"name"`
	Phone      string    `json:"phone" db:"phone"`
	NationalID *string   `json:"national_id" db:"national_id"`
	Status     int8      `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
