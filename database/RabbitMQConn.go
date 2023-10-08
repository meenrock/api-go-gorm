package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createRabbitMQConnection() *amqp.Channel {
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

	return ch
}
