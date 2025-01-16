package repository

import (
	"context"
	"snapp-food/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE phone=$1)`

	var exists bool
	err := r.db.QueryRowxContext(ctx, query, phone).Scan(&exists)

	return exists, err
}

func (r UserRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (first_name, last_name, phone, national_id) VALUES($1, $2, $3, $4) RETURNING(id)`

	var createdID int
	err := r.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Phone, user.NationalID).Scan(&createdID)
	if err != nil {
		return entity.User{}, err
	}

	user.ID = createdID

	return user, nil
}

func (r UserRepository) GetByPhone(ctx context.Context, phone string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE phone=$1`

	var user entity.User

	if err := r.db.
		QueryRowxContext(ctx, query, phone).
		StructScan(&user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r UserRepository) Get(ctx context.Context, userID int) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id=$1`

	var user entity.User
	err := r.db.QueryRowxContext(ctx, query, userID).StructScan(&user)

	return user, err
}

func (r UserRepository) Update(ctx context.Context, userID int, firstName, lastName, nationalID string, defaultAddress int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE users
    SET first_name=$1, last_name=$2, national_id=$3, default_address_id=$4
    WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query, firstName, lastName, nationalID, defaultAddress, userID)

	return err
}
