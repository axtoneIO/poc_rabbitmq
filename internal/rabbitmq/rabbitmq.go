package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Service interface{
	Connect() error
}

type RabbitMQ struct{
	Conn *amqp.Connection
}

func (r *RabbitMQ) Connect() error{
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to RabbitMQ")
	return nil
}

func NewRabbitMQService() *RabbitMQ{
	return &RabbitMQ{}
}