package productcategoryservice

import (
	"context"

	"snapp-food/internal/entity"
	"snapp-food/pkg/apperr"

	"github.com/gosimple/slug"
)

type CreateReq struct {
	Name  string
	Image string
}

func (s Service) Create(ctx context.Context, req CreateReq) error {
	if err := s.repo.Create(ctx, entity.ProductCategory{
		Name:  req.Name,
		Slug:  slug.MakeLang(req.Name, "fa"),
		Image: req.Image,
	}); err != nil {
		return apperr.New(apperr.Unexpected).WithErr(err)
	}

	return nil
}
