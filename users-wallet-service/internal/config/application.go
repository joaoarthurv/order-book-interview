package config

import (
	"context"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/api"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/dao"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/listener"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/service"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/utils"
)

type Application struct {
	AppName    string
	AppRouters []api.Router
}

type ClientsContainer struct {
	UserWalletService *service.UserWalletService
}

func NewApplication() *Application {
	client := getClients()
	return &Application{
		AppName: "users-wallet-service",
		AppRouters: []api.Router{
			api.NewHealthCheck(),
			api.NewUserWalletApi(client.UserWalletService),
		},
	}
}

func getClients() ClientsContainer {
	return ClientsContainer{
		UserWalletService: producesUserWalletService(producesUserWalletDao()),
	}
}

func producesUserWalletService(userWalletDao *dao.UserWalletDao) *service.UserWalletService {
	return service.NewUserWalletService(userWalletDao)
}

func producesUserWalletDao() *dao.UserWalletDao {
	return dao.NewIUserWalletDao(utils.GetDynamoClient())
}

func producesClientSqs() *utils.Queue {
	return utils.NewQueue(context.Background(), Config().AmazonConfig.Region)
}

func (app *Application) Start() {
	conf := Config()
	go api.Start(api.NewHTTPServer(app.AppRouters), conf.Server.Host.Port)
	go listener.UserWalletListener(producesUserWalletService(producesUserWalletDao()), producesClientSqs())
}
