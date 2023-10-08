package main

import (
	// client "restapi/proto/client"

	queue "restapi/controllers/queue"
	conn "restapi/database"
	server "restapi/proto/server"
	"restapi/routers"
)

func main() {
	conn, ch := conn.CreateRabbitMQConnection()
	defer conn.Close()
	defer ch.Close()

	go server.StartGrpcServer()
	// go client.StartGrpcClient()
	go queue.PublishMessage(ch, "order_queue", "order_1_cmpl")
	go queue.ConsumeMessage(ch, "order_queue")
	r := routers.SetupRouter()
	r.Run(":8080") // Replace with your desired port
}
