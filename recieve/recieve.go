package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Hello",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "fail to register a consumer")
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("recieved a message: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages. To exit, press CTRL + c")
	<-forever

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
