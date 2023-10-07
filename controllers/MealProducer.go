package controllers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetMyMealProducer() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("env unable to load")
	}

	User := os.Getenv("RABBIT_MQ_USERNAME")
	Pwd := os.Getenv("RABBIT_MQ_PASSWORD")
	Port := os.Getenv("RABBIT_MQ_PORT")
	Host := os.Getenv("RABBIT_MQ_SERVER")

	dialStr := "amqp://" + User + ":" + Pwd + "@" + Host + ":" + Port + "/"

	conn, err := amqp.Dial(dialStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueName := "order_queue"
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	q, err := ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Printf(" [*] Waiting for messages in %s. To exit press CTRL+C\n", queueName)

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
