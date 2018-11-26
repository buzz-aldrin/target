package service

import "target/models"

type ProductResp struct {
	*models.Product
	Desc string `json:"product_desc"`
}
