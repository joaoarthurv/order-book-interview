package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/client"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/dao"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/model"
)

type OrdersService struct {
	ordersDao        dao.IOrdersDao
	topicService     ITopicService
	userWalletClient client.IUsersWalletClient
}

type IOrdersService interface {
	PlaceBuyOrder(ctx context.Context, orderRequest *model.OrdersDataRow) error
	PlaceSellOrder(ctx context.Context, orderRequest *model.OrdersDataRow) error
}

func NewOrdersService(ordersDao dao.IOrdersDao, topicService ITopicService, userWalletClient client.IUsersWalletClient) *OrdersService {
	return &OrdersService{
		ordersDao:        ordersDao,
		topicService:     topicService,
		userWalletClient: userWalletClient,
	}
}

func (o *OrdersService) PlaceBuyOrder(ctx context.Context, orderRequest *model.OrdersDataRow) error {
	checkResponse, err := o.userWalletClient.UserWalletCheck(
		model.UserWallet{
			UserId:      orderRequest.OrderOwnerId,
			Price:       orderRequest.OrderPrice * float64(orderRequest.OrderQuantity),
			ProductType: orderRequest.OrderProductType,
		}, client.CheckUserBalancePath)

	if err != nil {
		return err
	}
	if checkResponse == false {
		errMessage := fmt.Sprintf("user: %v not enough balance: %v", orderRequest.OrderOwnerId, orderRequest.OrderProductType)
		return errors.New(errMessage)
	}

	ordersDatabase, err := o.hasOrderToSell(ctx, orderRequest)
	if err != nil {
		log.Error("error to find order sell")
		return err
	}

	orderRequest.OrderType = "BUY"
	orderRequest.OrderId = uuid.New().String()

	err = o.validateAndExecuteOrders(ctx, orderRequest, ordersDatabase)
	if err != nil {
		log.Error("error to validate and execute orders")
		return err
	}

	if orderRequest.OrderQuantity > 0 {
		err := o.ordersDao.InsertOrder(ctx, orderRequest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *OrdersService) PlaceSellOrder(ctx context.Context, orderRequest *model.OrdersDataRow) error {
	checkResponse, err := o.userWalletClient.UserWalletCheck(
		model.UserWallet{
			UserId:          orderRequest.OrderOwnerId,
			ProductQuantity: orderRequest.OrderQuantity,
			ProductType:     orderRequest.OrderProductType,
		}, client.CheckUserProductQuantity)

	if err != nil {
		return err
	}
	if checkResponse == false {
		errMessage := fmt.Sprintf("user: %v, not enough quantity of product: %v", orderRequest.OrderOwnerId, orderRequest.OrderProductType)
		return errors.New(errMessage)
	}

	ordersBuy, err := o.hasOrderToBuy(ctx, orderRequest)
	if err != nil {
		log.Error("error to find order buy")
		return err
	}

	orderRequest.OrderType = "SELL"
	orderRequest.OrderId = uuid.New().String()

	err = o.validateAndExecuteOrders(ctx, orderRequest, ordersBuy)
	if err != nil {
		log.Error("error to validate and execute orders")
		return err
	}

	if orderRequest.OrderQuantity > 0 {
		err = o.ordersDao.InsertOrder(ctx, orderRequest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *OrdersService) hasOrderToSell(ctx context.Context, orderRequest *model.OrdersDataRow) ([]model.OrdersDataRow, error) {
	orders, err := o.ordersDao.SelectSellOrdersToExecute(ctx, orderRequest)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrdersService) hasOrderToBuy(ctx context.Context, orderRequest *model.OrdersDataRow) ([]model.OrdersDataRow, error) {
	orders, err := o.ordersDao.SelectBuyOrdersToExecute(ctx, orderRequest)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrdersService) validateAndExecuteOrders(ctx context.Context, orderRequest *model.OrdersDataRow, ordersDatabase []model.OrdersDataRow) error {
	if len(ordersDatabase) > 0 {
		err := o.executeOrders(ctx, orderRequest, ordersDatabase)

		if err != nil {
			log.Error("error to execute orders")
			return err
		}
	}
	return nil
}

func (o *OrdersService) executeOrders(ctx context.Context, orderRequest *model.OrdersDataRow, ordersDatabase []model.OrdersDataRow) error {
	for _, orderDatabase := range ordersDatabase {
		if orderRequest.OrderQuantity >= orderDatabase.OrderQuantity {
			orderRequest.OrderQuantity = orderRequest.OrderQuantity - orderDatabase.OrderQuantity
		} else {
			orderDatabase.OrderQuantity = orderDatabase.OrderQuantity - orderRequest.OrderQuantity
			err := o.SendExecutionToTopic(orderRequest, &orderDatabase, orderRequest.OrderQuantity)
			if err != nil {
				return err
			}
			orderRequest.OrderQuantity = 0
			err = o.ordersDao.UpdateOrderQuantityById(ctx, orderDatabase.OrderId, orderDatabase.OrderQuantity)
			if err != nil {
				log.Errorf("error while update order quantity by Id when execute orders, orderId: %v", orderDatabase.OrderId)
				return err
			}
			return nil
		}
		err := o.SendExecutionToTopic(orderRequest, &orderDatabase, orderDatabase.OrderQuantity)
		if err != nil {
			return err
		}
		o.ordersDao.DeleteOrderById(ctx, orderDatabase.OrderId)
	}
	return nil
}

func (o *OrdersService) GetBuyOrders(ctx context.Context) ([]model.OrdersDataRow, error) {
	ordersResponse, err := o.ordersDao.SelectBuyOrdersByOrderType(ctx)

	if err != nil {
		log.Error("error while select buy orders")
		return nil, err
	}

	return ordersResponse, nil
}

func (o *OrdersService) GetSellOrders(ctx context.Context) ([]model.OrdersDataRow, error) {
	ordersResponse, err := o.ordersDao.SelectSellOrdersByOrderType(ctx)

	if err != nil {
		log.Error("error while select buy orders")
		return nil, err
	}

	return ordersResponse, nil
}

func (o *OrdersService) SendExecutionToTopic(orderRequest *model.OrdersDataRow, orderDatabase *model.OrdersDataRow, orderQuantity int) error {
	var err error
	var orderExecutionBody model.OrderExecutedBody
	if orderRequest.OrderType == "BUY" {
		err = o.topicService.SendEvents(orderExecutionBody.BuildOrderExecutedBody(orderRequest, orderDatabase, orderQuantity))
	} else {
		err = o.topicService.SendEvents(orderExecutionBody.BuildOrderExecutedBody(orderDatabase, orderRequest, orderQuantity))
	}
	if err != nil {
		return err
	}
	return nil
}
