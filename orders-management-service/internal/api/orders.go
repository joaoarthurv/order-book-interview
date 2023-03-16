package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/model"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/service"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/utils"
	"net/http"
)

type Orders struct {
	ordersService *service.OrdersService
}

func NewOrders(ordersService *service.OrdersService) *Orders {
	return &Orders{
		ordersService: ordersService,
	}
}

func (o *Orders) Register(echo *echo.Echo) {
	echo.POST("/v1/order/buy", o.placeBuyOrder)
	echo.POST("/v1/order/sell", o.placeSellOrder)
	echo.GET("/v1/order/buy", o.getBuyOrders)
	echo.GET("/v1/order/sell", o.getSellOrders)

	/*
		echo.PUT("/v1/order/buy/:id", o.updateBuyOrder)
		echo.PUT("/v1/order/sell/:id", o.updateSellOrder)
		echo.PUT("/v1/order/buy/:id", o.cancelBuyOrder)
		echo.PUT("/v1/order/sell/:id", o.cancelSellOrder)
	*/
}

func (o *Orders) placeBuyOrder(ctx echo.Context) error {
	order, err := utils.GenericDecodeRequest[model.OrdersDataRow](ctx)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to decode body, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	err = o.ordersService.PlaceBuyOrder(ctx.Request().Context(), order)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process buy order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, true)
}

func (o *Orders) placeSellOrder(ctx echo.Context) error {
	order, err := utils.GenericDecodeRequest[model.OrdersDataRow](ctx)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to decode body, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	err = o.ordersService.PlaceSellOrder(ctx.Request().Context(), order)

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process sell order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, true)
}

func (o *Orders) getBuyOrders(ctx echo.Context) error {
	ordersResponse, err := o.ordersService.GetBuyOrders(ctx.Request().Context())

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process get buy order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, ordersResponse)
}

func (o *Orders) getSellOrders(ctx echo.Context) error {
	ordersResponse, err := o.ordersService.GetSellOrders(ctx.Request().Context())

	if err != nil {
		response := make(map[string]string)
		response["error"] = fmt.Sprintf("Error to process get sell order, %v", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	return ctx.JSON(http.StatusOK, ordersResponse)
}
