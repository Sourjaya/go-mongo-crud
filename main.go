package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sourjaya/go-mongo-crud/pkg/routes"
)

func main() {
	fmt.Printf("Listening at port 8080 ...")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
