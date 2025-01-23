package repository

import (
	"context"
	"time"

	"snapp-food/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserAddressRepository struct {
	db *sqlx.DB
}

func NewUserAddressRepository(db *sqlx.DB) UserAddressRepository {
	return UserAddressRepository{db: db}
}

func (r UserAddressRepository) GetByUserID(ctx context.Context, userID int) ([]entity.UserAddress, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
	SELECT 
		user_addresses.id,
		user_addresses.name,
		user_addresses.user_id,
		user_addresses.city_id,
		user_addresses.province_id,
		user_addresses.address,
		user_addresses.detail,
		st_astext(user_addresses.location),
		c.id, c.name, p.id, p.name 
	FROM user_addresses 
	LEFT JOIN cities c on c.id = user_addresses.city_id 
	LEFT OUTER JOIN provinces p on p.id = user_addresses.province_id 
	WHERE user_id = $1`

	rows, err := r.db.QueryxContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := make([]entity.UserAddress, 0)
	for rows.Next() {
		var address entity.UserAddress

		if err := rows.Scan(
			&address.ID,
			&address.Name,
			&address.UserID,
			&address.CityID,
			&address.ProvinceID,
			&address.Address,
			&address.Detail,
			&address.Location,
			&address.City.ID,
			&address.City.Name,
			&address.Province.ID,
			&address.Province.Name,
		); err != nil {
			return nil, err
		}

		address.GenLatAndLong()

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (r UserAddressRepository) Create(ctx context.Context, userID int, address entity.UserAddress) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO user_addresses 
		(name, location, user_id, city_id, province_id, address, detail) 
		VALUES ($1, st_makepoint($2, $3), $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		address.Name,
		address.Latitude,
		address.Longitude,
		userID,
		address.CityID,
		address.ProvinceID,
		address.Address,
		address.Detail,
	)

	return err
}

func (r UserAddressRepository) BelongsToUser(ctx context.Context, addressID, userID int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT EXISTS(SELECT 1 FROM user_addresses WHERE user_id=$1 AND id=$2)`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, userID, addressID).Scan(&exists)

	return exists, err
}

func (r UserAddressRepository) GetUserDefaultAddress(ctx context.Context, userID int) (entity.UserAddress, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT id, name, address FROM user_addresses WHERE user_addresses.id = (
		SELECT default_address_id FROM users WHERE users.id = $1
	)`

	var address entity.UserAddress
	if err := r.db.QueryRowContext(ctx, query, userID).
		Scan(
			&address.ID,
			&address.Name,
			&address.Address,
		); err != nil {
		return entity.UserAddress{}, err
	}

	return address, nil
}
