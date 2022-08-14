package handlers

import (
	"github.com/goiste/web_gui_example/server/handlers/api/v1"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(srv *echo.Echo) {
	v1.RegisterHandlers(srv)
}
