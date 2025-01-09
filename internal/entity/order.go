package entity

import "time"

const (
	OrderStatusPending int8 = iota
	OrderStatusPaying
	OrderStatusStorePending
	OrderStatusSending
	OrderStatusDone

	OrderStatusPendingLabel = ""
)

type StatusLabel struct {
	Label   string `json:"status_label"`
	LabelFa string `json:"status_label_fa"`
}

var orderStatusLabels = map[int8]StatusLabel{
	OrderStatusPending: {
		Label:   "pending",
		LabelFa: "در انتظار پرداخت",
	},
	OrderStatusPaying: {
		Label:   "paying",
		LabelFa: "در حال پرداخت",
	},
	OrderStatusStorePending: {
		Label:   "store_pending",
		LabelFa: "در انتظار تایید رستوران",
	},
	OrderStatusSending: {
		Label:   "sending",
		LabelFa: "در حال ارسال",
	},
	OrderStatusDone: {
		Label:   "done",
		LabelFa: "تکمیل شده",
	},
}

type Order struct {
	ID            int       `json:"id" db:"id"`
	UserID        int       `json:"user_id" db:"user_id"`
	UserAddressID int       `json:"user_address_id" db:"user_address_id"`
	StoreID       int       `json:"store_id" db:"store_id"`
	Amount        int       `json:"amount" db:"amount"`
	Status        int8      `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`

	Store Store       `json:"store"`
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	ID        int  `json:"id" db:"id"`
	OrderID   int  `json:"order_id" db:"order_id"`
	ProductID int  `json:"product_id" db:"product_id"`
	Quantity  int8 `json:"quantity" db:"quantity"`
	Price     int  `json:"price" db:"price"`

	Product Product `json:"product"`
}

func (o Order) StatusLabel() StatusLabel {
	return orderStatusLabels[o.Status]
}
