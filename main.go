package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func init() {
}

func index(ctx *gin.Context) {
	bread := []string { "Account", "Checking", "Transactions" }
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Breadcrumbs": bread,
		"Sidebar": gin.H{
			"Title": "Accounts",
			"Selected": 0,
			"List": []gin.H{
				{
					"Title": "Checking",
					"Subtitle": "€ 1123.30",
					"Img": "https://i.pravatar.cc/160",
				},
				{
					"Title": "Savings",
					"Subtitle": "€ 9876.10",
					"Img": "https://i.pravatar.cc/120",
				},
			},
		},
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

  // Start API
  router.Run()
}
