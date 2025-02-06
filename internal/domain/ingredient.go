package domain

type Ingredient struct {
	ID       int32  `json:"id" db:"id"`
	UUID     string `json:"uuid" db:"uuid"`
	Name     string `json:"name" db:"name"`
	Quantity int32  `json:"quantity" db:"quantity"`
}
