package main

import (
	"context"
	"log"
	"notification-service/kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	consumer := kafka.NewConsumer("localhost:9092", "transfer-events", "notifier-group")

	ctx, cancel := context.WithCancel(context.Background())
	go consumer.Start(ctx)

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("🔴 Shutting down consumer...")
	cancel()
}
