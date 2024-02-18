package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"

	jobC "porter-management/internal/job/application"
)

func main() {
	route := gin.Default()

	// Connect to RabbitMQ
	amqp, err := amqp.Dial("amqp://rabbitmq:password@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer amqp.Close()

	// Create a channel
	ch, err := amqp.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	jobCTRLRegis := jobC.NewJobController(route, ch)
	jobCTRLRegis.RegisterRoutes()

	route.Run(":8080")
}
