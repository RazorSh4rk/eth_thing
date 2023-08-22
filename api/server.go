package api

import (
	"fmt"
	"net/http"
)

func Run(parser *Parser) {
	fmt.Println("Setting up API server")

	// /currentblock
	http.HandleFunc("/currentblock", HandleGetCurrentBlock(parser))

	// /subscribe?address=0xDEADBEEF
	http.HandleFunc("/subscribe", HandleSubscribe(parser))

	// /transactions?address=0xDEADBEEF
	http.HandleFunc("/transactions", HandleGetTransactions(parser))

	// /subs
	http.HandleFunc("/subs", HandleAllSubs)

	http.ListenAndServe(":8080", nil)
}
