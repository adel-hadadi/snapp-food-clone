package entity

type City struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
