package service

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/common/log"
)

// write data to response writer
func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	bs, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return
	}
	if _, err := w.Write(bs); err != nil {
		log.Error(err)
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

// returns internal server error
func writeErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	er := errorResponse{Message: "unable to process request"}
	bs, err := json.Marshal(er)
	if err != nil {
		log.Error(err)
		return
	}
	if _, err := w.Write(bs); err != nil {
		log.Error(err)
	}
}
