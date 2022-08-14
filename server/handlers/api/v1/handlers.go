package v1

import (
	"github.com/goiste/web_gui_example/server/handlers/api/v1/sysinfo"

	"github.com/labstack/echo/v4"
)

const basePath = "/api/v1"

func RegisterHandlers(srv *echo.Echo) {
	addURLs(srv, sysinfo.URLs)
}

func addURLs(srv *echo.Echo, urls map[string]map[string]echo.HandlerFunc) {
	for path, handlers := range urls {
		for method, handler := range handlers {
			srv.Add(method, basePath+path, handler)
		}
	}
}
