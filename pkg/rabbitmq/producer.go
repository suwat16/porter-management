package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	done       chan error
}

func NewProducer(amqpURI, queueName string) (*Producer, error) {
	p := &Producer{
		connection: nil,
		channel:    nil,
		queue:      amqp.Queue{},
		done:       make(chan error),
	}

	var err error

	log.Printf("dialing %s", amqpURI)
	p.connection, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	go func() {
		log.Printf("closing: %s", <-p.connection.NotifyClose(make(chan *amqp.Error)))
	}()

	p.channel, err = p.connection.Channel()
	if err != nil {
		return nil, err
	}

	p.queue, err = p.channel.QueueDeclare(
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

	return p, nil
}
