package models

type CurrencyCode string

type CurrentPrice struct {
	Value        string       `json:"value" bson:"value"`
	CurrencyCode CurrencyCode `json:"currency_code" bson:"currency_code"`
}

type Product struct {
	ID           string       `json:"id" bson:"_id"`
	CurrentPrice CurrentPrice `json:"current_price" bson:"current_price"`
	ProductDesc  string       `json:"product_desc" bson:"product_desc"`
}
