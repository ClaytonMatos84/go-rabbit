package main

import (
	"fmt"

	"github.com/ClaytonMatos84/go-rabbit/internal"
)

func main() {
	fmt.Println("Test connection with RabbitMQ in go.")
	conn := internal.CreateConnection()
	fmt.Println("Connect RabbitMQ in go completed with success.")
	defer conn.Close()
	fmt.Println("Connection executed with success.")
}
