package domain

import "time"

type Potion struct {
	UUID       string    `json:"uuid" db:"uuid"`
	WitchUuid  int32     `json:"witch_uuid" db:"witch_uuid"`
	RecipeUuid int32     `json:"recipe_uuid" db:"recipe_uuid"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
