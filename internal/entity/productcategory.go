package entity

type ProductCategory struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Slug  string `json:"slug" db:"slug"`
	Image string `json:"image" db:"image"`
}
