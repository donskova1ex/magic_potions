package processors

import (
	"context"
	"fmt"
	"github.com/donskova1ex/magic_potions/internal/domain"
	"log/slog"
)

type IngredientsRepository interface {
	CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error)
	IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error)
	IngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error)
	DeleteIngredientByUUID(ctx context.Context, uuid string) error
	UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error)
}

type IngredientsLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type ingredients struct {
	ingredientsRepository IngredientsRepository
	log                   IngredientsLogger
}

func NewIngredient(ingredientsRepository IngredientsRepository, log IngredientsLogger) *ingredients {
	return &ingredients{ingredientsRepository: ingredientsRepository, log: log}
}

func (i *ingredients) IngredientsList(ctx context.Context) ([]*domain.Ingredient, error) {

	r, err := i.ingredientsRepository.IngredientsAll(ctx)
	if err != nil {
		i.log.Error("it is impossible to get a ingredients list", slog.String("err", err.Error()))
		return nil, fmt.Errorf("it is impossible to get a ingredients list: %w", err)
	}

	return r, nil
}

func (i *ingredients) GetIngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ing, err := i.ingredientsRepository.IngredientByUUID(ctx, uuid)
	if err != nil {
		i.log.Error("unable to get ingredient by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return nil, fmt.Errorf("unable to get ingredient by uuid: %s, error: %w", uuid, err)
	}
	return ing, nil
}
func (i *ingredients) DeleteIngredientByUUID(ctx context.Context, uuid string) error {
	err := i.ingredientsRepository.DeleteIngredientByUUID(ctx, uuid)
	if err != nil {
		i.log.Error("unable to delete ingredient by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return fmt.Errorf("unable to delete ingredient by uuid: %s, error: %w", uuid, err)
	}
	return nil
}

func (i *ingredients) UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	ing, err := i.ingredientsRepository.UpdateIngredientByUUID(ctx, ingredient)
	if err != nil {
		i.log.Error("unable to update ingredient by uuid")
		return nil, fmt.Errorf("unable to update ingredient by uuid: %s, error: %w", ingredient.UUID, err)
	}
	return ing, nil
}

func (i *ingredients) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	ing, err := i.ingredientsRepository.CreateIngredient(ctx, ingredient)

	if err != nil {
		i.log.Error("unable to create ingredient",
			slog.String("err", err.Error()))
		return nil, fmt.Errorf("unable to create ingredient: %s, error: %w", ingredient.Name, err)
	}

	return ing, nil
}
