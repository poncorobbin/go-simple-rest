package main

import "net/http"
import "github.com/poncorobbin/go-simple-rest/pkg/controllers"

func main() {
    controllers.New()

    server := new(http.Server)
    server.Addr = ":8090"

    server.ListenAndServe()
}
