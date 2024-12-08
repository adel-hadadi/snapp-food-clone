package userservice

import (
	"context"
	"snapp-food/pkg/apperr"
)

type UserRes struct {
	ID         int
	FirstName  *string
	LastName   *string
	Phone      string
	NationalID *string
}

func (s Service) Get(ctx context.Context, userID int) (UserRes, error) {
	const getUserSysMsg = "user service get user"

	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return UserRes{}, apperr.New(apperr.NotFound)
		}

		return UserRes{}, apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(getUserSysMsg)
	}

	return UserRes{
		ID:         userID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Phone:      user.Phone,
		NationalID: user.NationalID,
	}, nil
}
