package main

import (
	"SeaBattle/battle"
	"SeaBattle/web"
	"flag"
	"log"
)

func main() {
	listenPort := *flag.Int("port", 8080, "server port")

	model := battle.CreateSeaBattle()

	server := web.NewGinServer(model)

	if err := server.Run(listenPort); err != nil {
		log.Fatalf("api: failed to start server %v", err)
	}
}
