package main

import (
	"log"
	"net/http"
	"transfer-service/client"
	"transfer-service/handler"
	"transfer-service/kafka"
	"transfer-service/service"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	accountsClient := client.NewAccountsClient("http://localhost:8081")
	producer := kafka.NewProducer("localhost:9092", "transfer-events")
	transferSvc := service.NewUPITransferService(accountsClient, producer) // DI
	transferHandler := handler.NewTransferHandler(transferSvc)

	r.HandleFunc("/transfer", transferHandler.Transfer).Methods("POST")

	log.Println("🚀 Transfer service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
