package main

import (
	"porter-management/config"
	"porter-management/internal/job/application"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Connect to Postgres
	db, err := config.NewPgConfig("postgres", "password", "postgres").Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Connect to RabbitMQ
	amqp, err := config.NewAmqpConfig("rabbitmq", "password", "localhost", "5672").Init()
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

	// Initialize JobController
	jobCTRL := application.NewJobController(route, db, ch)
	jobCTRL.RegisterRoutes("/jobs")

	route.Run(":8080")
}
