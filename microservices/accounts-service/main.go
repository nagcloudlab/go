package main

import (
	"accounts-service/handler"
	"accounts-service/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	store := store.NewAccountStore()
	h := handler.NewAccountHandler(store) // DI

	r.HandleFunc("/accounts", h.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", h.GetAccount).Methods("GET")
	r.HandleFunc("/accounts/{id}/credit", h.CreditAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}/debit", h.DebitAccount).Methods("POST")

	log.Println("✅ Accounts service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))

}
