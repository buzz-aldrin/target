package models

import (
	"fmt"
	"strconv"
	"target/dal"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func (v *Value) Validate() (formatted *Value, err error) {
	if v == nil {
		return nil, errors.Errorf("models_Value.Validate missing product value")
	}

	fVal, err := strconv.ParseFloat(string(*v), 64)
	if err != nil {
		return nil, errors.Errorf("models_Value.Validate invalid product value:%s, err:%v", *v, err)
	}

	tempVal := Value(fmt.Sprintf("%.2f", fVal))

	return &tempVal, err
}

func (pid ProductID) Validate() (err error) {
	if len(pid) == 0 {
		return errors.Errorf("models_ProductID.Validate missing product id")
	}

	if _, err = strconv.Atoi(string(pid)); err != nil {
		return errors.Errorf("models_ProductID.Validate invalid product id:%s, err:%v", pid, err)
	}
	return
}

func (prod *Product) Find(dbName string, prodID string) (err error) {
	findQ := bson.M{"_id": prodID}
	if err = dal.FindOne(dbName, prodCollName, findQ, prod); err != nil {
		return errors.Wrap(err, "prod.Find failed to fetch product")
	}
	return
}

func (prodDesc *ProductDesc) Find(dbName string, prodID string) (err error) {
	findQ := bson.M{"_id": prodID}
	if err = dal.FindOne(dbName, prodDescCollName, findQ, prodDesc); err != nil {
		return errors.Wrap(err, "prod.Find failed to fetch product description")
	}
	return
}

func (prod *Product) Upsert(dbName string) (err error) {
	findQ := bson.M{"_id": prod.ID}

	if err = prod.ID.Validate(); err != nil {
		return errors.Wrap(err, "prod.Upsert invalid product.ID")
	}

	if prod.CurrentPrice.Value, err = prod.CurrentPrice.Value.Validate(); err != nil {
		return errors.Wrap(err, "prod.Upsert invalid product.CurrentPrice.Value")
	}

	if err = dal.UpsertOne(dbName, prodCollName, findQ, prod); err != nil {
		return errors.Wrap(err, "prod.Upsert failed upsert product")
	}
	return
}

func (prod *Product) Delete(dbName string, prodID string) (err error) {
	findQ := bson.M{"_id": prodID}
	if err = dal.DeleteOne(dbName, prodCollName, findQ); err != nil {
		return errors.Wrap(err, "prod.Delete failed to delete product")
	}
	return
}
