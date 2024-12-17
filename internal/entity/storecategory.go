package entity

type StoreCategory struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	StoreID  int       `json:"store_id" db:"store_id"`
	Products []Product `json:"products"`
}
