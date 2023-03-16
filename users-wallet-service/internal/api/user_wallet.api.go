package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/model"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/service"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/utils"
	"net/http"
)

type UserWalletApi struct {
	userWalletService *service.UserWalletService
}

func NewUserWalletApi(userWalletService *service.UserWalletService) *UserWalletApi {
	return &UserWalletApi{
		userWalletService: userWalletService,
	}
}

func (u *UserWalletApi) Register(echo *echo.Echo) {
	echo.GET("/v1/wallet", u.getUserWallet)
	echo.POST("/v1/wallet/check/balance", u.checkUserBalance)
	echo.POST("/v1/wallet/check/product-quantity", u.checkUserProductQuantity)
}

func (u *UserWalletApi) getUserWallet(ctx echo.Context) error {
	userWallet, err := utils.GenericDecodeRequest[model.UserWallet](ctx)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to decode body, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response, err := u.userWalletService.GetUserWalletByUserIdAndProductType(userWallet.UserId, userWallet.ProductType)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process buy order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (u *UserWalletApi) checkUserBalance(ctx echo.Context) error {
	userWalletCheck, err := utils.GenericDecodeRequest[model.UserWalletCheck](ctx)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to decode body, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	checkResponse, err := u.userWalletService.CheckUserBalance(userWalletCheck)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process buy order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, model.UserWalletCheckResponse{CheckValue: checkResponse})
}

func (u *UserWalletApi) checkUserProductQuantity(ctx echo.Context) error {
	userWalletCheck, err := utils.GenericDecodeRequest[model.UserWalletCheck](ctx)
	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to decode body, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	checkResponse, err := u.userWalletService.CheckUserProductQuantity(userWalletCheck)
	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process buy order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	return ctx.JSON(http.StatusOK, model.UserWalletCheckResponse{CheckValue: checkResponse})
}
