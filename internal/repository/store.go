package repository

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"snapp-food/internal/entity"
)

type StoreRepository struct {
	db *sqlx.DB
}

func NewStoreRepository(db *sqlx.DB) StoreRepository {
	return StoreRepository{db: db}
}

func (r StoreRepository) Find(ctx context.Context, id int) (entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM stores WHERE id=$1`

	var store entity.Store
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&store.ID,
		&store.Name,
		&store.Slug,
		&store.Latitude,
		&store.Longitude,
		&store.Logo,
		&store.StoreTypeID,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.ManagerFirstName,
		&store.ManagerLastName,
		&store.Phone,
	)

	return store, err
}

func (r StoreRepository) FindBySlug(ctx context.Context, slug string) (entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM stores WHERE slug=$1`

	var store entity.Store
	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&store.ID,
		&store.Name,
		&store.Slug,
		&store.Latitude,
		&store.Longitude,
		&store.Logo,
		&store.StoreTypeID,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.ManagerFirstName,
		&store.ManagerLastName,
		&store.Phone,
	)

	return store, err
}

func (r StoreRepository) Get(ctx context.Context) ([]entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM stores`

	stores := make([]entity.Store, 0)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var store entity.Store

		err := rows.Scan(
			&store.ID,
			&store.Name,
			&store.Slug,
			&store.Latitude,
			&store.Longitude,
			&store.Logo,
			&store.StoreTypeID,
			&store.Status,
			&store.CreatedAt,
			&store.UpdatedAt,
			&store.ManagerFirstName,
			&store.ManagerLastName,
			&store.Phone,
		)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	return stores, nil
}

func (r StoreRepository) Create(ctx context.Context, store entity.Store) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	log.Println("store slug", store.Slug)

	query := `INSERT INTO stores 
		(name, slug, latitude, longitude, logo, manager_first_name, manager_last_name, phone, store_type_id) 
		VALUES 
		($1, $2, $3, $4, $5,$6, $7, $8, $9)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		store.Name,
		store.Slug,
		store.Latitude,
		store.Longitude,
		store.Logo,
		store.ManagerFirstName,
		store.ManagerLastName,
		store.Phone,
		store.StoreTypeID,
	)

	return err
}

func (r StoreRepository) Update(ctx context.Context, id int, store entity.Store) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE stores SET name=$1, latitude=$2, longitude=$3, logo=$4, manager_first_name=$5, manager_last_name=$6, phone=$7, store_type_id=$8, status=$9 WHERE id=$10`

	_, err := r.db.ExecContext(
		ctx,
		query,
		store.Name,
		store.Latitude,
		store.Longitude,
		store.Logo,
		store.ManagerFirstName,
		store.ManagerLastName,
		store.Phone,
		store.StoreTypeID,
		store.Status,
		id,
	)

	return err
}
