package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/donskova1ex/magic_potions/internal"
	"github.com/donskova1ex/magic_potions/internal/domain"
	"github.com/google/uuid"
)

func (r *Repository) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	var id uint32

	query := "INSERT INTO witches (uuid, name) values ($1, $2) on conflict on constraint witches_name_key RETURNING uuid"

	newUUID := uuid.NewString()
	row := r.db.QueryRowContext(ctx, query, newUUID, witch.Name)
	err := row.Err()
	if err != nil {
		return nil, fmt.Errorf("can not read witch from db: %w", err)
	}
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("can not scan witch for id: %w", err)
	}
	newWitch := &domain.Witch{
		Name: witch.Name,
		UUID: newUUID,
	}
	return newWitch, nil
}

func (r *Repository) WitchesAll(ctx context.Context) ([]*domain.Witch, error) {
	witches := []*domain.Witch{}
	err := r.db.SelectContext(ctx, &witches, "SELECT uuid, name FROM witches")
	if errors.Is(err, sql.ErrNoRows) {
		return witches, nil
	}

	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", internal.ErrReadRows)
	}
	return witches, nil
}

func (r *Repository) WitchByUUID(ctx context.Context, uuid string) (*domain.Witch, error) {
	witch := &domain.Witch{}
	query := "SELECT name, uuid FROM witches WHERE uuid = $1"
	err := r.db.GetContext(ctx, witch, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}

	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}

	return witch, nil
}

func (r *Repository) DeleteWitchByUUID(ctx context.Context, uuid string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM witches WHERE uuid = $1", uuid)
	if err != nil {
		return fmt.Errorf("there is no object with this ID: %w", err)
	}
	return nil
}

func (r *Repository) UpdateWitchByUUID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	query := "UPDATE witches SET name = $1 WHERE uuid = $2"
	_, err := r.db.ExecContext(ctx, query, witch.Name, witch.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err) // TODO: Уточнить вопрос,пример в ingredients
	}
	return witch, nil
}
