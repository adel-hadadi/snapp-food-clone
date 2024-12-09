package repository

import (
	"context"
	"time"

	"snapp-food/internal/entity"

	"github.com/jmoiron/sqlx"
)

type StoreTypeRepository struct {
	db *sqlx.DB
}

func NewStoreTypeRepository(db *sqlx.DB) StoreTypeRepository {
	return StoreTypeRepository{db: db}
}

func (r StoreTypeRepository) Get(ctx context.Context) ([]entity.StoreType, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM store_types`

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	storeTypes := make([]entity.StoreType, 0)
	for rows.Next() {
		var st entity.StoreType

		err := rows.StructScan(&st)
		if err != nil {
			return nil, err
		}

		storeTypes = append(storeTypes, st)
	}

	return storeTypes, nil
}

func (r StoreTypeRepository) Create(ctx context.Context, name, url, image string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO store_types (name, url, image) VALUES($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, name, url, image)

	return err
}
