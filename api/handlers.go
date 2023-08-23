package api

import (
	"encoding/json"
	"eth_tracker/eth"
	"eth_tracker/storage"
	"fmt"
	"net/http"
)

type BlockResponse struct {
	Id uint64 `json:"id"`
}

type TransactionsResponse struct {
	Transactions []eth.EthereumTransaction `json:"transactions"`
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func HandleGetCurrentBlock(parser *Parser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		block := parser.GetCurrentBlock()

		response := BlockResponse{Id: block}
		jsonRes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
	}
}

func HandleSubscribe(parser *Parser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		queryParams := r.URL.Query()
		address := queryParams.Get("address")

		if address != "" {
			fmt.Println("Added a sub for ", address)
			sub_success := parser.Subscribe(address)

			if !sub_success {
				http.Error(w, "Can't sub with that address", http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			http.Error(w, "Add a /address to subscribe", http.StatusBadRequest)
		}
	}
}

func HandleGetTransactions(parser *Parser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		queryParams := r.URL.Query()
		address := queryParams.Get("address")

		if address != "" {
			trans := parser.GetTransactions(address)
			resp := TransactionsResponse{Transactions: trans}
			jsonRes, err := json.Marshal(resp)

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		} else {
			http.Error(w, "Add a /address to get transactions", http.StatusBadRequest)
		}
	}
}

func HandleAllSubs(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	addr := storage.GetWatchedAddresses()
	jsonRes, err := json.Marshal(addr)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
