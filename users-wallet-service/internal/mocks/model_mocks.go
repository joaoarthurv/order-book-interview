package mocks

import (
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/model"
	"time"
)

var OrderExecutedListenerMock = model.OrderExecutedBody{
	OrderBuyId:         "a2b6782c-8f43-4c4e-99f4-9e3b70f85737",
	OrderSellId:        "75f57348-d508-49d6-a6a0-ad644153287d",
	OrderSellOwnerId:   "10",
	Price:              150,
	OrderQuantity:      5,
	OrderProductType:   "VIBRANIUM",
	OrderBuyOwnerId:    "2",
	ExecutionTimestamp: time.Date(2023, 03, 16, 0, 40, 8, 37878000, time.UTC),
}

var UserWalletMock = model.UserWallet{
	UserId:          "3",
	ProductType:     "VIBRANIUM",
	Balance:         3127,
	ProductQuantity: 1963,
	CreatedAt:       time.Date(2023, 03, 16, 0, 40, 8, 37878000, time.UTC),
	UpdatedAt:       time.Date(2023, 03, 16, 0, 42, 8, 37878000, time.UTC),
}
