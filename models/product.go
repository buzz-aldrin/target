package models

var (
	prodCollName     = "Product"
	prodDescCollName = "ProductDesc"
)

type CurrencyCode string

type Value string

type ProductID string

type CurrentPrice struct {
	Value        *Value       `json:"value" bson:"value"`
	CurrencyCode CurrencyCode `json:"currency_code" bson:"currency_code"`
}

type Product struct {
	ID           ProductID    `json:"id" bson:"_id"`
	CurrentPrice CurrentPrice `json:"current_price" bson:"current_price"`
}

type ProductDesc struct {
	ID   ProductID `json:"id" bson:"_id"`
	Desc string    `json:"product_desc" bson:"product_desc"`
}
