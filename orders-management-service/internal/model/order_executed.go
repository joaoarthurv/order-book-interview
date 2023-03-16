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

func (o *OrderExecutedBody) BuildOrderExecutedBody(orderBuy *OrdersDataRow, orderSell *OrdersDataRow, orderQuantity int) *OrderExecutedBody {
	return &OrderExecutedBody{
		OrderBuyId:         orderBuy.OrderId,
		OrderSellId:        orderSell.OrderId,
		OrderProductType:   orderSell.OrderProductType,
		OrderQuantity:      orderQuantity,
		Price:              orderSell.OrderPrice,
		OrderBuyOwnerId:    orderBuy.OrderOwnerId,
		OrderSellOwnerId:   orderSell.OrderOwnerId,
		ExecutionTimestamp: time.Now(),
	}
}
