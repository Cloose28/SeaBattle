package main

import (
	"SeaBattle/battle"
	"SeaBattle/web"
	"flag"
	"log"
)

func main() {
	const defaultPort = 4040
	listenPort := *flag.Int("port", defaultPort, "server port")

	model := battle.CreateSeaBattle()

	server := web.NewGinServer(model)

	if err := server.Run(listenPort); err != nil {
		log.Fatalf("api: failed to start server %v", err)
	}
}
