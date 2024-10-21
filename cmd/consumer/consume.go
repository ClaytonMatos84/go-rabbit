package main

import (
	"fmt"
	"log"

	"github.com/ClaytonMatos84/go-rabbit/internal"
)

func main() {
	fmt.Println("Start consume message on rabbit.")

	conn := internal.CreateConnection()
	defer conn.Close()

	ch := internal.GetChannel(conn)
	defer ch.Close()
	internal.ConfigureExchange(ch)

	internal.ConfigureQueue(ch, internal.BUYER_QUEUE_NAME, internal.BUYER_BINDING_NAME)
	internal.ConfigureQueue(ch, internal.SELLER_QUEUE_NAME, internal.SELLER_BINDING_NAME)
	internal.ConfigureQueue(ch, internal.PAYMENT_QUEUE_NAME, internal.PAYMENT_BINDING_NAME)

	buyerMessages := internal.ConsumeMessage(ch, internal.BUYER_QUEUE_NAME)
	sellerMessages := internal.ConsumeMessage(ch, internal.SELLER_QUEUE_NAME)
	paymentMessages := internal.ConsumeMessage(ch, internal.PAYMENT_QUEUE_NAME)

	forever := make(chan bool)

	go func() {
		for message := range buyerMessages {
			log.Println("Receive buyer message:", string(message.Body))
		}
	}()

	go func() {
		for message := range sellerMessages {
			log.Println("Receive seller message:", string(message.Body))
		}
	}()

	go func() {
		for message := range paymentMessages {
			log.Println("Receive payment message:", string(message.Body))
		}
	}()

	fmt.Println("Waiting messages on queue")
	<-forever
}
