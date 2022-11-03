package main

import (
	"challange/api/services/walletService/routes"
	"log"
)

type Server struct {
	Port string
	Host string
}

// entry point of program
func main() {
	e := routes.Router()

	s := Server{
		Port: "8080",
		Host: "localhost",
	}

	log.Fatal(e.Start(s.Host + ":" + s.Port))
}
