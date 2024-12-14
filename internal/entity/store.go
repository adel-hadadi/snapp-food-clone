package entity

import "time"

type Store struct {
	ID               int       `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Slug             string    `json:"slug" db:"slug"`
	ManagerFirstName string    `json:"manager_first_name" db:"manager_first_name"`
	ManagerLastName  string    `json:"manager_last_name" db:"manager_last_name"`
	Phone            string    `json:"phone" db:"phone"`
	Latitude         float64   `json:"latitude" db:"latitude"`
	Longitude        float64   `json:"longitude" db:"longitude"`
	Address          string    `json:"address" db:"address"`
	Logo             string    `json:"logo" db:"logo"`
	StoreTypeID      int       `json:"store_type_id" db:"store_type_id"`
	Status           int       `json:"status" db:"status"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
