package config

import "github.com/streadway/amqp"

type AmqpConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
}

func NewAmqpConfig(userName, password, host, port string) *AmqpConfig {
	return &AmqpConfig{
		UserName: userName,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (ac *AmqpConfig) GetConnStr() string {
	return "amqp://" + ac.UserName + ":" + ac.Password + "@" + ac.Host + ":" + ac.Port + "/"
}

func (ac *AmqpConfig) Init() (*amqp.Connection, error) {
	conn, err := amqp.Dial(ac.GetConnStr())
	if err != nil {
		panic("failed to connect to RabbitMQ")
	}

	return conn, nil
}
