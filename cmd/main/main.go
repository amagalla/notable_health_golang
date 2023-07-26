package main

import (
	"notable_health/cmd/main/routes"
	"notable_health/pckg/db"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// mysql connection
	db.InitDB()

	router := gin.Default()

	router.Use(cors.Default())

	// Set trusted proxy
	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
