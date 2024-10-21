package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ClaytonMatos84/go-rabbit/internal"
)

func main() {
	fmt.Println("Publish message on queue.")

	conn := internal.CreateConnection()
	defer conn.Close()

	ch := internal.GetChannel(conn)
	defer ch.Close()
	internal.ConfigureExchange(ch)

	// call: go run <class> <routingKey> <message>
	if len(os.Args) < 3 {
		log.Printf("Usage 2 parameters for publish message on RabbitMQ")
		os.Exit(0)
	}
	routingKey := os.Args[1]
	message := os.Args[2]

	internal.PublishMessage(ch, routingKey, message)
	fmt.Println("Published message on queue.", message)
}
