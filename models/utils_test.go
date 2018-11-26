package models

import (
	"testing"
)

func TestValue_Validate(t *testing.T) {
	var err error
	var v *Value

	if _, err = v.Validate(); err == nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			err, "models_Value.Validate missing product value")
	}

	temp := Value("invalid")
	v = &temp
	if _, err = v.Validate(); err == nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			err, "models_Value.Validate invalid product value")
	}

	temp = Value("99.999999")
	v = &temp
	if v, err = v.Validate(); err != nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			nil, err)
	}
	if *v != Value("100.00") {
		t.Errorf("returned wrong value: got %v expecting %v",
			*v, "100.00")
	}

	temp = Value("99.99")
	v = &temp
	if v, err = v.Validate(); err != nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			nil, err)
	}
	if *v != temp {
		t.Errorf("returned wrong value: got %v expecting %v",
			*v, temp)
	}
}

func TestProductID_Validate(t *testing.T) {
	var err error
	var prodID ProductID

	if err = prodID.Validate(); err == nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			err, "models_ProductID.Validate missing product id")
	}

	prodID = ProductID("invalid")
	if err = prodID.Validate(); err == nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			err, "models_ProductID.Validate invalid product id")
	}

	prodID = ProductID("1234")
	if err = prodID.Validate(); err != nil {
		t.Errorf("returned wrong value: got %v expecting %v",
			nil, err)
	}
}
