package main

import (
	client "restapi/proto/client"
	server "restapi/proto/server"
	"restapi/routers"
)

func main() {
	go server.StartGrpcServer()
	go client.StartGrpcClient()
	r := routers.SetupRouter()
	r.Run(":8080") // Replace with your desired port
}
