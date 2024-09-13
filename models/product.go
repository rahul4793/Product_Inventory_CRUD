package models

type Product struct {
	ID       string  `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Quantity int     `json:"quantity" bson:"quantity"`
	Price    float64 `json:"price" bson:"price"`
}
