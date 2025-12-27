package main

import (
	"witty-server/internal/config"
	"witty-server/server"
)

func main() {
	appConfig := config.LoadConfig()

	server.NewServer(appConfig).Start()
}
