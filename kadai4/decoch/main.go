package main

import (
	"log"
	"net/http"

	"github.com/decoch/dojo4/kadai4/decoch/controller"
)

func main() {
	http.HandleFunc("/", controller.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
