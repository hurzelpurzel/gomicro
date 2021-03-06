package service

/*
https://dzone.com/articles/go-microservices-part-2-building-our-first-service?fromrel=true
*/

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	r := NewRouter()    // NEW
	http.Handle("/", r) // NEW

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
