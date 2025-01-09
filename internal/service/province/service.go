package provinceservice

import (
	"context"
	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"
	"snapp-food/pkg/convert"
)

type Service struct {
	repo provinceRepository
}

type provinceRepository interface {
	Get(ctx context.Context) ([]entity.Province, error)
}

func New(repo provinceRepository) Service {
	return Service{repo: repo}
}

type ProvinceRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s Service) List(ctx context.Context) ([]ProvinceRes, error) {
	const getListOfProvincesSysMSG = "province service get list"

	provinces, err := s.repo.Get(ctx)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(getListOfProvincesSysMSG)
	}

	provinceRes := make([]ProvinceRes, len(provinces))
	for k, province := range provinces {
		p, _ := convert.ToStruct[ProvinceRes](province)
		provinceRes[k] = p
	}

	return provinceRes, nil
}
