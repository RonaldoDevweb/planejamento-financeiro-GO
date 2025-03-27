package http

import (
	"github.com/RonaldoDevweb/planejamento-financeiro-GO/adapter/http/transaction"
	"net/http"
)

// Init coment
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	// ListenAndServe ; coment
	http.ListenAndServe(":8080", nil)
}
