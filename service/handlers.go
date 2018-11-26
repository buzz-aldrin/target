package service

import (
	"net/http"

	"target/service/models"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

var dbName = "target"

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prodID := vars["id"]

	prod := new(models.Product)
	if err := prod.FindProduct(dbName, prodID); err != nil {
		err = errors.Errorf("GET /product/%s failed, err:%v", prodID, err)
		log.Error(err)
		writeErrorResponse(w)
	}

	writeSuccessResponse(w, prod)
}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	prod := new(models.Product)
	if err := json.NewDecoder(r.Body).Decode(prod); err != nil {
		err = errors.Errorf("PUT /product/%s failed to decode request body, err:%v", prod.ID, err)
		log.Error(err)
		writeErrorResponse(w)
	}

	if err := prod.UpsertProduct(dbName); err != nil {
		err = errors.Errorf("PUT /product/%s failed to update/create product:%+v, err:%v", prod.ID, *prod, err)
		log.Error(err)
		writeErrorResponse(w)
	}

	writeSuccessResponse(w, prod)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prodID := vars["id"]

	prod := new(models.Product)
	if err := prod.DeleteProduct(dbName, prodID); err != nil {
		err = errors.Errorf("DELETE /product/%s failed to delete product, err:%v", prodID, err)
		log.Error(err)
		writeErrorResponse(w)
	}
	writeSuccessResponse(w, prod)
}
