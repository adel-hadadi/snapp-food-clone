package entity

import "time"

type User struct {
	ID         int
	FirstName  *string   `json:"first_name" db:"first_name"`
	LastName   *string   `json:"last_name" db:"last_name"`
	Phone      string    `json:"phone" db:"phone"`
	NationalID *string   `json:"national_id" db:"national_id"`
	Status     int8      `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`

	DefaultAddressID *int `json:"default_address_id" db:"default_address_id"`
}

func (u User) FullName() string {
	if u.FirstName == nil && u.LastName == nil {
		return ""
	}
	return *u.FirstName + " " + *u.LastName
}
