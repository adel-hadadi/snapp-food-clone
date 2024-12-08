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
