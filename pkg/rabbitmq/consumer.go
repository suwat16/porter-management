package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	done       chan error
}

func NewConsumer(amqpURI, queueName string) (*Consumer, error) {
	c := &Consumer{
		connection: nil,
		channel:    nil,
		queue:      amqp.Queue{},
		done:       make(chan error),
	}

	var err error

	log.Printf("dialing %s", amqpURI)
	c.connection, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	go func() {
		log.Printf("closing: %s", <-c.connection.NotifyClose(make(chan *amqp.Error)))
	}()

	c.channel, err = c.connection.Channel()
	if err != nil {
		return nil, err
	}

	c.queue, err = c.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}
