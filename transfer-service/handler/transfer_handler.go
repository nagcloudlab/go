package handler

import (
	"encoding/json"
	"net/http"
	"transfer-service/model"
	"transfer-service/service"
)

type TransferHandler struct {
	Service *service.UPITransferService
}

func NewTransferHandler(svc *service.UPITransferService) *TransferHandler {
	return &TransferHandler{Service: svc}
}

func (h *TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req model.TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	result := h.Service.TransferFunds(req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
