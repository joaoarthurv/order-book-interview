package model

import "time"

type OrderExecutedBody struct {
	OrderBuyId         string
	OrderSellId        string
	OrderProductType   string
	OrderQuantity      int
	Price              float64
	OrderBuyOwnerId    string
	OrderSellOwnerId   string
	ExecutionTimestamp time.Time
}
