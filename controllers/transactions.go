package controllers

import (
	"net/http"

	"finy.be/api/data"
	"finy.be/api/rendering"
	"finy.be/api/structs"
	"finy.be/api/utils"
)


func GetTransactions(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDatabase("finy")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	var transactions []structs.Transaction
	err = data.SelectQuery(db, "SELECT * FROM Transactions WHERE DeletedAt = -1;", &transactions)
	if err != nil {
		panic(err)
	}

	rendering.TranscationsList(w, &transactions)
}


func AddTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rendering.TransactionsModalAdd(w)
		break;
	case http.MethodPost:
		db, err := data.ConnectDatabase("finy")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		var transaction structs.Transaction
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		
		transaction.New()
		transaction.Name = r.PostForm.Get("name")
		transaction.Amount = utils.StringDecimalInt(r.PostForm.Get("amount"))
		transaction.Date = r.PostForm.Get("date")

		var transactions []structs.Transaction
		transactions = append(transactions, transaction)

		err = data.InsertMany(db, &transactions)
		if err != nil {
			panic(err)
		}

		GetTransactions(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDatabase("finy")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := r.PathValue("id")
	err = data.DeleteTransaction(db, id)
	if err != nil {
		panic(err)
	}
	
	w.WriteHeader(http.StatusNoContent)
}
