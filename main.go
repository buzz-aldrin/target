package main

import (
	"net/http"
	"target/service"

	"github.com/prometheus/common/log"
)

func main() {
	log.Info("Starting product service at port:8080")
	if err := http.ListenAndServeTLS(":8080", "service/keys/server.crt", "service/keys/server.key",
		service.Router()); err != nil {
		log.Fatal(err)
	}
}
