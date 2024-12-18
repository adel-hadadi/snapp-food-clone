package repository

import (
	"context"
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

	query := `SELECT 
id, name, slug, logo, store_type_id, status, created_at,
updated_at, manager_first_name, manager_last_name, phone, st_astext(location) 
FROM stores WHERE id=$1`

	var store entity.Store
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&store.ID,
		&store.Name,
		&store.Slug,
		&store.Logo,
		&store.StoreTypeID,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.ManagerFirstName,
		&store.ManagerLastName,
		&store.Phone,
		&store.Location,
	)

	return store, err
}

func (r StoreRepository) FindBySlug(ctx context.Context, slug string) (entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT 
id, name, slug, logo, store_type_id, status, created_at,
updated_at, manager_first_name, manager_last_name, phone, st_astext(location) 
FROM stores WHERE slug=$1`

	var store entity.Store
	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&store.ID,
		&store.Name,
		&store.Slug,
		&store.Logo,
		&store.StoreTypeID,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.ManagerFirstName,
		&store.ManagerLastName,
		&store.Phone,
		&store.Location,
	)

	store.GenLatAndLong()

	return store, err
}

func (r StoreRepository) FindByPhone(ctx context.Context, phone string) (entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT 
id, name, slug, logo, store_type_id, status, created_at,
updated_at, manager_first_name, manager_last_name, phone, st_astext(location) 
FROM stores WHERE phone=$1`

	var store entity.Store
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&store.ID,
		&store.Name,
		&store.Slug,
		&store.Logo,
		&store.StoreTypeID,
		&store.Status,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.ManagerFirstName,
		&store.ManagerLastName,
		&store.Phone,
		&store.Location,
	)

	return store, err
}

func (r StoreRepository) Get(ctx context.Context) ([]entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT 
id, name, slug, logo, store_type_id, status, created_at,
updated_at, manager_first_name, manager_last_name, phone, st_astext(location) 
FROM stores`

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
			&store.Logo,
			&store.StoreTypeID,
			&store.Status,
			&store.CreatedAt,
			&store.UpdatedAt,
			&store.ManagerFirstName,
			&store.ManagerLastName,
			&store.Phone,
			&store.Location,
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

	query := `INSERT INTO stores 
		(name, slug, location, logo, manager_first_name, manager_last_name, phone, store_type_id, city_id) 
		VALUES 
		($1, $2, st_makepoint($3, $4), $5, $6, $7, $8, $9, $10)`

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
		store.CityID,
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

func (r StoreRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM stores WHERE phone=$1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, phone).Scan(&exists)

	return exists, err
}

func (r StoreRepository) Nearest(ctx context.Context, userID int) ([]entity.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
	SELECT stores.id, stores.name, stores.slug, stores.logo FROM stores
	RIGHT JOIN user_addresses ON user_addresses.id = (select default_address_id from users where users.id=$1)
	WHERE stores.city_id = user_addresses.city_id
	ORDER BY stores.location <-> user_addresses.location
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stores := make([]entity.Store, 0)

	for rows.Next() {
		var store entity.Store

		if err := rows.Scan(
			&store.ID,
			&store.Name,
			&store.Slug,
			&store.Logo,
		); err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	return stores, nil
}
