package processors

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/donskova1ex/magic_potions/internal/domain"
)

//go:generate mockgen -destination=./mocks/witches_repository.go -package=mocks -mock_names=WitchesRepository=WitchesRepository . WitchesRepository
type WitchesRepository interface {
	CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error)
	WitchesAll(ctx context.Context) ([]*domain.Witch, error)
	WitchByUUID(ctx context.Context, uuid string) (*domain.Witch, error)
	DeleteWitchByUUID(ctx context.Context, uuid string) error
	UpdateWitchByUUID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error)
}

//go:generate mockgen -destination=./mocks/witches_logger.go -package=mocks -mock_names=WitchesLogger=WitchesLogger . WitchesLogger
type WitchesLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type Witches struct {
	witchesRepository WitchesRepository
	log               WitchesLogger
}

func NewWitch(witchesRepository WitchesRepository, log WitchesLogger) *Witches {
	return &Witches{witchesRepository: witchesRepository, log: log}
}

func (wtch *Witches) WitchesList(ctx context.Context) ([]*domain.Witch, error) {
	r, err := wtch.witchesRepository.WitchesAll(ctx)
	if err != nil {
		wtch.log.Error("it is impossible to get a Witches list",
			slog.String("err", err.Error()))
		return nil, fmt.Errorf("witches list getting error: %w", err)
	}
	return r, nil
}

func (wtch *Witches) GetWitchByUUID(ctx context.Context, uuid string) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.WitchByUUID(ctx, uuid)
	if err != nil {
		wtch.log.Error("unable to get witch by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return nil, fmt.Errorf("unable to get witch by uuid: %s, error: %w", uuid, err)
	}
	return w, nil
}
func (wtch *Witches) DeleteWitchByUUID(ctx context.Context, uuid string) error {
	err := wtch.witchesRepository.DeleteWitchByUUID(ctx, uuid)
	if err != nil {
		wtch.log.Error("unable to delete witch by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return fmt.Errorf("unable to delete witch by uuid: %s, error: %w", uuid, err)
	}
	return nil
}

func (wtch *Witches) UpdateWitchByUUID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.UpdateWitchByUUID(ctx, witch)
	if err != nil {
		wtch.log.Error("unable to update witch by uuid")
		return nil, fmt.Errorf("unable to update witch by uuid: %s, error: %w", witch.Name, err)
	}
	return w, nil
}

func (wtch *Witches) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.CreateWitch(ctx, witch)

	if err != nil {
		wtch.log.Error("unable to create witch",
			slog.String("err", err.Error()))
		return nil, fmt.Errorf("unable to create witch: %s, error: %w", witch.Name, err)
	}

	return w, nil
}
