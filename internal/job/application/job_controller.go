package controller

import (
	"github.com/gin-gonic/gin"
)

type JobController struct {
	// using job use case
}

func InitJobController(router *gin.Engine) {

	jobController := router.Group("/jobs")
	{
		jobController.POST("/create", CreateJob)
	}
}

func CreateJob(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": "success",
		"data":   "{}",
	})
}
