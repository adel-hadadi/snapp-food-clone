package entity

type StoreType struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	URL   string `json:"url" db:"url"`
	Image string `json:"image" db:"image"`
}
