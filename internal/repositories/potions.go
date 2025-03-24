package repositories

import (
	"context"
	"fmt"
	"log"
)

func (r *Repository) CreatePotion(ctx context.Context, potionUUID, witchUUID, recipeUUID, status string) error {
	var txCommited bool

	query := `INSERT INTO potions (uuid, witch_uuid, recipe_uuid, status, created_at, updated_at)
        			VALUES ($1, $2, $3, $4, NOW(), NOW())`

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	defer func() {
		if !txCommited {
			if err := tx.Rollback(); err != nil {
				log.Printf("error rolling back transaction: %v", err)
				return
			}
		}
	}()

	_, err = tx.Exec(query, potionUUID, witchUUID, recipeUUID, status)
	if err != nil {
		return fmt.Errorf("error inserting potion: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	txCommited = true

	return nil
}

func (r *Repository) UpdatePotion(ctx context.Context, potionUUID, status string) error {
	var txCommited bool

	query := `UPDATE potions SET status = $1, updated_at = NOW() WHERE uuid = $2`
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	defer func() {
		if !txCommited {
			if err := tx.Rollback(); err != nil {
				log.Printf("error rolling back transaction: %v", err)
				return
			}
		}
	}()

	_, err = tx.Exec(query, status, potionUUID)
	if err != nil {
		return fmt.Errorf("error updating potion: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	txCommited = true

	return nil
}
