package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pleum/hexgo/cmd/api/internal/wire"
)

func main() {
	_, cleanup, err := wire.InitializeContainer()
	if err != nil {
		log.Panic(err)
	}
	defer cleanup()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	log.Println("Gracefully shutting down...")
}
