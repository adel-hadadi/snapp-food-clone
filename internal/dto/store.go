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
