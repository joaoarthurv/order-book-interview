package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GenericDecodeRequest[K any](ctx echo.Context) (*K, error) {
	var body K
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		log.Error("Error while parsing request: ", err)
		return nil, err
	}
	return &body, nil
}
