package service

import (
	"transfer-service/client"
	"transfer-service/kafka"
	"transfer-service/model"
)

type UPITransferService struct {
	AccountsClient *client.AccountsClient
	Producer       *kafka.EventProducer
}

func NewUPITransferService(accountsClient *client.AccountsClient, producer *kafka.EventProducer) *UPITransferService {
	return &UPITransferService{
		AccountsClient: accountsClient,
		Producer:       producer,
	}
}

func (s *UPITransferService) TransferFunds(req model.TransferRequest) model.TransferResponse {

	if err := s.AccountsClient.PostAmount("debit", req.FromAccount, req.Amount); err != nil {
		return model.TransferResponse{Status: "FAILED", Message: "Debit failed: " + err.Error()}
	}

	if err := s.AccountsClient.PostAmount("credit", req.ToAccount, req.Amount); err != nil {
		return model.TransferResponse{Status: "FAILED", Message: "Credit failed: " + err.Error()}
	}

	// ✅ Send Kafka Avro event
	event := model.TransferEvent{
		FromAccount: req.FromAccount,
		ToAccount:   req.ToAccount,
		Amount:      req.Amount,
	}

	if err := s.Producer.SendTransferEvent(event); err != nil {
		return model.TransferResponse{Status: "FAILED", Message: "Kafka publish failed: " + err.Error()}
	}

	return model.TransferResponse{Status: "SUCCESS", Message: "Transfer completed successfully"}

}
