package userservice

import (
	"context"
	"snapp-food/pkg/apperr"
)

type UpdateUserReq struct {
	FirstName        string
	LastName         string
	NationalID       string
	DefaultAddressID int
}

const (
	msgAddressDoesNotBelongsToUser = "آدرس وارد شده متعلق به کاربر وارد شده نمی‌باشد"
)

func (s Service) Update(ctx context.Context, userID int, req UpdateUserReq) error {
	const checkAddressBelongsToUserSysMSG = "user service update method check address belongs to user"

	ok, err := s.userAddressRepo.BelongsToUser(ctx, req.DefaultAddressID, userID)
	if err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(checkAddressBelongsToUserSysMSG)
	}

	if !ok {
		return apperr.New(apperr.Invalid).WithMsg(msgAddressDoesNotBelongsToUser)
	}

	const ErrUpdateUserSysMsg = "user service update"

	if err := s.userRepo.Update(
		ctx,
		userID,
		req.FirstName,
		req.LastName,
		req.NationalID,
		req.DefaultAddressID,
	); err != nil {
		if apperr.IsSQLDuplicateEntry(err) {
			return apperr.New(apperr.Conflict)
		}

		return apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(ErrUpdateUserSysMsg)
	}

	return nil
}
