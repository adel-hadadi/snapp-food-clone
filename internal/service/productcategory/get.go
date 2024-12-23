package productcategoryservice

import (
	"context"
	"snapp-food/pkg/apperr"
)

type ProductCategoryRes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Slug  string `json:"slug"`
}

func (s Service) List(ctx context.Context) ([]ProductCategoryRes, error) {
	categories, err := s.repo.Get(ctx)
	if err != nil {
		return nil, apperr.New(apperr.Unexpected).WithErr(err)
	}

	categoryRes := make([]ProductCategoryRes, 0, len(categories))
	for _, category := range categories {
		categoryRes = append(categoryRes, ProductCategoryRes{
			ID:    category.ID,
			Name:  category.Name,
			Image: category.Image,
			Slug:  category.Slug,
		})
	}

	return categoryRes, nil
}
