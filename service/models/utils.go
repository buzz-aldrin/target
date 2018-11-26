package models

import (
	"target/service/dal"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func (prod *Product) FindProduct(dbName string, prodID string) (err error) {
	findQ := bson.M{"_id": prodID}
	if err = dal.FindOne(dbName, prodCollName, findQ, prod); err != nil {
		return errors.Wrap(err, "prod.FindProduct failed to fetch product")
	}

	prodDesc := new(ProductDesc)
	if err = dal.FindOne(dbName, prodDescCollName, findQ, prodDesc); err != nil {
		return errors.Wrap(err, "prod.FindProduct failed to fetch product description")
	}
	prod.Desc = prodDesc.Desc

	return
}

func (prod *Product) UpsertProduct(dbName string) (err error) {
	findQ := bson.M{"_id": prod.ID}
	if err = dal.UpsertOne(dbName, prodCollName, findQ, prod); err != nil {
		return errors.Wrap(err, "prod.UpsertProduct failed upsert product")
	}
	return
}

func (prod *Product) DeleteProduct(dbName string, prodID string) (err error) {
	findQ := bson.M{"_id": prodID}
	if err = dal.DeleteOne(dbName, prodCollName, findQ); err != nil {
		return errors.Wrap(err, "prod.DeleteProduct failed to delete product")
	}
	return
}
