package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createRabbitMQConnection() (*amqp.Connection, *amqp.Channel) {
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

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}
