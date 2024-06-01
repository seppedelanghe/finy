package main

import (
	"net/http"
	"strconv"
	"time"

	"finy.be/api/structs/viewmodel"
	"finy.be/api/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func init() {
}

func sidebar(ctx *gin.Context) {
	selected := ctx.Param("index")
	index, err := strconv.Atoi(selected)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	sidebar := viewmodel.SidebarVM {
		Title: "Accounts",
		Selected: index,
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

	ctx.HTML(http.StatusOK, "sidebar.tmpl", sidebar)
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
		"Transactions": gin.H {
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
					{ Value: "€ 14,15", Bold: false },
				},
			},
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

	// HTMX
	router.POST("/sidebar/menu/select/:index", sidebar)

	// Start API
	router.Run()
}
