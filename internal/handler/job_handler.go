package handler

import (
	"github.com/Medharametlanaveen/jobflow/internal/service"
	"github.com/gin-gonic/gin"
)

type JobHandler struct{
	service *service.JobService
}

func NewJobHandler(service *service.JobService)*JobHandler{
	return &JobHandler{service: service}
}

func (h *JobHandler)Health(c *gin.Context){
	status:=h.service.Health()
	c.JSON(200, gin.H{"status": status})
}