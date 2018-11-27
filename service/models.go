package service

import "target/models"

type productResp struct {
	*models.Product
	Desc string `json:"product_desc"`
}
