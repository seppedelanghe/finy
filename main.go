package main

import (
	"net/http"

	"finy.be/api/controllers"
	"finy.be/api/rendering"
	"finy.be/api/ws"
)

type H = map[any]any


func init() {
}



func index(w http.ResponseWriter, r *http.Request) {
	rendering.Renderer.HTML(w, http.StatusOK, "index", nil)
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
	compRouter.HandleFunc("GET /modal/transaction/add", controllers.ModalAddTransaction)

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

