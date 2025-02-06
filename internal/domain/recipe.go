package domain

type Recipe struct {
	UUID            string `json:"uuid" db:"uuid"`
	Name            string `json:"name" db:"name"`
	BrewTimeSeconds int32  `json:"brew_time_seconds" db:"brew_time_seconds"`
	Ingredients     []*Ingredient
}

//TODO: пример итерирования ингредиентов
// {
// 	"name":"recipe1",
// 	"brew_time_second":15,
// 	"ingredients":
// 	["asdad", "asdads"]
// }

// for ing := range ricipe.ingredients {

// }
