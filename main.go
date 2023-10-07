package main

import (
	// client "restapi/proto/client"
	producer "restapi/controllers"
	server "restapi/proto/server"
	"restapi/routers"
)

func main() {
	go server.StartGrpcServer()
	// go client.StartGrpcClient()
	go producer.GetMyMealProducer()
	r := routers.SetupRouter()
	r.Run(":8080") // Replace with your desired port
}
