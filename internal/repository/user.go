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

func (r UserRepository) CheckExistsByPhoneNumber(ctx context.Context, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE phone=$1)`

	var exists bool
	err := r.db.QueryRowxContext(ctx, query, phone).Scan(&exists)

	return exists, err
}

func (r UserRepository) Create(ctx context.Context, phone string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (phone) VALUES($1)`

	_, err := r.db.ExecContext(ctx, query, phone)
	if err != nil {
		return entity.User{}, nil
	}

	return entity.User{Phone: phone}, nil
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

func (r UserRepository) Update(ctx context.Context, userID int, firstName, lastName, nationalID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE users
    SET first_name=$1, last_name=$2, national_id=$3
    WHERE id = $4`

	_, err := r.db.ExecContext(ctx, query, firstName, lastName, nationalID, userID)

	return err
}
