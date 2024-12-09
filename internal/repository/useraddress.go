package repository

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"snapp-food/internal/entity"
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

	query := `SELECT user_addresses.*, c.id, c.name, p.id, p.name FROM user_addresses 
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
			&address.Latitude,
			&address.Longitude,
			&address.UserID,
			&address.CityID,
			&address.ProvinceID,
			&address.Address,
			&address.Detail,
			&address.City.ID,
			&address.City.Name,
			&address.Province.ID,
			&address.Province.Name,
		); err != nil {
			return nil, err
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (r UserAddressRepository) Create(ctx context.Context, userID int, address entity.UserAddress) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO user_addresses 
		(name, latitude, longitude, user_id, city_id, province_id, address, detail) 
		VALUES ($1,  $2, $3, $4, $5, $6, $7, $8)`

	log.Printf("user address %+v", address)

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
