package repository

import (
	"context"
	"snapp-food/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type CityRepository struct {
	db *sqlx.DB
}

func NewCityRepository(db *sqlx.DB) CityRepository {
	return CityRepository{db: db}
}

func (r CityRepository) GetByProvinceID(ctx context.Context, provinceID int) ([]entity.City, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM cities WHERE province_id = $1`

	rows, err := r.db.QueryxContext(ctx, query, provinceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := make([]entity.City, 0)
	for rows.Next() {
		var city entity.City
		if err := rows.StructScan(&city); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}
