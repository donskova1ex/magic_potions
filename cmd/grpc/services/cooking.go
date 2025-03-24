package services

import (
	"context"
	"fmt"
	"github.com/donskova1ex/magic_potions/cmd/grpc/generated"
	"log/slog"
	"sync"
	"time"

	"github.com/donskova1ex/magic_potions/internal/processors"
)

type Server struct {
	generated.UnimplementedBrewingServiceServer
	cookingStatus map[string]generated.GetCookingStatusResponse_Status
	mu            *sync.Mutex
	recipes       *processors.Recipes
	witches       *processors.Witches
	logger        *slog.Logger
}

func NewServer(
	cookingStatus map[string]generated.GetCookingStatusResponse_Status,
	mu *sync.Mutex,
	recipes *processors.Recipes,
	witches *processors.Witches,
	logger *slog.Logger,

) *Server {
	return &Server{
		cookingStatus: cookingStatus,
		mu:            mu,
		recipes:       recipes,
		witches:       witches,
		logger:        logger,
	}
}

func (s *Server) StartCooking(ctx context.Context, request *generated.StartCookingRequest) (*generated.StartCookingResponse, error) {
	recipeUUID := request.GetRecipeUuid()
	witchUUID := request.GetWitchUuid()
	brewTimeSeconds := request.GetBrewTimeSeconds()

	recipe, err := s.recipes.RecipeByUUID(ctx, recipeUUID)
	if err != nil {
		s.logger.Error("error getting recipe", slog.String("recipe_uuid", recipeUUID), slog.String("err", err.Error()))
		return nil, fmt.Errorf("could not get recipe by uuid: %w", err)
	}
	witch, err := s.witches.GetWitchByUUID(ctx, witchUUID)
	if err != nil {
		s.logger.Error("error getting witch", slog.String("witch_uuid", witchUUID), slog.String("err", err.Error()))
		return nil, fmt.Errorf("could not get witch by uuid: %w", err)
	}

	go func() {
		time.Sleep(time.Duration(brewTimeSeconds) * time.Second)
		s.mu.Lock()
		defer s.mu.Unlock()
		s.cookingStatus[recipeUUID] = generated.GetCookingStatusResponse_COMPLETED
	}()

	s.mu.Lock()
	s.cookingStatus[recipeUUID] = generated.GetCookingStatusResponse_IN_PROGRESS
	s.mu.Unlock()

	return &generated.StartCookingResponse{
		RecipeUuid: recipeUUID,
		RecipeName: recipe.Name,
		WitchUuid:  witch.UUID,
		WitchName:  witch.Name,
		Status:     "Started",
	}, nil
}
func (s *Server) GetCookingStatus(ctx context.Context, request *generated.GetCookingStatusRequest) (*generated.GetCookingStatusResponse, error) {
	recipeUUID := request.GetRecipeUuid()

	s.mu.Lock()
	status, exists := s.cookingStatus[recipeUUID]
	s.mu.Unlock()

	if !exists {
		s.logger.Error("error getting cooking status", slog.String("recipe_uuid", recipeUUID))
		return nil, fmt.Errorf("could not get cooking status for recipe uuid %s", recipeUUID)
	}

	return &generated.GetCookingStatusResponse{
		Status: status,
	}, nil
}
