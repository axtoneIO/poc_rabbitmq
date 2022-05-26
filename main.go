package main

import "fmt"


func Run() error{
	fmt.Println("Go RabbitMQ POC")
	return nil
}

func main(){
	if err := Run(); err != nil{
		fmt.Println("Error setting up our application")
		fmt.Println(err)
	}
}