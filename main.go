package main

import (
	// client "restapi/proto/client"
	"log"
	"os"
	queue "restapi/controllers/queue"
	server "restapi/proto/server"
	"restapi/routers"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
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

func main() {
	conn, ch := createRabbitMQConnection()
	defer conn.Close()
	defer ch.Close()

	go server.StartGrpcServer()
	// go client.StartGrpcClient()
	go queue.PublishMessage(ch, "order_queue", "order_1_cmpl")
	go queue.ConsumeMessage(ch, "order_queue")
	r := routers.SetupRouter()
	r.Run(":8080") // Replace with your desired port
}
