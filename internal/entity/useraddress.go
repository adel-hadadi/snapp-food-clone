package entity

import (
	"regexp"
	"strconv"
	"strings"
)

type UserAddress struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Latitude   float64 `json:"latitude" db:"latitude"`
	Longitude  float64 `json:"longitude" db:"longitude"`
	Location   string  `json:"location" db:"location"`
	UserID     int     `json:"user_id" db:"user_id"`
	CityID     int     `json:"city_id" db:"city_id"`
	ProvinceID int     `json:"province_id" db:"province_id"`
	Address    string  `json:"address" db:"address"`
	Detail     string  `json:"detail" db:"detail"`

	Province Province `json:"province"`
	City     City     `json:"city"`
}

func (a *UserAddress) GenLatAndLong() {
	rgx := regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(a.Location)

	points := strings.Split(rs[1], " ")

	lat, _ := strconv.ParseFloat(points[0], 64)
	long, _ := strconv.ParseFloat(points[1], 64)

	a.Latitude = lat

	a.Longitude = long
}
