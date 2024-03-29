package main

import (
	"log"
	"wallet-service/routes"
)

type Server struct {
	Port string
	Host string
}

// entry point of service
func main() {
	e := routes.Router()

	s := Server{
		Port: "8080",
		Host: "localhost",
	}

	log.Fatal(e.Start(s.Host + ":" + s.Port))
}
