package userservice

import (
	"context"
	"snapp-food/pkg/apperr"
)

type UpdateUserReq struct {
	FirstName  string
	LastName   string
	NationalID string
}

func (s Service) Update(ctx context.Context, userID int, req UpdateUserReq) error {
	const ErrUpdateUserSysMsg = "user service update"

	err := s.userRepo.Update(ctx, userID, req.FirstName, req.LastName, req.NationalID)
	if err != nil {
		if apperr.IsSQLDuplicateEntry(err) {
			return apperr.New(apperr.Conflict)
		}

		return apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(ErrUpdateUserSysMsg)
	}

	return nil
}
