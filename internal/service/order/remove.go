package orderservice

import (
	"context"
	"snapp-food/pkg/apperr"
)

func (s Service) RemovePending(ctx context.Context) error {
	const removePendingOrdersSysMsg = "order service remove pending method"

	err := s.repo.RemovePending(ctx)
	if err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(removePendingOrdersSysMsg)
	}

	return nil
}
