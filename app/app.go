package main

import (
	service "app/services"
	"fmt"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767") // NEW
}
