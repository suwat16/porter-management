package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/streadway/amqp"

	"github.com/gin-gonic/gin"

	jobC "porter-management/internal/job/application"
)

func main() {
	route := gin.Default()

	connStr := "user=postgres password=password dbname=postgres sslmode=disable"
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

	jobCTRLRegis := jobC.NewJobController(route, db, ch)
	jobCTRLRegis.RegisterRoutes()

	route.Run(":8080")
}
