package model

type TransferEvent struct {
	FromAccount string  `avro:"from_account"`
	ToAccount   string  `avro:"to_account"`
	Amount      float64 `avro:"amount"`
}
