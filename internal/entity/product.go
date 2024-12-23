package entity

import "time"

type Product struct {
	ID                int       `json:"id" db:"id"`
	Name              string    `json:"name" db:"name"`
	Slug              string    `json:"slug" db:"slug"`
	Image             string    `json:"image" db:"image"`
	Rate              float32   `json:"rate" db:"rate"`
	StoreID           int       `json:"store_id" db:"store_id"`
	CategoryID        int       `json:"category_id" db:"category_id"`
	ProductCategoryID *int      `json:"product_category_id" db:"product_category_id"`
	Status            int8      `json:"status" db:"status"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	Price             int       `json:"price" db:"price"`

	Store Store `json:"store,omitempty"`
}
