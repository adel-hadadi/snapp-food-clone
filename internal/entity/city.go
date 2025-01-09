package entity

type City struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProvinceID int    `json:"province_id" db:"province_id"`
}
