package controllers

import (
	"net/http"

	"finy.be/api/rendering"
	"finy.be/api/structs"
	"finy.be/api/utils"
)

var transactions []structs.Transaction


func GetTransactions(w http.ResponseWriter, r *http.Request) {
	rendering.TranscationsList(w, &transactions)
}


func AddTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rendering.TransactionsModalAdd(w)
		break;
	case http.MethodPost:
		var transaction structs.Transaction
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		
		transaction.Name = r.PostForm.Get("name")
		transaction.Amount = utils.StringDecimalInt(r.PostForm.Get("amount"))
		transaction.Date = r.PostForm.Get("date")

		transactions = append(transactions, transaction)

		err = rendering.TranscationsList(w, &transactions)
		if err != nil {
			panic(err)
		}
		break;
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
