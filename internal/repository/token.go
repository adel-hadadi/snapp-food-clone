package repository

import (
	"context"
	"snapp-food/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) TokenRepository {
	return TokenRepository{db: db}
}

func (r TokenRepository) Create(ctx context.Context, userID int, token string, expireAt time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// TODO: add expire at
	query := `INSERT INTO user_tokens (user_id, token) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, userID, token)

	return err
}

func (r TokenRepository) Get(ctx context.Context, userID int, refreshToken string) (entity.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM user_tokens WHERE user_id = $1 AND token = $2`

	var token entity.Token
	err := r.db.QueryRowxContext(ctx, query, userID, refreshToken).
		StructScan(&token)

	return token, err
}

func (r TokenRepository) ExistsUserRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT true FROM user_tokens WHERE token = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, refreshToken).
		Scan(&exists)

	return exists, err
}
