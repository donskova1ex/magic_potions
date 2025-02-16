package processors

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/donskova1ex/magic_potions/internal/consumers"
	"github.com/donskova1ex/magic_potions/internal/domain"
)

type RecipesRepository interface {
	RecipesAll(ctx context.Context) ([]*domain.Recipe, error)
	CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error)
	RecipeByUUID(ctx context.Context, uuid string) (*domain.Recipe, error)
	DeleteRecipeByUUID(ctx context.Context, uuid string) error
	UpdateRecipeByUUID(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error)
	CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error)
}

type RecipesLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type recipes struct {
	recipesRepository RecipesRepository
	log               RecipesLogger
}

func NewRecipe(recipesRepository RecipesRepository, log RecipesLogger) *recipes {
	return &recipes{recipesRepository: recipesRepository, log: log}
}

// TODO: тестирование не забывать делать после методов
func (rec *recipes) RecipesList(ctx context.Context) ([]*domain.Recipe, error) {
	r, err := rec.recipesRepository.RecipesAll(ctx)
	if err != nil {
		rec.log.Error("it is impossible to get a recipes list",
			slog.String("err", err.Error()))
		return nil, fmt.Errorf("recipes list getting error: %w", err)
	}
	return r, nil
}

func (rec *recipes) CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {
	r, err := rec.recipesRepository.CreateRecipe(ctx, recipe)
	if err != nil {
		rec.log.Error("unable to create recipe",
			slog.String("err", err.Error()))
		return nil, fmt.Errorf("can not create recipe: %s, error: %w", recipe.Name, err)
	}
	return r, nil
}
func (rec *recipes) RecipeByUUID(ctx context.Context, uuid string) (*domain.Recipe, error) {
	r, err := rec.recipesRepository.RecipeByUUID(ctx, uuid)
	if err != nil {
		rec.log.Error("unable to get recipe by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return nil, fmt.Errorf("can not get recipe by uuid: %s, error: %w", uuid, err)
	}
	return r, nil
}

func (rec *recipes) DeleteRecipeByUUID(ctx context.Context, uuid string) error {
	err := rec.recipesRepository.DeleteRecipeByUUID(ctx, uuid)
	if err != nil {
		rec.log.Error("unable to delete witch by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return fmt.Errorf("unable to delete witch by uuid: %s, error: %w", uuid, err)
	}
	return nil
}

func (rec *recipes) UpdateRecipeByID(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {

	r, err := rec.recipesRepository.UpdateRecipeByUUID(ctx, recipe)
	if err != nil {
		rec.log.Error("unable to update recipe",
			slog.String("err", err.Error()))
		return nil, fmt.Errorf("can not update recipe: %s, error: %w", recipe.Name, err)
	}
	return r, nil
}

// TODO вынести в отдельную функцию конверт анмаршла в домаин структуры
func (rec *recipes) Save(ctx context.Context, key []byte, body []byte, timeStamp time.Time) error {
	//TODO: валидация структур
	recipe := &consumers.Recipe{}
	if err := json.Unmarshal(body, recipe); err != nil {
		rec.log.Error(
			"can not unmarshal recipe",
			slog.String("err", err.Error()),
			slog.String("value", string(body)),
		)
		return nil
	}
	domainRecipe := &domain.Recipe{
		Name:            recipe.Name,
		BrewTimeSeconds: recipe.BrewTimeSeconds,
	}
	domainIngredients := []*domain.Ingredient{}
	for _, ingredient := range recipe.Ingredients {
		domainIngredients = append(domainIngredients, &domain.Ingredient{
			Name:     ingredient.Name,
			Quantity: ingredient.Quantity,
		})
	}
	domainRecipe.Ingredients = domainIngredients

	/*TODO: creating retry (повторение сохранения 3-5 раз, если сообщение так и не вычиталось, писать в логи).
	RetryStruct и методы по ретраю*/
	_, err := rec.CreateRecipe(ctx, domainRecipe)
	if err != nil {
		return fmt.Errorf("failed to greate recipe: %w", err)
	}

	return nil
}
