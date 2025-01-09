package productserviec

import (
	"context"

	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"

	"github.com/gosimple/slug"
)

type CreateProductReq struct {
	Name       string
	Image      string
	CategoryID int
	Price      float64
}

func (s Service) Create(ctx context.Context, storeID int, req CreateProductReq) error {
	const createProductSysMSG = "product service create"

	if err := s.repo.Create(ctx, entity.Product{
		Name:       req.Name,
		Slug:       slug.MakeLang(req.Name, "fa"),
		Image:      req.Image,
		StoreID:    storeID,
		CategoryID: &req.CategoryID,
		Price:      int(req.Price * 10),
	}); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err).WithSysMsg(createProductSysMSG)
	}

	return nil
}
