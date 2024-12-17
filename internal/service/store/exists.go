package storeservice

import (
	"context"

	"snapp-food/pkg/apperr"
)

func (s Service) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	const checkStoreExistsByPhoneSysMSG = "store service check exists by phone"

	exists, err := s.repo.ExistsByPhone(ctx, phone)
	if err != nil {
		return false, apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(checkStoreExistsByPhoneSysMSG)
	}

	return exists, nil
}
