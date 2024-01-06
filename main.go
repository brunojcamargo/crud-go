package main

import (
	"crud/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	routes.AppRoutes(app)

	app.Run("localhost:3001")
}
