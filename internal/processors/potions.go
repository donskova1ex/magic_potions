package processors

import (
	"context"
	"fmt"
	"log/slog"
)

type PotionsRepository interface {
	CreatePotion(ctx context.Context, potionUUID, witchUUID, recipeUUID, status string) error
	UpdatePotion(ctx context.Context, potionUUID, status string) error
}

type PotionsLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type Potions struct {
	potionsRepository PotionsRepository
	logger            PotionsLogger
}

func NewPotions(potionsRepository PotionsRepository, log PotionsLogger) *Potions {
	return &Potions{
		potionsRepository: potionsRepository,
		logger:            log,
	}
}

func (p *Potions) CreatePotion(ctx context.Context, potionUUID, witchUUID, recipeUUID, status string) error {
	err := p.potionsRepository.CreatePotion(ctx, potionUUID, witchUUID, recipeUUID, status)
	if err != nil {
		p.logger.Error("error creating potion", slog.String("err", err.Error()))
		return fmt.Errorf("error creating potion: %w", err)
	}
	return nil
}

func (p *Potions) UpdatePotion(ctx context.Context, potionUUID, status string) error {
	err := p.potionsRepository.UpdatePotion(ctx, potionUUID, status)
	if err != nil {
		p.logger.Error("error updating potion", slog.String("err", err.Error()))
		return fmt.Errorf("error updating potion: %w", err)
	}
	return nil

}
