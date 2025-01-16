package sellerstorehandler

import storeservice "snapp-food/internal/service/store"

type Handler struct {
	storeSvc storeservice.Service
}

func New(storeSvc storeservice.Service) Handler {
	return Handler{storeSvc: storeSvc}
}
