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
	//Exchange Declare
	err := jobController.ch.ExchangeDeclare("job", "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	jobController.gin.GET("/job", jobController.GetJobs)
}

func (jobController *JobController) GetJobs(c *gin.Context) {
	// Publish a message to the exchange job with the routing key create-job

	// Message is object
	message := `{"name": "Job 1", "description": "This is job 1"}`

	// Publish a message to the exchange job with the routing key create-job
	err := jobController.ch.Publish("job", "create-job", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(message),
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
