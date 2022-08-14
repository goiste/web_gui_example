package server

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/goiste/web_gui_example/server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

const pingTimeout = 3 * time.Second

func Init(publicFS embed.FS, port uint, debug bool, exitChan chan<- struct{}) *echo.Echo {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			fmt.Sprintf("http://127.0.0.1:%d", port),
			fmt.Sprintf("http://localhost:%d", port),
		},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))

	server.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	if debug {
		server.Use(middleware.Logger())
	}
	server.Use(middleware.Recover())

	server.GET("/", echo.WrapHandler(handlers.HandleFS(publicFS, "public")))
	server.GET("/assets/*", echo.WrapHandler(
		http.StripPrefix("/assets/", handlers.HandleFS(publicFS, "public/assets")),
	))

	pingChan := make(chan struct{})
	server.GET("/ws", handlers.WSHandler(pingChan))

	handlers.RegisterHandlers(server)

	go handlePing(pingTimeout, pingChan, exitChan)

	go func() {
		if err := server.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			server.Logger.Fatal("shutting down the server...")
		}
	}()

	return server
}

func handlePing(timeout time.Duration, in <-chan struct{}, exit chan<- struct{}) {
	for {
		select {
		case <-in:
		case <-time.NewTimer(timeout).C:
			log.Debug("GUI closed")
			exit <- struct{}{}
		}
	}
}
