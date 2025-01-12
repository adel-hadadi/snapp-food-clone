package handler

import (
	"context"
	"log"
	"net/http"
	orderservice "snapp-food/internal/service/order"
	"snapp-food/pkg/httpres"
	"snapp-food/pkg/server/httpreq"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	orderSvc orderService
}

type orderService interface {
	Create(ctx context.Context, userID int, storeSlug string, items []orderservice.OrderItemReq) error
	List(ctx context.Context, userID int) ([]orderservice.OrderRes, error)
	Pay(ctx context.Context, userID, orderID int) error
}

func NewOrderHandler(cartSvc orderService) OrderHandler {
	return OrderHandler{orderSvc: cartSvc}
}

type CreateOrderReq struct {
	// TODO: it can be store id
	StoreSlug string `json:"store_slug"`
	Items     []struct {
		ProductID int  `json:"product_id"`
		Quantity  int8 `json:"quantity"`
	} `json:"items"`
}

func (h OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := httpreq.AuthID(r)

	req, err := httpreq.Bind[CreateOrderReq](r)
	if err != nil {
		httpres.ValidationErr(w, err, http.StatusBadRequest)
		return
	}

	orderItems := make([]orderservice.OrderItemReq, len(req.Items))
	for i, item := range req.Items {
		orderItems[i] = orderservice.OrderItemReq{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	if err := h.orderSvc.Create(r.Context(), userID, req.StoreSlug, orderItems); err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, nil, http.StatusCreated)
}

func (h OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := httpreq.AuthID(r)
	log.Println("order list => ", userID)

	orders, err := h.orderSvc.List(r.Context(), userID)
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	httpres.Success(w, orders, http.StatusOK)
}

func (h OrderHandler) Pay(w http.ResponseWriter, r *http.Request) {
	userID := httpreq.AuthID(r)

	orderID, err := strconv.Atoi(chi.URLParam(r, "orderID"))
	if err != nil {
		httpres.WithErr(w, err)
		return
	}

	h.orderSvc.Pay(r.Context(), userID, orderID)
}
