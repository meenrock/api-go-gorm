package main

import (
	"restapi/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080") // Replace with your desired port
}
