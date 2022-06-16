package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/poncorobbin/go-simple-rest/pkg/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := new(http.ServeMux)
	controllers.New(mux)

	server := new(http.Server)
	server.Addr = ":8090"
	server.Handler = mux

	server.ListenAndServe()
}
