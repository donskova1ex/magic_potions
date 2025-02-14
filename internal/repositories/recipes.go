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

// TODO: проверить транзакции, если норм, то сделать аналогично во все запросы
func (r *Repository) CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {
	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("error start transaction: %w", internal.ErrRecipeTransaction)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			r.logger.Error("error rollbacking transaction", slog.String("err", err.Error()))
			return
		}
	}()
	for _, ingredient := range recipe.Ingredients {
		ingredient.UUID = uuid.NewString()
	}

	ingredientsMap, err := r.createIngredientsTx(ctx, tx, recipe.Ingredients)
	if err != nil {
		return nil, fmt.Errorf("error creating ingredients from recipe: %w", err)
	}

	newRecipe, err := r.createRecipeTx(ctx, tx, recipe)
	if err != nil {
		return nil, fmt.Errorf("error creating new recipe: %w", internal.ErrCreateRecipe)
	}

	ingredients, err := saveRecipesToIngredients(tx, newRecipe.ID, recipe.Ingredients, ingredientsMap)
	if err != nil {
		return nil, fmt.Errorf("error saving recipe ingredients: %w", err)
	}

	newRecipe.Ingredients = ingredients
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", internal.ErrRecipeTransaction)
	}

	return newRecipe, nil
}

func (r *Repository) createRecipeTx(ctx context.Context, tx *sqlx.Tx, recipe *domain.Recipe) (*domain.Recipe, error) {
	var id int32
	newUUID := uuid.NewString()
	query := `INSERT INTO recipes (uuid, Name, BrewTimeSeconds) values ($1, $2, $3) 
				on conflict on constraint recipes_name_key RETURNING id`
	row := tx.QueryRowxContext(ctx, query, newUUID, recipe.Name)

	if row.Err() != nil {
		return nil, fmt.Errorf("error query with name %s: %w", recipe.Name, row.Err())
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
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}
	if err != nil {
		return nil, fmt.Errorf("%w by uuid: [%s]", internal.ErrReadRows, uuid)
	}
	return recipe, nil
}

func (r *Repository) DeleteRecipeByUUID(ctx context.Context, uuid string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM recipes WHERE uuid = $1", uuid)

	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrGetByUUID, uuid)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}

	if rows != 1 {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotDelete, uuid)
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
	ingredientsMap map[string]int32) ([]*domain.Ingredient, error) {

	var createdIngredients []*domain.Ingredient

	for _, ingredient := range ingredients {
		ingredientID, exists := ingredientsMap[ingredient.Name]
		if !exists {
			return nil, fmt.Errorf("ingredient [%s] does not exist", ingredient.Name)
		}

		query := `INSERT into recipes_to_ingredients (recipe_id, ingredient_id, quantity) 
				values (:recipe_id, :ingredient_id, :quantity)
				on conflict on constraint do nothing`
		queryParams := map[string]interface{}{
			"recipe_id":     recipeId,
			"ingredient_id": ingredientID,
			"quantity":      ingredient.Quantity,
		}
		_, err := tx.NamedExec(query, queryParams)
		if err != nil {
			return nil, fmt.Errorf("failed to query recipe ingredients: %w", err)
		}
		createdIngredient := &domain.Ingredient{
			ID:       ingredientID,
			UUID:     ingredient.UUID,
			Name:     ingredient.Name,
			Quantity: ingredient.Quantity,
		}
		createdIngredients = append(createdIngredients, createdIngredient)
	}
	return createdIngredients, nil
}
