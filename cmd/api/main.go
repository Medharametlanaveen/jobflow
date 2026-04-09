package main

import (
	"github.com/Medharametlanaveen/jobflow/internal/service"
	"github.com/Medharametlanaveen/jobflow/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	jobService:=service.NewJobService()
	jobHandler:=handler.NewJobHandler(jobService)

	router.GET("/health", jobHandler.Health)

	router.Run(":8080")
}