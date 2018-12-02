package main

import (
	"log"
	"net/http"

	"github.com/decoch/dojo4/kadai4/decoch/application"
	"github.com/decoch/dojo4/kadai4/decoch/handler"
)

func main() {
	omikujiHandler := handler.OmikujiHandler{
		Service: application.NewOmikujiService(),
	}
	http.Handle("/", &omikujiHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
