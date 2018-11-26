package service

import (
	"net/http"

	"target/models"

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
	if err := prod.Find(dbName, prodID); err != nil {
		err = errors.Errorf("GET /product/%s failed, err:%v", prodID, err)
		log.Error(err)
		writeErrorResponse(w)
		return
	}
	prodDesc := new(models.ProductDesc)
	if err := prodDesc.Find(dbName, prodID); err != nil {
		err = errors.Errorf("GET /product/%s failed, err:%v", prodID, err)
		log.Error(err)
		writeErrorResponse(w)
		return
	}
	resp := ProductResp{Product: prod, Desc: prodDesc.Desc}

	writeSuccessResponse(w, resp)
}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	prod := new(models.Product)
	if err := json.NewDecoder(r.Body).Decode(prod); err != nil {
		err = errors.Errorf("PUT /product/%s failed to decode request body, err:%v", prod.ID, err)
		log.Error(err)
		writeErrorResponse(w)
		return
	}

	if err := prod.Upsert(dbName); err != nil {
		err = errors.Errorf("PUT /product/%s failed to update/create product:%+v, err:%v", prod.ID, *prod, err)
		log.Error(err)
		writeErrorResponse(w)
		return
	}

	writeSuccessResponse(w, prod)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prodID := vars["id"]

	prod := new(models.Product)
	if err := prod.Delete(dbName, prodID); err != nil {
		err = errors.Errorf("DELETE /product/%s failed to delete product, err:%v", prodID, err)
		log.Error(err)
		writeErrorResponse(w)
		return
	}
	writeSuccessResponse(w, prod)
}
