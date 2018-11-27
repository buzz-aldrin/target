package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"target/dal"
	"target/models"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var (
	prodID1 = models.ProductID("15117729")
	//prodID2 = models.ProductID("16483589")
	prodID3 = models.ProductID("16696652")

	value1 = models.Value("13.49")
	//value2 = models.Value("99.99")
	value3 = models.Value("999.99")

	currencyCode1 = models.CurrencyCode("USD")
	//currencyCode2 = models.CurrencyCode("INR")
	currencyCode3 = models.CurrencyCode("EUR")

	prod1 = models.Product{
		ID: prodID1,
		CurrentPrice: models.CurrentPrice{
			Value:        &value1,
			CurrencyCode: currencyCode1,
		},
	}

	//prod2 = models.Product{
	//	ID: prodID2,
	//	CurrentPrice: models.CurrentPrice{
	//		Value:        &value2,
	//		CurrencyCode: currencyCode2,
	//	},
	//}

	prod3 = models.Product{
		ID: prodID3,
		CurrentPrice: models.CurrentPrice{
			Value:        &value3,
			CurrencyCode: currencyCode3,
		},
	}

	prodDesc1 = models.ProductDesc{
		ID:   prodID1,
		Desc: "Shirt",
	}

	//prodDesc2 = models.ProductDesc{
	//	ID:   prodID2,
	//	Desc: "Pants",
	//}

	//prodDesc3 = models.ProductDesc{
	//	ID:   prodID3,
	//	Desc: "Shoes",
	//}
)

var (
	invalidValue1 = models.Value("invalid")

	invalidProd2 = models.Product{
		ID: "invalid",
		CurrentPrice: models.CurrentPrice{
			Value:        &value3,
			CurrencyCode: currencyCode3,
		},
	}

	invalidProd3 = models.Product{
		ID: prodID1,
		CurrentPrice: models.CurrentPrice{
			Value:        &invalidValue1,
			CurrencyCode: currencyCode3,
		},
	}
)

//func dataSetup(t *testing.T) {
//	if err := dal.CreateOne("target", "Product", prod1); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.CreateOne("target", "Product", prod2); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.CreateOne("target", "Product", prod3); err != nil {
//		t.Fatal(err)
//	}
//
//	if err := dal.CreateOne("target", "ProductDesc", prodDesc1); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.CreateOne("target", "ProductDesc", prodDesc2); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.CreateOne("target", "ProductDesc", prodDesc3); err != nil {
//		t.Fatal(err)
//	}
//}

//func dataCleanup(t *testing.T) {
//	if err := dal.DeleteOne("target", "Product", bson.M{"_id": prodID1}); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.DeleteOne("target", "Product", bson.M{"_id": prodID2}); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.DeleteOne("target", "Product", bson.M{"_id": prodID3}); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.DeleteOne("target", "ProductDesc", bson.M{"_id": prodID1}); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.DeleteOne("target", "ProductDesc", bson.M{"_id": prodID2}); err != nil {
//		t.Fatal(err)
//	}
//	if err := dal.DeleteOne("target", "ProductDesc", bson.M{"_id": prodID3}); err != nil {
//		t.Fatal(err)
//	}
//}

func TestGetProductHandler1(t *testing.T) {
	req, err := http.NewRequest("GET", "/product/"+string(prodID1), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", GetProductHandler).Methods(http.MethodGet)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestGetProductHandler2(t *testing.T) {
	if err := dal.CreateOne("target", "Product", prod1); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/product/"+string(prodID1), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", GetProductHandler).Methods(http.MethodGet)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	if err := dal.DeleteOne("target", "Product", bson.M{"_id": prodID1}); err != nil {
		t.Fatal(err)
	}
}

func TestGetProductHandler3(t *testing.T) {
	if err := dal.CreateOne("target", "Product", prod1); err != nil {
		t.Fatal(err)
	}
	if err := dal.CreateOne("target", "ProductDesc", prodDesc1); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/product/"+string(prodID1), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", GetProductHandler).Methods(http.MethodGet)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	resp := new(productResp)
	if err = json.NewDecoder(rr.Body).Decode(resp); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*resp.Product, prod1) {
		t.Errorf("handler returned wrong response: got %v want %v",
			*resp.Product, prod1)
	}

	if err := dal.DeleteOne("target", "Product", bson.M{"_id": prodID1}); err != nil {
		t.Fatal(err)
	}
	if err := dal.DeleteOne("target", "ProductDesc", bson.M{"_id": prodID1}); err != nil {
		t.Fatal(err)
	}
}

func TestPutProductHandler1(t *testing.T) {
	bc := make([]byte, 0)

	req, err := http.NewRequest("PUT", "/product/invalid", bytes.NewReader(bc))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", PutProductHandler).Methods(http.MethodPut)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestPutProductHandler2(t *testing.T) {
	bc, err := json.Marshal(invalidProd2)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/product/"+string(invalidProd2.ID), bytes.NewReader(bc))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", PutProductHandler).Methods(http.MethodPut)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestPutProductHandler3(t *testing.T) {
	bc, err := json.Marshal(invalidProd3)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/product/"+string(invalidProd3.ID), bytes.NewReader(bc))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", PutProductHandler).Methods(http.MethodPut)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestPutProductHandler4(t *testing.T) {
	bc, err := json.Marshal(prod3)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/product/"+string(prod3.ID), bytes.NewReader(bc))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", PutProductHandler).Methods(http.MethodPut)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	prod3.CurrentPrice.Value = &value1
	bc, err = json.Marshal(prod3)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("PUT", "/product/"+string(prod3.ID), bytes.NewReader(bc))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router = mux.NewRouter()
	router.HandleFunc("/product/{id}", PutProductHandler).Methods(http.MethodPut)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if err := dal.DeleteOne("target", "Product", bson.M{"_id": prodID3}); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteProductHandler1(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/product/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", DeleteProductHandler).Methods(http.MethodDelete)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestDeleteProductHandler2(t *testing.T) {
	if err := dal.CreateOne("target", "Product", prod1); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/product/"+string(prod1.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", DeleteProductHandler).Methods(http.MethodDelete)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
