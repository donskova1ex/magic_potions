package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/donskova1ex/magic_potions/internal"
	"github.com/donskova1ex/magic_potions/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *Repository) createIngredientsTx(tx *sqlx.Tx, ingredients []*domain.Ingredient) ([]*domain.Ingredient, error) {
	query := `
	INSERT INTO ingredients (name, uuid) 
	values (:name, :uuid) 
	on conflict on constraint ingredients_name_key 
	do update set name=excluded.name 
	returning id, name;
	`
	rows, err := tx.NamedQuery(query, ingredients)
	if err != nil {
		return nil, fmt.Errorf("failed to query ingredients: %w", err)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to reading rows: %w", rows.Err())
	}

	ingredientsMap := make(map[string]int32)
	var id int32
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("failed scan ingredient values: %w", err)
		}
		ingredientsMap[name] = id

	}

	savedIngredients := make([]*domain.Ingredient, 0, len(ingredients))
	for _, i := range ingredients {
		savedIngredients = append(savedIngredients, &domain.Ingredient{
			ID:       ingredientsMap[i.Name],
			UUID:     i.UUID,
			Name:     i.Name,
			Quantity: i.Quantity,
		})

	}

	return savedIngredients, nil
}

func (r *Repository) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {

	var entUUID string
	var pqErr *pq.Error

	query := `INSERT INTO ingredients (name, uuid) values ($1, $2) on conflict on constraint ingredients_name_key do nothing RETURNING uuid`

	newUUID := uuid.NewString()

	row := r.db.QueryRowContext(ctx, query, ingredient.Name, newUUID)
	err := row.Err()
	if err != nil {
		if errors.As(err, &pqErr) {
			if pqErr.Constraint == "ingredients_uuid_key" {
				return nil, fmt.Errorf("this uuid  is already in use: %w", err)
			}

			if pqErr.Constraint == "ingredients_name_key" {
				return nil, fmt.Errorf("this name is already in use: %w", err)
			}
		}
		return nil, fmt.Errorf("can not read ingredient from db: %w", err)
	}

	if err := row.Scan(&entUUID); err != nil {
		return nil, fmt.Errorf("impossible to create an entity: %w", err)
	}

	newIngredient := &domain.Ingredient{
		Name: ingredient.Name,
		UUID: newUUID,
	}
	return newIngredient, nil
}

func (r *Repository) IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error) {
	var ingredients []*domain.Ingredient
	err := r.db.SelectContext(ctx, &ingredients, "SELECT uuid, name FROM ingredients")
	if errors.Is(err, sql.ErrNoRows) {
		return ingredients, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", internal.ErrReadRows)
	}
	return ingredients, nil

}

func (r *Repository) IngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{}
	query := "SELECT name, uuid FROM ingredients WHERE uuid = $1"
	err := r.db.GetContext(ctx, ingredient, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}

	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}
	return ingredient, nil
}

func (r *Repository) DeleteIngredientByUUID(ctx context.Context, uuid string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM ingredients WHERE uuid = $1", uuid)
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrGetByUUID, uuid)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}

	if rows == 0 {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotDelete, uuid)
	}
	return nil
}

func (r *Repository) UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	query := "UPDATE ingredients SET name = $1 WHERE uuid = $2"
	_, err := r.db.ExecContext(ctx, query, ingredient.Name, ingredient.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return ingredient, nil

}
