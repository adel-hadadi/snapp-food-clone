package useraddressservice

import (
	"context"

	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
	"snapp-food/pkg/convert"
)

type UserAddressRes struct {
	ID        int
	Name      string
	Latitude  float64
	Longitude float64
	Address   string
	Detail    string
	Province  string
	City      string
}

func (s Service) Get(ctx context.Context, userID int) ([]UserAddressRes, error) {
	const getUserAddressesSysMsg = "user address service get by user id"

	addresses, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg(getUserAddressesSysMsg)
	}

	addressesRes := make([]UserAddressRes, 0, len(addresses))
	for a := range addresses {
		addressesRes = append(addressesRes, UserAddressRes{
			ID:        addresses[a].ID,
			Name:      addresses[a].Name,
			Latitude:  addresses[a].Latitude,
			Longitude: addresses[a].Longitude,
			Address:   addresses[a].Address,
			Detail:    addresses[a].Detail,
			Province:  addresses[a].Province.Name,
			City:      addresses[a].City.Name,
		})
	}

	return addressesRes, nil
}

type CreateReq struct {
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Detail     string  `json:"detail"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	ProvinceID int     `json:"province_id"`
	CityID     int     `json:"city_id"`
}

func (s Service) Create(ctx context.Context, userID int, req CreateReq) error {
	const createUserAddressSysMsg = "user address service create"

	userAddress, _ := convert.ToStruct[entity.UserAddress](req)

	if err := s.repo.Create(ctx, userID, userAddress); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(createUserAddressSysMsg)
	}

	return nil
}
