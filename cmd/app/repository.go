package app

import (
	"snapp-food/internal/repository"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	OTPRepo  repository.OTPRepository
	UserRepo repository.UserRepository
}

func (a *Application) setupRepositories(db *sqlx.DB) {
	a.Repositories = Repositories{
		OTPRepo:  repository.NewOTPRepository(db),
		UserRepo: repository.NewUserRepository(db),
	}
}
