package models

var (
	// product collection name
	prodCollName = "Product"
	// product description collection name
	prodDescCollName = "ProductDesc"
)

// CurrencyCode type
type CurrencyCode string

// Value type
type Value string

// ProductID type
type ProductID string

// CurrentPrice type
type CurrentPrice struct {
	Value        *Value       `json:"value" bson:"value"`
	CurrencyCode CurrencyCode `json:"currency_code" bson:"currency_code"`
}

// Product type
type Product struct {
	ID           ProductID    `json:"id" bson:"_id"`
	CurrentPrice CurrentPrice `json:"current_price" bson:"current_price"`
}

// ProductDesc product description type
type ProductDesc struct {
	ID   ProductID `json:"id" bson:"_id"`
	Desc string    `json:"product_desc" bson:"product_desc"`
}
