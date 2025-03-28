package http

import (
	"net/http"
    "github.com/RonaldoDevweb/planejamento-financeiro-GO/adapter/http/actuator"
	"github.com/RonaldoDevweb/planejamento-financeiro-GO/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init coment
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())
	// ListenAndServe ; coment
	http.ListenAndServe(":8080", nil)
}
