package cityservice

import (
	"context"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
	"snapp-food/pkg/convert"
)

type Service struct {
	repo cityRepository
}

type cityRepository interface {
	GetByProvinceID(ctx context.Context, provinceID int) ([]entity.City, error)
}

func New(repo cityRepository) Service {
	return Service{repo: repo}
}

type CityRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s Service) ListByProvinceID(ctx context.Context, provinceID int) ([]CityRes, error) {
	cities, err := s.repo.GetByProvinceID(ctx, provinceID)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).
			WithErr(err).
			WithSysMsg("city service get list by province id")
	}

	cityRes := make([]CityRes, len(cities))
	for k, city := range cities {
		c, _ := convert.ToStruct[CityRes](city)
		cityRes[k] = c
	}

	return cityRes, nil
}
