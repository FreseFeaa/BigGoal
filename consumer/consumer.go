package consumer

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Consumer struct {
	UserName    string
	Password    string
	Host        string
	Port        string
	QueueName   string
	ServiceName string
}

func (c *Consumer) Consume() {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", c.UserName, c.Password, c.Host, c.Port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		c.QueueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

}
