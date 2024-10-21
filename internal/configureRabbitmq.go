package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

const EXCHANGE_NAME string = "app_exchange"
const BUYER_QUEUE_NAME string = "buyer"
const SELLER_QUEUE_NAME string = "seller"
const PAYMENT_QUEUE_NAME string = "payment"
const BUYER_BINDING_NAME string = "buy"
const SELLER_BINDING_NAME string = "sell"
const PAYMENT_BINDING_NAME string = "pay"

func CreateConnection() *amqp091.Connection {
	err := godotenv.Load()
	CheckError(err, "Error on loading env")

	user := os.Getenv("RABBIT_USER")
	pass := os.Getenv("RABBIT_PASS")

	path := fmt.Sprintf("amqp://%s:%s@localhost:5672/", user, pass)
	conn, err := amqp091.Dial(path)
	CheckError(err, "Failed connect to RabbitMQ")

	return conn
}

func GetChannel(conn *amqp091.Connection) *amqp091.Channel {
	ch, err := conn.Channel()
	CheckError(err, "Error open channel")

	return ch
}

func ConfigureExchange(ch *amqp091.Channel) {
	ch.ExchangeDeclare(
		EXCHANGE_NAME, // name
		"direct",      // type
		false,         // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
}

func PublishMessage(ch *amqp091.Channel, routingKey string, message string) {
	ch.Publish(
		EXCHANGE_NAME, // exchange
		routingKey,    // routing key
		false,         // mandatory
		false,         // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		})
}

func ConfigureQueue(ch *amqp091.Channel, queueName string, bindingKey string) {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	CheckError(err, "Failed on create queue")
	bindingQueue(ch, q.Name, bindingKey)
}

func bindingQueue(ch *amqp091.Channel, queueName string, bindingKey string) {
	ch.QueueBind(
		queueName,     // queue name
		bindingKey,    // routing key
		EXCHANGE_NAME, // exchange
		false,
		nil,
	)
}

func ConsumeMessage(ch *amqp091.Channel, queueName string) <-chan amqp091.Delivery {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	CheckError(err, "Error on consume messages")

	return msgs
}
