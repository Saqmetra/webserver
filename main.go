package main

import (
	"webserver/config"
	"webserver/server"
)

func main() {
	cfg := config.LoadConfig()
	server.StartServer(cfg.Port)
}
