package rendering

import (
	"fmt"
	"net/http"

	"finy.be/api/structs"
	"finy.be/api/structs/viewmodel"
)

func TranscationsList(w http.ResponseWriter, transactions *[]structs.Transaction) error {
	var rows []viewmodel.TableRow
	for _, t := range *transactions {
		rows = append(rows, viewmodel.TableRow{
			Data: []viewmodel.TableDataVM {
				{ Value: t.Name, Bold: true },
				{ Value: t.Date, Bold: false },
				{ Value: fmt.Sprintf("%.2f", float32(t.Amount) / 100), Bold: false },
			},
		})
	}

	vm := viewmodel.TransactionsVM {
		Title: "Transactions",
		Subtitle: fmt.Sprintf("%d total", len(*transactions)),
		Table: viewmodel.TableVM {
			Headers: []viewmodel.TableHeaderVM {
				{ Name: "Name" },
				{ Name: "Date" },
				{ Name: "Amount" },
			},
			Rows: rows,
		},
	}

	return Renderer.HTML(w, http.StatusOK, "transactions", vm)
}

func TransactionsModalAdd(w http.ResponseWriter) error {
	vm := viewmodel.ModalAddTransactionVM{}

	return Renderer.HTML(w, http.StatusOK, "modals/transactions/add", vm)
}

