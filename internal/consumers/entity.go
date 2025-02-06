package consumers

type Recipe struct {
	Name            string        `json:"name"`
	BrewTimeSeconds int32         `json:"brew_time_seconds"`
	Ingredients     []*Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}
