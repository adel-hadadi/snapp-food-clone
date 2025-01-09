package orderservice

import (
	"context"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
)

type OrderItemReq struct {
	ProductID int
	Quantity  int8
}

func (s Service) Create(ctx context.Context, userID int, storeSlug string, items []OrderItemReq) error {
	const getStoreSysMSG = "order service get store by slug"

	store, err := s.storeRepo.FindBySlug(ctx, storeSlug)
	if err != nil {
		if apperr.IsSQLNoRows(err) {
			return apperr.New(apperr.NotFound)
		}

		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getStoreSysMSG)
	}

	const getUserDefaultAddressSysMSG = "order service get user default address"

	userAddress, err := s.userAddressRepo.GetUserDefaultAddress(ctx, userID)
	if err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getUserDefaultAddressSysMSG)
	}

	var totalAmount int

	const checkProductExistsSysMSG = "order service check products exists"
	orderItems := make([]entity.OrderItem, len(items))

	for i, item := range items {
		product, err := s.productRepo.Find(ctx, item.ProductID)
		if err != nil {
			if apperr.IsSQLNoRows(err) {
				return apperr.New(apperr.NotFound)
			}

			return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(checkProductExistsSysMSG)
		}

		totalAmount += product.Price * int(item.Quantity)

		orderItems[i] = entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
	}

	const createOrderSysMSG = "order service create order"
	if err := s.repo.Create(
		ctx,
		userID,
		userAddress.ID,
		store.ID,
		totalAmount,
		orderItems,
	); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(createOrderSysMSG)
	}

	return nil
}
