package main

import (
	"fmt"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/routes"
)

func main() {

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes.RegisterRoutes(),
	}

	fmt.Println("Server Started")

	server.ListenAndServe()
}
