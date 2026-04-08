package main

import (

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run(":8080")
}