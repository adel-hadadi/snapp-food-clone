package storetypeservice

import (
	"context"

	"snapp-food/pkg/apperr"
)

type StoreType struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

func (s Service) Get(ctx context.Context) ([]StoreType, error) {
	const getStoreTypesSysMsg = "store type service get list"

	types, err := s.repo.Get(ctx)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getStoreTypesSysMsg)
	}

	storeTypesRes := make([]StoreType, 0, len(types))
	for t := range types {
		storeTypesRes = append(storeTypesRes, StoreType{
			ID:    types[t].ID,
			Name:  types[t].Name,
			Image: types[t].Image,
			URL:   types[t].URL,
		})
	}

	return storeTypesRes, nil
}
