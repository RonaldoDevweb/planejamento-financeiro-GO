package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RonaldoDevweb/planejamento-financeiro-GO/Model/transaction"
	"github.com/RonaldoDevweb/planejamento-financeiro-GO/util"
)


//GetTransactions blabla
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
   
	w.Header().Set("Content-Type", "application/json")

    

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title: "Salário",
			Amount: 1500.0,
			Type: 0,
			CreatedAt: util.StringToTime("2025-03-26T11:20:05"),
		},

	}

	json.NewEncoder(w).Encode(transactions)

}
//CreateATransaction conver
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)


}