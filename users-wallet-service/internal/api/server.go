package api

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/validation"
)

type Router interface {
	Register(e *echo.Echo)
}

func NewHTTPServer(routers []Router) *echo.Echo {
	ech := echo.New()
	ech.Use(middleware.Recover())
	setServerProperties(ech)
	registerAPI(ech, routers)
	ech.Validator = &validation.CustomValidator{Validator: validator.New()}
	return ech
}

func setServerProperties(e *echo.Echo) {
	e.HideBanner = true
	e.HidePort = true
}

func registerAPI(e *echo.Echo, routers []Router) {
	for i := range routers {
		routers[i].Register(e)
	}
}

func Start(echo *echo.Echo, port string) {
	log.Print(fmt.Sprintf("Starting a api at %s", host(port)))
	echo.Start(host(port))
}

func host(port string) string {
	return fmt.Sprintf("0.0.0.0:%s", port)
}
