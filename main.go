package main

import (
	"evilteccorp.com/database"
	"evilteccorp.com/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := database.InitGDB(); err != nil {
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

	r.POST("/api/project", routes.PostProject)
	r.GET("/api/project", routes.GetProject)
	r.POST("/api/project/solution", routes.PostSolution)

	r.GET("/api/employee", routes.GetEmployee)

	r.POST("/api/register", routes.PostRegister)
	r.POST("/api/token", routes.PostToken)

	if os.Getenv("HTTPS") == "false" {
		log.Fatal(r.Run(":" + os.Getenv("APP_PORT")))
	} else {
		log.Fatal(r.RunTLS(":443",
			os.Getenv("SSL_PEM_PATH"), os.Getenv("SSL_KEY_PATH")))
	}
}
