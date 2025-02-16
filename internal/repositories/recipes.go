package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/donskova1ex/magic_potions/internal"
	"github.com/donskova1ex/magic_potions/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *Repository) CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {
	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("error start transaction: %w", internal.ErrRecipeTransaction)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			r.logger.Error("error roll back transaction", slog.String("err", err.Error()))
			return
		}
	}()
	for _, ingredient := range recipe.Ingredients {
		ingredient.UUID = uuid.NewString()
	}

	savedIngredients, err := r.createIngredientsTx(tx, recipe.Ingredients)
	if err != nil {
		return nil, fmt.Errorf("error creating ingredients from recipe: %w", err)
	}

	newRecipe, err := r.createRecipeTx(ctx, tx, recipe)
	if err != nil {
		return nil, fmt.Errorf("error creating new recipe: %w", err)
	}

	if err := saveRecipesToIngredients(tx, newRecipe.ID, savedIngredients); err != nil {
		return nil, fmt.Errorf("error saving recipe ingredients: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	newRecipe.Ingredients = savedIngredients

	return newRecipe, nil
}

func (r *Repository) createRecipeTx(ctx context.Context, tx *sqlx.Tx, recipe *domain.Recipe) (*domain.Recipe, error) {
	var id int32
	newUUID := uuid.NewString()
	//TODO: Upper case in query
	query := `INSERT INTO recipes (uuid, name, brew_time_seconds) VALUES ($1, $2, $3) 
				ON CONFLICT ON CONSTRAINT recipes_name_key DO NOTHING RETURNING id`

	row := tx.QueryRowxContext(ctx, query, newUUID, recipe.Name, recipe.BrewTimeSeconds)
	if row.Err() != nil {
		return nil, fmt.Errorf("error query with name [%s]: %w", recipe.Name, row.Err())
	}

	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("error scanning id from row: %w", err)
	}

	newRecipe := &domain.Recipe{
		ID:              id,
		UUID:            newUUID,
		Name:            recipe.Name,
		BrewTimeSeconds: recipe.BrewTimeSeconds,
		Ingredients:     nil,
	}

	return newRecipe, nil

}

func (r *Repository) RecipesAll(ctx context.Context) ([]*domain.Recipe, error) {
	var recipes []*domain.Recipe

	err := r.db.SelectContext(ctx, &recipes, "SELECT uuid, name, brew_time_seconds FROM recipes")
	if errors.Is(err, sql.ErrNoRows) {
		return recipes, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", internal.ErrReadRows)
	}
	return recipes, nil
}

func (r *Repository) RecipeByUUID(ctx context.Context, uuid string) (*domain.Recipe, error) {
	recipe := &domain.Recipe{}
	query := "SELECT uuid, name, brew_time_seconds FROM recipes WHERE uuid = $1"
	err := r.db.GetContext(ctx, recipe, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("err getting recipe with uuid [%s]: %w", uuid, err)
	}
	if err != nil {
		return nil, fmt.Errorf("err getting recipe with uuid [%s]: %w", uuid, err)
	}
	return recipe, nil
}

func (r *Repository) DeleteRecipeByUUID(ctx context.Context, uuid string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM recipes WHERE uuid = $1", uuid)
	if err != nil {
		return fmt.Errorf("error delete recipe with uuid [%s]: %w", uuid, err)
	}

	return nil
}

func (r *Repository) UpdateRecipeByUUID(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {

	query := "UPDATE recipes SET name = $1, brew_time_seconds = $2 WHERE uuid = $3"
	_, err := r.db.ExecContext(ctx, query, recipe.Name, recipe.BrewTimeSeconds, recipe.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return recipe, nil
}

func saveRecipesToIngredients(
	tx *sqlx.Tx,
	recipeId int32,
	ingredients []*domain.Ingredient,
) error {

	var queryParameters []map[string]interface{}

	for _, ingredient := range ingredients {
		queryParameters = append(queryParameters, map[string]interface{}{
			"recipe_id":     recipeId,
			"ingredient_id": ingredient.ID,
			"quantity":      ingredient.Quantity,
		})
	}
	deleteQuery := `
	DELETE FROM ingredients 
	WHERE 
		recipe_id=$1 AND 
		ingredient_id <> ANY(:ingredient_id)`

	_, err := tx.NamedExec(deleteQuery, queryParameters)
	if err != nil {
		return fmt.Errorf("failed to delete old recipe ingredients: %w", err)
	}

	query := `
	INSERT INTO recipes_to_ingredients rti (recipe_id, ingredient_id, quantity) 
	VALUES (:recipe_id, :ingredient_id, :quantity)
	ON CONFLICT ON CONSTRAINT recipes_to_ingredients_pkey 
	DO UPDATE SET 
		rti.quantity = EXCLUDED.quantity`

	_, err = tx.NamedExec(query, queryParameters)
	if err != nil {
		return fmt.Errorf("failed to query recipe ingredients: %w", err)
	}

	return nil
}
