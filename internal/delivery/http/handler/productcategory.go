package handler

import (
	"context"
	"net/http"

	productcategoryservice "snapp-food/internal/service/productcategory"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
)

type ProductCategoryHandler struct {
	productCategorySvc productCategoryService
}

type productCategoryService interface {
	Create(ctx context.Context, req productcategoryservice.CreateReq) error
	List(ctx context.Context) ([]productcategoryservice.ProductCategoryRes, error)
}

func NewProductCategoryHandler(service productCategoryService) ProductCategoryHandler {
	return ProductCategoryHandler{productCategorySvc: service}
}

type CreateProductCategoryReq struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (h ProductCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := httpreq.Bind[CreateProductCategoryReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	if err := h.productCategorySvc.Create(r.Context(), productcategoryservice.CreateReq{
		Name:  req.Name,
		Image: req.Image,
	}); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}

func (h ProductCategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	productCategories, err := h.productCategorySvc.List(r.Context())
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, productCategories, http.StatusOK)
}
