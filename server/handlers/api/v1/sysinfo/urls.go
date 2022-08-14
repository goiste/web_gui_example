package sysinfo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var URLs = map[string]map[string]echo.HandlerFunc{
	"/host": {
		http.MethodGet: GetHostInfo,
	},
	"/cpu": {
		http.MethodGet: GetCPUInfo,
	},
	"/mem": {
		http.MethodGet: GetMemInfo,
	},
}
