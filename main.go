package main

import (
	"net/http"

	"finy.be/api/controllers"
	"finy.be/api/rendering"
	"finy.be/api/ws"
)

type H = map[any]any


func init() {
	// err := data.EnsureDbExists("finy")
	// if err != nil {
	// 	panic(err)
	// }
}


func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":	
		rendering.Renderer.HTML(w, http.StatusOK, "index", nil)
		break

	case "/favicon.ico":
		http.Redirect(w, r, "/assets/favicon.ico", 303)
		break

	default:
		rendering.Renderer.Text(w, http.StatusNotFound, "Not Found")
	}
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
    p := "." + r.URL.Path
    http.ServeFile(w, r, p)
}

func main() {
	mainRouter := http.NewServeMux()
	
	mainRouter.HandleFunc("GET /", index)
	mainRouter.HandleFunc("GET /assets/", serveFiles)

	compRouter := http.NewServeMux()
	compRouter.HandleFunc("GET /sidebar", controllers.GetSidebar)
	compRouter.HandleFunc("POST /sidebar/menu/select/{index}", controllers.PostSidebar)
	compRouter.HandleFunc("GET /transactions", controllers.GetTransactions)
	compRouter.HandleFunc("GET /modal/transaction/add", controllers.AddTransaction)
	compRouter.HandleFunc("POST /modal/transaction/add", controllers.AddTransaction)
	compRouter.HandleFunc("GET /modal/empty/{id}", controllers.GetEmptyModal)


	// Main router
	router := http.NewServeMux()
	router.Handle("/", mainRouter)
	router.Handle("/comp/", http.StripPrefix("/comp", compRouter))
	
	// dev routes
	router.HandleFunc("GET /ws", ws.HandleWebSocket)

	// Start API
	server := &http.Server{ Addr: "localhost:8080", Handler: router }
	server.ListenAndServe()
}

