package repository

import (
	"context"
	"snapp-food/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type ProvinceRepository struct {
	db *sqlx.DB
}

func NewProvinceRepository(db *sqlx.DB) ProvinceRepository {
	return ProvinceRepository{db: db}
}

func (r ProvinceRepository) Get(ctx context.Context) ([]entity.Province, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM provinces`

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	provinces := make([]entity.Province, 0)
	for rows.Next() {
		var province entity.Province
		if err := rows.StructScan(&province); err != nil {
			return nil, err
		}
		provinces = append(provinces, province)
	}

	return provinces, nil
}
