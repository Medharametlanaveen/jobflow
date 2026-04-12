package main

import (
	"github.com/Medharametlanaveen/jobflow/internal/db"
	"github.com/Medharametlanaveen/jobflow/internal/handler"
	"github.com/Medharametlanaveen/jobflow/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Connect()

	jobService := service.NewJobService()
	jobHandler := handler.NewJobHandler(jobService)

	router := gin.Default()
	router.GET("/health", jobHandler.Health)

	_ = database // will be used when we set up repository

	router.Run(":8080")
}
