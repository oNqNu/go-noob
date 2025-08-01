package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("htmlTemplate/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "gin で作った api サーバーから html を返せたよ ><",
			"subtitle": "サブタイトルだよ",
		})
	})
	router.Run(":8080")
}
