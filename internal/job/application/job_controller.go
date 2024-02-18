package controller

import "github.com/gin-gonic/gin"

type JobController struct {
	gin *gin.Engine
}

func NewJobController(gin *gin.Engine) *JobController {
	return &JobController{gin: gin}
}

func (jobController *JobController) CreateNewJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	// ...
}
