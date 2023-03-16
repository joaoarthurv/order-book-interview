package main

import (
	"context"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.NewApplication().Start()
	shutDownHook()
}

func shutDownHook() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		log.Println(context.Background(), "Received SIGINT, stopping...")
	case syscall.SIGTERM:
		log.Println(context.Background(), "Received SIGTERM, stopping...")
	}
}
