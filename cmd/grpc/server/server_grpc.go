package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"sync"

	"github.com/donskova1ex/magic_potions/cmd/grpc/services"
	pb "github.com/donskova1ex/magic_potions/generated"
	"github.com/donskova1ex/magic_potions/internal/processors"
	"github.com/donskova1ex/magic_potions/internal/repositories"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJSONHandler)
	slog.SetDefault(logger)

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logger.Info("env file loaded")
	pgDSN := os.Getenv("POSTGRES_DSN")
	if pgDSN == "" {
		logger.Error("empty POSTGRES_DSN")
		os.Exit(1)
	}
	logger.Info("postgres DSN loaded")

	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort == "" {
		logger.Error("empty GRPC_PORT")
		os.Exit(1)
	}
	logger.Info("gRPC port loaded")

	lis, err := net.Listen("tcp", ":"+gRPCPort)
	if err != nil {
		logger.Error("failed to listen",
			slog.String("port", gRPCPort),
			slog.String("err", err.Error()),
		)
		os.Exit(1)
	}

	db, err := repositories.NewPostgresDB(ctx, pgDSN)
	if err != nil {
		logger.Error("can not create postgres db connection", slog.String("error", err.Error()))
		return
	}
	defer db.Close()

	repository := repositories.NewRepository(db, logger)
	ingredientProcessor := processors.NewIngredient(repository, logger)
	recipeProcessor := processors.NewRecipe(repository, logger)
	witchProcessor := processors.NewWitch(repository, logger)

	s := grpc.NewServer()

	cookingStatus := make(map[string]pb.GetCookingStatusResponse_Status)
	mu := &sync.Mutex{}

	newServer := services.NewServer(cookingStatus, mu, ingredientProcessor, recipeProcessor, witchProcessor, logger)

	pb.RegisterBrewingServiceServer(s, newServer)

	reflection.Register(s)

	logger.Info("gRPC starting and listening port " + gRPCPort)
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve", err)
	}

}
