package model

import (
	"time"
)

type OrdersDataRow struct {
	OrderId          string    `json:"id,omitempty" db:"id"`
	OrderType        string    `json:"orderType,omitempty" db:"order_type"`
	OrderPrice       float64   `json:"price,omitempty" db:"price"`
	OrderQuantity    int       `json:"quantity,omitempty" db:"quantity"`
	OrderProductType string    `json:"productType,omitempty" db:"product_type"`
	OrderOwnerId     string    `json:"ownerId,omitempty" db:"owner_id"`
	OrderTimestamp   time.Time `json:"timestamp,omitempty" db:"order_timestamp"`
}
