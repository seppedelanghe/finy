package main

import (
	"net/http"
	"strconv"

	"finy.be/api/ws"
	"finy.be/api/rendering"
	"finy.be/api/structs/viewmodel"
)

type H = map[any]any


func init() {
}


func sidebar(w http.ResponseWriter, r *http.Request) {
	selected := r.PathValue("index")
	index, err := strconv.Atoi(selected)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sidebar := viewmodel.SidebarVM {
		Selected: 1,
		SelectedMenu: index,
		Pages: []viewmodel.SidebarPageVM {
			{ Icon: "bell.svg", Name: "Notifications" },
			{ Icon: "chart.svg", Name: "Accounts" },
			{ Icon: "settings.svg", Name: "Settings" },
		},
		Menus: []viewmodel.SidebarMenuVM {
			{ Title: "Checking", Subtitle: "1114.13", Icon: "" },
			{ Title: "Savings", Subtitle: "9876.10", Icon: "" },
		},
	}

	rendering.Renderer.HTML(w, http.StatusOK, "sidebar", sidebar)
	rendering.Renderer.Text(w, http.StatusOK, "<p id=\"transactions\">Swapped</p>")
}


func index(w http.ResponseWriter, r *http.Request) {
	bread := []string { "Account", "Checking", "Transactions" }
	sidebar := viewmodel.SidebarVM {
		Selected: 1,
		SelectedMenu: 0,
		Pages: []viewmodel.SidebarPageVM {
			{ Icon: "bell.svg", Name: "Notifications" },
			{ Icon: "chart.svg", Name: "Accounts" },
			{ Icon: "settings.svg", Name: "Settings" },
		},
		Menus: []viewmodel.SidebarMenuVM {
			{ Title: "Checking", Subtitle: "1114.13", Icon: "" },
			{ Title: "Savings", Subtitle: "9876.10", Icon: "" },
		},
	}

	data := H {
		"Breadcrumbs": bread,
		"Sidebar": sidebar,
		"Transactions": H {
			"Title": "Transactions",
			"Subtitle": "23 total",
			"Table": viewmodel.TableVM {
				Headers: []viewmodel.TableHeaderVM {
					{ Name: "Name" },
					{ Name: "Date" },
					{ Name: "Amount" },
				},
				Items: []viewmodel.TableItemVM {
					{ Value: "Food", Bold: true },
					{ Value: "31-05-2024", Bold: false },
					{ Value: "14,15", Bold: false },
				},
			},
		},
	}
		
	rendering.Renderer.HTML(w, http.StatusOK, "index", data)
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
    p := "." + r.URL.Path
    http.ServeFile(w, r, p)
}

func main() {
	router := http.NewServeMux()
	
	router.HandleFunc("GET /assets/", serveFiles)
	router.HandleFunc("GET /", index)
	router.HandleFunc("GET /ws", ws.HandleWebSocket)

	router.HandleFunc("POST /sidebar/menu/select/{index}", sidebar)

	// Start API
	server := &http.Server{ Addr: "localhost:8080", Handler: router }
	server.ListenAndServe()
}

