package main

import (
	"net/http"
	"onefc/service"
	"os"

	"github.com/prometheus/common/log"
)

func main() {
	log.Info("Starting product service at port:8080")
	http.ListenAndServe(":8080", service.Router(os.Args[1]))
}
