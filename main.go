package main

import (
	"log"
	"net/http"

	"github.com/olendr/boilerplate/routes"
)

func main() {
	port := "8000"
	r := routes.New()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
