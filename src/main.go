package main

import (
	"fmt"
	"net"
	"src/api"
	"src/config"
	"src/db"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// urlSkipper middleware ignores metrics on some route
func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics")
}

func main() {
	echoMainServer := echo.New()
	echoMainServer.Logger.SetLevel(log.Lvl(1))
	echoMainServer.Logger.SetHeader("${time_rfc3339} ${level} ${short_file}:${line} message:${message}")

	p := prometheus.NewPrometheus("echo", urlSkipper)
	p.Use(echoMainServer)

	// Inisialisasi Env
	err := godotenv.Load()
	if err != nil {
		echoMainServer.Logger.Error(err)
	}

	// Inisialisasi DB
	db.Init(echoMainServer, true, true)
	// Inisialisasi Server
	echoMainServer = api.Init(echoMainServer)

	// Server Listener
	port := config.GetConfig(echoMainServer).Port
	echoMainServer.Logger.Fatal(echoMainServer.Start(fmt.Sprintf(":%s", port)))
	echoMainServer.Logger.Info("Port is:", echoMainServer.Listener.Addr().(*net.TCPAddr).Port)
}
