package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Register(server *echo.Echo) {
	server.GET("/liveness", h.liveness)
	server.GET("/readiness", h.readiness)
}

func (h *HealthCheck) liveness(c echo.Context) error {
	response := make(map[string]string)
	response["status"] = "UP"
	return c.JSON(http.StatusOK, response)
}

func (h *HealthCheck) readiness(c echo.Context) error {
	response := make(map[string]string)
	response["status"] = "OK"
	return c.JSON(http.StatusOK, response)
}
