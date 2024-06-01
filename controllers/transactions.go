package controllers

import (
	"fmt"
	"net/http"

	"finy.be/api/rendering"
	"finy.be/api/structs"
	"finy.be/api/structs/viewmodel"
)

var transactions = make([]structs.Transaction, 100)


func GetTransactions(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.TransactionsVM {
		Title: "Transactions",
		Subtitle: fmt.Sprintf("%d total", len(transactions)),
		Table: viewmodel.TableVM {
			Headers: []viewmodel.TableHeaderVM {
				{ Name: "Name" },
				{ Name: "Date" },
				{ Name: "Amount" },
			},
			Rows: []viewmodel.TableRow {
				{ 
					Data: []viewmodel.TableDataVM {
						{ Value: "Food", Bold: true },
						{ Value: "31-05-2024", Bold: false },
						{ Value: "14,15", Bold: false },
					},
				},
			},
		},
	}

	rendering.Renderer.HTML(w, http.StatusOK, "transactions", vm)
}


func ModalAddTransaction(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.ModalAddTransactionVM{}

	rendering.Renderer.HTML(w, http.StatusOK, "modals/transaction_add", vm)
}
