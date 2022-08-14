package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/goiste/web_gui_example/server"

	log "github.com/sirupsen/logrus"
)

const (
	minPort     = 1024
	maxPort     = 49151
	defaultPort = 8099
)

//go:embed public
var publicFS embed.FS

func main() {
	port := flag.Uint("p", defaultPort, fmt.Sprintf("A port for the local app server between %d and %d", minPort, maxPort))
	debug := flag.Bool("d", false, "Run in debug mode")
	flag.Parse()

	if *port < minPort || *port > maxPort {
		log.Fatalf("invalid port: %d", *port)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	exitChan := make(chan struct{}, 1)

	srv := server.Init(publicFS, *port, *debug, exitChan)

	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := openGUI(fmt.Sprintf("http://localhost:%d", *port))
		if err != nil {
			log.Errorf("cannot open GUI: %v", err)
			exitChan <- struct{}{}
		}
	}()

	select {
	case <-exitChan:
	case <-quitSignal:
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Debug("shutting down...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

// openGUI opens the specified URL in the default browser of the user's OS
func openGUI(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}
