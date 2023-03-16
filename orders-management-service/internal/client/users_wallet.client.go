package client

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/model"
)

const (
	//urlBase                  = "http://localhost:8082"
	urlBase                  = "http://users-wallet-service:8082"
	CheckUserBalancePath     = "/v1/wallet/check/balance"
	CheckUserProductQuantity = "/v1/wallet/check/product-quantity"
)

type UsersWalletClient struct {
	client *resty.Client
}

func NewUsersWalletClient(client *resty.Client) *UsersWalletClient {
	return &UsersWalletClient{
		client: client,
	}
}

func (u *UsersWalletClient) UserWalletCheck(userWalletCheck model.UserWallet, path string) (bool, error) {
	response, err := u.doRequest(userWalletCheck, path)
	if err != nil {
		return false, err
	}
	userWalletResponse, err := unmarshallDecision(response.Body())
	if err != nil {
		return false, err
	}
	return userWalletResponse, nil
}

func (u *UsersWalletClient) doRequest(userWalletCheck model.UserWallet, path string) (*resty.Response, error) {
	response, err := u.client.R().
		SetBody(userWalletCheck).
		Post(urlBase + path)

	if err != nil {
		return response, err
	}

	if response != nil && response.IsError() {
		return response, err
	}

	return response, nil
}

func unmarshallDecision(data []byte) (bool, error) {
	var userWalletResponse model.UserWalletCheckResponse
	if err := json.Unmarshal(data, &userWalletResponse); err != nil {
		return false, err
	}
	return userWalletResponse.CheckValue, nil
}
