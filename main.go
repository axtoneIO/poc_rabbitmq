package main

import (
	"fmt"
	"github.com/axtoneIO/poc_rabbitmq/internal/rabbitmq"
)

type App struct{
	Rmq *rabbitmq.RabbitMQ
}

func Run() error{
	fmt.Println("Go RabbitMQ POC")
	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("Hello World")
	if err != nil{
		return err
	}

	app.Rmq.Consume()

	return nil
}

func main(){
	if err := Run(); err != nil{
		fmt.Println("Error setting up our application")
		fmt.Println(err)
	}
}