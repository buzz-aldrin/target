package models

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

var (
	prodCollName     = "Product"
	prodDescCollName = "Desc"
)

type CurrencyCode string

type Value string

func (v *Value) Validate() (err error) {
	if v == nil {
		return errors.Errorf("missing product value")
	}

	fVal, err := strconv.ParseFloat(string(*v), 64)
	if err != nil {
		return errors.Errorf("invalid product value:%s, err:%v", *v, err)
	}

	tempVal := Value(fmt.Sprintf("%.2f", fVal))
	v = &tempVal

	return
}

type ProductID string

func (pid ProductID) Validate() (err error) {
	if len(pid) == 0 {
		return errors.Errorf("missing product id")
	}

	if _, err = strconv.Atoi(string(pid)); err != nil {
		return errors.Errorf("invalid product id:%s, err:%v", pid, err)
	}
	return
}

type CurrentPrice struct {
	Value        *Value       `json:"value" bson:"value"`
	CurrencyCode CurrencyCode `json:"currency_code" bson:"currency_code"`
}

type Product struct {
	ID           ProductID    `json:"id" bson:"_id"`
	CurrentPrice CurrentPrice `json:"current_price" bson:"current_price"`
	Desc         string       `json:"product_desc" bson:"product_desc"`
}

type ProductDesc struct {
	ID   ProductID `json:"id" bson:"_id"`
	Desc string    `json:"product_desc" bson:"product_desc"`
}
