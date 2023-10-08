package queue

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func PublishMessage(ch *amqp.Channel, QueueName string, message string) {

	ctx := context.Background()
	err := ch.PublishWithContext(
		ctx,
		"",        // exchange
		QueueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	failOnError(err, "Failed to publish a message")

	fmt.Printf(" [x] Sent: %s\n", message)
}

func ConsumeMessage(ch *amqp.Channel, QueueName string) {

	// err := ch.Qos(
	// 	1,     // prefetch count
	// 	0,     // prefetch size
	// 	false, // global
	// )
	// failOnError(err, "Failed to set QoS")

	q, err := ch.QueueDeclare(
		QueueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Printf(" [*] Waiting for messages in %s. To exit press CTRL+C\n", QueueName)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received")
			log.Printf("%s", d.Body)

			// Simulate processing time
			secs := bytes.Count(d.Body, []byte("."))
			time.Sleep(time.Duration(secs) * time.Second)

			log.Printf(" [x] Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
