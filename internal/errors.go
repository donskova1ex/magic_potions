package internal

import (
	"errors"
)

var ErrNotDelete = errors.New("ent not delete")
var ErrNotFound = errors.New("entity not found")
var ErrGetByUUID = errors.New("entity not found")
var ErrReadRows = errors.New("can not read rows")
var ErrCreateRecipe = errors.New("recipe creation error")
var ErrCreateConsumerGroup = errors.New("consumer group creation error")
var ErrClosingCosumerGroupClient = errors.New("closing cosumer group client error")
var ErrRecipeTransaction = errors.New("recipe transaction creation error")
var ErrRecipeRollback = errors.New("recipe transaction creation error")
