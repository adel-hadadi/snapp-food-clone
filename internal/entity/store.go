package entity

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Store struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Slug        string    `json:"slug" db:"slug"`
	Rate        float32   `json:"rate" db:"rate"`
	Latitude    float64   `json:"latitude" db:"latitude"`
	Longitude   float64   `json:"longitude" db:"longitude"`
	Location    string    `json:"location" db:"location"`
	CityID      *int      `json:"city_id" db:"city_id"`
	Address     string    `json:"address" db:"address"`
	Logo        string    `json:"logo" db:"logo"`
	StoreTypeID int       `json:"store_type_id" db:"store_type_id"`
	ManagerID   int       `json:"manager_id" db:"manager_id"`
	Status      int       `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	StoreType StoreType `json:"store_type"`
}

func (s *Store) GenLatAndLong() {
	rgx := regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(s.Location)

	if len(rs) == 0 {
		return
	}

	points := strings.Split(rs[1], " ")

	lat, _ := strconv.ParseFloat(points[0], 64)
	long, _ := strconv.ParseFloat(points[1], 64)

	s.Latitude = lat

	s.Longitude = long
}
