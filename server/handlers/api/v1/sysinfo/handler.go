package sysinfo

import (
	"net/http"

	"github.com/goiste/web_gui_example/src/sysinfo"

	"github.com/labstack/echo/v4"
)

func GetHostInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"host": sysinfo.GetHostInfo(),
	})
}

func GetCPUInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cpu": sysinfo.GetCPUInfo(),
	})
}

func GetMemInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"mem": sysinfo.GetMemInfo(),
	})
}
