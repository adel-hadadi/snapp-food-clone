package orderservice

import (
	"context"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
)

const (
	orderNotFoundMSG = "سفارش یافت نشد"
)

func (s Service) Pay(ctx context.Context, userID, orderID int) error {
	const getOrderSysMSG = "order service pay method get order"

	order, err := s.repo.GetByID(ctx, orderID)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return apperr.New(apperr.NotFound).WithMsg(orderNotFoundMSG)
		}
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getOrderSysMSG)
	}

	if order.UserID != userID {
		return apperr.New(apperr.Invalid).WithMsg(orderNotFoundMSG)
	}

	if order.Status != entity.OrderStatusPending {
		return apperr.New(apperr.Invalid).WithMsg(orderNotFoundMSG)
	}

	const updateOrderStatusSysMSG = "order service pay method update order status"

	if err := s.repo.UpdateStatus(
		ctx,
		orderID,
		entity.OrderStatusStorePending,
	); err != nil {
		return apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(updateOrderStatusSysMSG)
	}

	return nil
}
