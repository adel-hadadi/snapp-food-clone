package dto

type StoreCreateReq struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Logo        string  `json:"logo"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	StoreTypeID int     `json:"store_type_id"`
	CityID      int     `json:"city_id"`
	ManagerID   int     `json:"manager_id"`
}

type StoreRes struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Logo        string  `json:"logo"`
	Slug        string  `json:"slug"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	StoreTypeID int     `json:"store_type_id"`
	CityID      int     `json:"city_id"`
	ManagerID   int     `json:"manager_id"`
}
