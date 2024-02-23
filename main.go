package main

import (
	"log"
	"net/http"

	"github.com/uguremirmustafa/go_with_tests/greet"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.MyGreeterHandler)))
}
