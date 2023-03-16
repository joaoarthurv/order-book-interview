package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/mocks"
	"testing"
)

func Test_OrdersService(t *testing.T) {
	t.Run("[Buy Order]Should the operation be successful, any excess orders will be added to the database.", func(t *testing.T) {
		//give
		ordersDaoMock := mocks.NewIOrdersDao(t)
		topicMock := mocks.NewITopicService(t)
		userWalletClientMock := mocks.NewIUsersWalletClient(t)
		orderService := NewOrdersService(ordersDaoMock, topicMock, userWalletClientMock)

		userWalletClientMock.EXPECT().UserWalletCheck(mock.Anything, mock.Anything).Call.Return(true, nil)
		ordersDaoMock.EXPECT().SelectSellOrdersToExecute(mock.Anything, mock.Anything).Call.Return(mocks.SelectSellOrdersMock, nil)
		topicMock.EXPECT().SendEvents(mock.Anything).Call.Return(nil)
		ordersDaoMock.EXPECT().DeleteOrderById(mock.Anything, mock.Anything).Call.Return(nil).Return(nil)
		ordersDaoMock.EXPECT().InsertOrder(mock.Anything, mock.Anything).Call.Return(nil)

		//when
		err := orderService.PlaceBuyOrder(context.Background(), &mocks.BuyOrderRequestMock)

		//then
		assert.Nil(t, err)
	})

	t.Run("[Buy Order] Should the requested quantity in a purchase request for Vibraniums be lower than the quantity being sold, "+
		"the system will subtract the requested quantity and return the remaining quantity to the database.", func(t *testing.T) {
		//give
		ordersDaoMock := mocks.NewIOrdersDao(t)
		topicMock := mocks.NewITopicService(t)
		userWalletClientMock := mocks.NewIUsersWalletClient(t)
		orderService := NewOrdersService(ordersDaoMock, topicMock, userWalletClientMock)

		userWalletClientMock.EXPECT().UserWalletCheck(mock.Anything, mock.Anything).Call.Return(true, nil)
		ordersDaoMock.EXPECT().SelectSellOrdersToExecute(mock.Anything, mock.Anything).Call.Return(mocks.SelectSellOrdersMock, nil)
		topicMock.EXPECT().SendEvents(mock.Anything).Call.Return(nil)
		ordersDaoMock.EXPECT().DeleteOrderById(mock.Anything, mock.Anything).Call.Return(nil).Return(nil)
		ordersDaoMock.EXPECT().UpdateOrderQuantityById(mock.Anything, mock.Anything, mock.Anything).Call.Return(nil)

		//when
		err := orderService.PlaceBuyOrder(context.Background(), &mocks.BuyOrderLessRequestMock)

		//then
		assert.Nil(t, err)
	})

	t.Run("[Sell Order]Should the operation be successful, any excess orders will be added to the database.", func(t *testing.T) {
		//give
		ordersDaoMock := mocks.NewIOrdersDao(t)
		topicMock := mocks.NewITopicService(t)
		userWalletClientMock := mocks.NewIUsersWalletClient(t)
		orderService := NewOrdersService(ordersDaoMock, topicMock, userWalletClientMock)

		userWalletClientMock.EXPECT().UserWalletCheck(mock.Anything, mock.Anything).Call.Return(true, nil)
		ordersDaoMock.EXPECT().SelectBuyOrdersToExecute(mock.Anything, mock.Anything).Call.Return(mocks.SelectBuyOrdersMock, nil)
		topicMock.EXPECT().SendEvents(mock.Anything).Call.Return(nil)
		ordersDaoMock.EXPECT().DeleteOrderById(mock.Anything, mock.Anything).Call.Return(nil).Return(nil)
		ordersDaoMock.EXPECT().InsertOrder(mock.Anything, mock.Anything).Call.Return(nil)

		//when
		err := orderService.PlaceSellOrder(context.Background(), &mocks.SellOrderRequestMock)

		//then
		assert.Nil(t, err)
	})

	t.Run("[Sell Order] Should the requested quantity in a purchase request for Vibraniums be lower than the quantity being sold, "+
		"the system will subtract the requested quantity and return the remaining quantity to the database.", func(t *testing.T) {
		//give
		ordersDaoMock := mocks.NewIOrdersDao(t)
		topicMock := mocks.NewITopicService(t)
		userWalletClientMock := mocks.NewIUsersWalletClient(t)
		orderService := NewOrdersService(ordersDaoMock, topicMock, userWalletClientMock)

		userWalletClientMock.EXPECT().UserWalletCheck(mock.Anything, mock.Anything).Call.Return(true, nil)
		ordersDaoMock.EXPECT().SelectBuyOrdersToExecute(mock.Anything, mock.Anything).Call.Return(mocks.SelectBuyOrdersMock, nil)
		topicMock.EXPECT().SendEvents(mock.Anything).Call.Return(nil)
		ordersDaoMock.EXPECT().DeleteOrderById(mock.Anything, mock.Anything).Call.Return(nil).Return(nil)
		ordersDaoMock.EXPECT().UpdateOrderQuantityById(mock.Anything, mock.Anything, mock.Anything).Call.Return(nil)

		//when
		err := orderService.PlaceSellOrder(context.Background(), &mocks.SellOrderLessRequestMock)

		//then
		assert.Nil(t, err)
	})

}
