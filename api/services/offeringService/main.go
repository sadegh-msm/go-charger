package main

import (
	"log"
	"offering-service/routes"
)

type Server struct {
	Port string
	Host string
}

// entry point of service
func main() {
	e := routes.Router()

	s := Server{
		Port: "8081",
		Host: "localhost",
	}

	log.Fatal(e.Start(s.Host + ":" + s.Port))
}
