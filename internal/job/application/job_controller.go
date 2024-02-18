package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type JobController struct {
	gin *gin.Engine
	ch  *amqp.Channel
}

func NewJobController(gin *gin.Engine, ch *amqp.Channel) *JobController {
	return &JobController{gin: gin, ch: ch}
}

func (jobController *JobController) RegisterRoutes() {
	jobController.gin.GET("/job", jobController.GetJobs)
}

func (jobController *JobController) GetJobs(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
