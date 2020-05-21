package main

import (
	"fmt"
	service "hello/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767") // NEW
}
