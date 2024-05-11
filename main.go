package main

import (
	"evilteccorp.com/database"
	"evilteccorp.com/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	err := database.InitGDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := gin.Default()

	r.Static("/public", "./public")

	r.LoadHTMLGlob("views/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/apply", func(c *gin.Context) {
		c.HTML(http.StatusOK, "apply.html", nil)
	})

	r.GET("/api", func(c *gin.Context) {
		c.HTML(http.StatusOK, "api.html", nil)
	})

	r.POST("/api/register", routes.PostRegister)
	r.POST("/api/token", routes.PostToken)

	r.Run(":3000")
}
