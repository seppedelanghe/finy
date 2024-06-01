package main

import (
	"net/http"
	"time"

	"finy.be/api/structs/viewmodel"
	"finy.be/api/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func init() {
}

func index(ctx *gin.Context) {
	bread := []string { "Account", "Checking", "Transactions" }
	sidebar := viewmodel.SidebarVM {
		Title: "Accounts",
		Selected: 0,
		Pages: []viewmodel.SidebarPageVM {
			{ Icon: "", Name: "" },
			{ Icon: "", Name: "" },
			{ Icon: "", Name: "" },
			{ Icon: "", Name: "" },
		},
		Menus: []viewmodel.SidebarMenuVM {
			{ Title: "Checking", Subtitle: "€ 1114.13", Icon: "" },
			{ Title: "Savings", Subtitle: "€ 9876.10", Icon: "" },
		},
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Breadcrumbs": bread,
		"Sidebar": sidebar,
		"Table": gin.H{
			"Title": "Transactions",
			"Subtitle": "23 total",
		},
	})
}

func main() {
  router := gin.Default()
  router.MaxMultipartMemory = 8 << 20 // 8 MB
  router.Static("/assets", "./assets")
  router.LoadHTMLGlob("templates/*")

  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "POST"},
    AllowHeaders:     []string{"*"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
  }))

  // No auth
  router.GET("/", index)
  router.GET("/ws", ws.HandleWebSocket)

  // Start API
  router.Run()
}
