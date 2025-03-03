package repository

import (
	"context"
	"log"
	"time"

	"snapp-food/internal/entity"

	"github.com/jmoiron/sqlx"
)

type OTPRepository struct {
	db *sqlx.DB
}

func NewOTPRepository(db *sqlx.DB) OTPRepository {
	return OTPRepository{db: db}
}

func (r OTPRepository) Create(ctx context.Context, phone, code, prefix string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO otps (phone, code, prefix) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, phone, code, prefix)

	return err
}

func (r OTPRepository) Check(ctx context.Context, phone, code, prefix string) (entity.OTP, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var otp entity.OTP
	query := `SELECT * FROM otps
        WHERE phone=$1
        AND code=$2
        AND status=$3
          AND prefix=$4
        AND created_at >= NOW() - INTERVAL '2 minutes'`

	log.Println("check status", entity.StatusUnUsed)
	err := r.db.QueryRowxContext(
		ctx,
		query,
		phone,
		code,
		entity.StatusUnUsed,
		prefix,
	).StructScan(&otp)

	return otp, err
}

func (r OTPRepository) Used(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE otps SET status=$1 WHERE id=$2`
	_, err := r.db.ExecContext(ctx, query, entity.StatusUsed, id)

	return err
}
