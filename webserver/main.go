package main

import (
	"log"
	"net/http"
)

const (
	Addr   = ":8080"
	WebDir = "../"
)

func main() {
	server := http.FileServer(http.Dir(WebDir))
	log.Println("WebServer starting on", Addr)
	err := http.ListenAndServe(Addr, server)

	if err != nil {
		log.Fatalln(err)
	}
}
