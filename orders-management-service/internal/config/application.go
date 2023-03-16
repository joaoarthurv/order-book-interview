package config

import (
	"context"
	"github.com/go-resty/resty/v2"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/api"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/client"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/dao"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/service"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/utils"
)

type Application struct {
	AppName    string
	AppRouters []api.Router
}

type ClientsContainer struct {
	OrdersDao     *dao.OrdersDao
	OrdersService *service.OrdersService
}

func NewApplication() *Application {
	clients := getClients()
	return &Application{
		AppName: config().AppName,
		AppRouters: []api.Router{
			api.NewHealthCheck(),
			api.NewOrders(clients.OrdersService),
		},
	}
}

func getClients() ClientsContainer {
	conf := config()
	return ClientsContainer{
		OrdersService: producesOrdersService(producesOrdersDao(conf), conf),
	}
}

func producesOrdersService(ordersDao *dao.OrdersDao, config Configuration) *service.OrdersService {
	return service.NewOrdersService(ordersDao, producesTopicClient(config), producesUserWalletClient())
}

func producesUserWalletClient() *client.UsersWalletClient {
	return client.NewUsersWalletClient(resty.New())
}

func producesOrdersDao(config Configuration) *dao.OrdersDao {
	writePool, _ := config.PostgresConfig.CreateWritePool()
	readPool, _ := config.PostgresConfig.CreateReadPool()
	return dao.NewOrdersDao(writePool, readPool)
}

func producesTopicClient(config Configuration) *service.TopicService {
	region := config.AmazonConfig.Region
	snsArn := config.AmazonConfig.OrdersExecutedSnsArn
	snsClient := utils.NewTopicClient(context.Background(), region)
	return service.NewTopicService(snsClient, snsArn)
}

func (app *Application) Start() {
	conf := config()
	go api.Start(api.NewHTTPServer(app.AppRouters), conf.Server.Host.Port)
}
