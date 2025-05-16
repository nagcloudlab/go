package model

type TransferRequest struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
}

type TransferResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
