package domain

type Recipe struct {
	ID              int32  `db:"id"`
	UUID            string `json:"uuid" db:"uuid"`
	Name            string `json:"name" db:"name"`
	BrewTimeSeconds int32  `json:"brew_time_seconds" db:"brew_time_seconds"`
	Ingredients     []*Ingredient
}
