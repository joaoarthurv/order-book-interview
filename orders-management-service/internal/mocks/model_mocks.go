package mocks

import (
	"gitlab.com/projetos/orderbook/orders-management-service/internal/model"
	"time"
)

var SelectSellOrdersMock = []model.OrdersDataRow{
	{
		OrderId:          "f409d80d-52e3-445d-9832-088a10f27e2e",
		OrderType:        "SELL",
		OrderPrice:       100,
		OrderQuantity:    10,
		OrderProductType: "VIBRANIUM",
		OrderOwnerId:     "5",
		OrderTimestamp:   time.Date(2023, 03, 16, 0, 39, 8, 37878000, time.UTC),
	},
	{
		OrderId:          "a2b6782c-8f43-4c4e-99f4-9e3b70f85737",
		OrderType:        "SELL",
		OrderPrice:       150,
		OrderQuantity:    5,
		OrderProductType: "VIBRANIUM",
		OrderOwnerId:     "2",
		OrderTimestamp:   time.Date(2023, 03, 16, 0, 40, 8, 37878000, time.UTC),
	},
	{
		OrderId:          "c1d3457e-1a2b-4f6d-b9e8-6c5a9d03c1b0",
		OrderType:        "SELL",
		OrderPrice:       300,
		OrderQuantity:    20,
		OrderProductType: "VIBRANIUM",
		OrderOwnerId:     "7",
		OrderTimestamp:   time.Date(2023, 03, 16, 0, 55, 8, 37878000, time.UTC),
	},
}

var SelectBuyOrdersMock = []model.OrdersDataRow{
	{
		OrderId:          "f409d80d-52e3-445d-9832-088a10f27e2e",
		OrderType:        "BUY",
		OrderPrice:       400,
		OrderQuantity:    10,
		OrderProductType: "VIBRANIUM",
		OrderOwnerId:     "5",
		OrderTimestamp:   time.Date(2023, 03, 16, 0, 39, 8, 37878000, time.UTC),
	},
	{
		OrderId:          "a2b6782c-8f43-4c4e-99f4-9e3b70f85737",
		OrderType:        "BUY",
		OrderPrice:       350,
		OrderQuantity:    5,
		OrderProductType: "VIBRANIUM",
		OrderOwnerId:     "2",
		OrderTimestamp:   time.Date(2023, 03, 16, 0, 40, 8, 37878000, time.UTC),
	},
	{
		OrderId:          "c1d3457e-1a2b-4f6d-b9e8-6c5a9d03c1b0",
		OrderType:        "BUY",
		OrderPrice:       100,
		OrderQuantity:    20,
		OrderProductType: "VIBRANIUM",
		OrderOwnerId:     "7",
		OrderTimestamp:   time.Date(2023, 03, 16, 0, 55, 8, 37878000, time.UTC),
	},
}

var BuyOrderRequestMock = model.OrdersDataRow{
	OrderPrice:       500,
	OrderQuantity:    300,
	OrderProductType: "VIBRANIUM",
	OrderOwnerId:     "65",
}

var BuyOrderLessRequestMock = model.OrdersDataRow{
	OrderPrice:       500,
	OrderQuantity:    18,
	OrderProductType: "VIBRANIUM",
	OrderOwnerId:     "65",
}

var SellOrderRequestMock = model.OrdersDataRow{
	OrderPrice:       10,
	OrderQuantity:    50,
	OrderProductType: "VIBRANIUM",
	OrderOwnerId:     "65",
}

var SellOrderLessRequestMock = model.OrdersDataRow{
	OrderPrice:       10,
	OrderQuantity:    30,
	OrderProductType: "VIBRANIUM",
	OrderOwnerId:     "65",
}
