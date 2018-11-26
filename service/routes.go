package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
The goal for this exercise is to create an end-to-end Proof-of-Concept for a products API,
which will aggregate product data from multiple sources and return it as JSON to the
caller.
Your goal is to create a RESTful service that can retrieve product and price details by ID.
The URL structure is up to you to define, but try to follow some sort of logical
convention.

1. Write a rest API to do a CURD operations and store the data id, current_price, currency_code
2. Accepts an HTTP PUT request at the same path (/products/{id}), containing a JSON request body similar to the
	GET response, and updates the product’s price in the data store.
	Example: {"id":13860428,"current_price":{"value":13.49,"currency_code":"USD"}}
3. Rest API - HTTP GET request at /products/{id} and delivers product data as JSON (where {id} will be a number.
	Example product IDs: 15117729, 16483589, 16696652, 16752456, 15643793)
	Example response: {"id":13860428,"current_price":{"value":13.49,"currency_code":"USD"}}
4. Assume there is a mongo DB which has product description and reads product information from a NoSQL data store
	and combines it with the product id and name from the HTTP request into a single response.
	Example: {"id":13860428,"current_price":{"value":13.49,"currency_code":"USD", "product_desc":"Shirt"}}
	The service should be secured.
*/

// Router adds service routes
func Router() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/product/{id}", GetProductHandler).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", PutProductHandler).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", DeleteProductHandler).Methods(http.MethodDelete)
	return
}
