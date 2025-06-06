package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/donskova1ex/magic_potions/internal"
	"github.com/donskova1ex/magic_potions/internal/metrics"
	"github.com/donskova1ex/magic_potions/internal/middleware"
	"github.com/donskova1ex/magic_potions/internal/processors"
	"github.com/donskova1ex/magic_potions/internal/repositories"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	openapi "github.com/donskova1ex/magic_potions/openapi"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("Server started")

	logJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJSONHandler)

	slog.SetDefault(logger)

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pgDSN := os.Getenv("POSTGRES_DSN")
	if pgDSN == "" {
		logger.Error("empty POSTGRES_DSN")
		os.Exit(1)
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		logger.Error("empty API_PORT")
		os.Exit(1)
	}

	db, err := repositories.NewPostgresDB(ctx, pgDSN)
	if err != nil {
		logger.Error("can not create postgres db connection", slog.String("error", err.Error()))
		return
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			logger.Error("can not close postgres db connection", slog.String("error", err.Error()))
		}
	}(db)
	rdb, err := repositories.NewRedisDB(ctx)
	if err != nil {
		logger.Error("can not create redis db connection", slog.String("error", err.Error()))
	}
	repository := repositories.NewRepository(db, rdb, logger)

	ingredientProcessor := processors.NewIngredient(repository, logger)
	IngredientAPIService := openapi.NewIngredientAPIService(ingredientProcessor, logger)
	IngredientAPIController := openapi.NewIngredientAPIController(IngredientAPIService)

	recipeProcessor := processors.NewRecipe(repository, logger)
	RecipeAPIService := openapi.NewRecipeAPIService(recipeProcessor, logger)
	RecipeAPIController := openapi.NewRecipeAPIController(RecipeAPIService)

	witchProcessor := processors.NewWitch(repository, logger)
	WitchAPIService := openapi.NewWitchAPIService(witchProcessor, logger)
	WitchAPIController := openapi.NewWitchAPIController(WitchAPIService)

	router := openapi.NewRouter(IngredientAPIController, RecipeAPIController, WitchAPIController)
	requestLogger := middleware.RequestLogger(logger)

	metricsNew := metrics.NewMetrics()
	metricsMiddleware := middleware.MetricsMiddleware(metricsNew)

	router.Use(middleware.RequestIDMiddleware, requestLogger, metricsMiddleware)
	router.Path("/metrics").Handler(promhttp.Handler())

	httpServer := http.Server{
		Addr:     ":" + apiPort,
		ErrorLog: slog.NewLogLogger(logJSONHandler, slog.LevelError),
		Handler:  router,
	}

	gracefulCloser := internal.NewGracefulCloser()
	gracefulCloser.Add(func() error {
		logger.Info("closing db connection")
		if err := db.Close(); err != nil {
			logger.Error(
				"error closing db connection",
				slog.String("err", err.Error()),
			)
			return err
		}
		logger.Info("db connection closed successfully")
		return nil
	})

	gracefulCloser.Add(func() error {
		logger.Info("shutting down HTTP server")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error(
				"error shutting down HTTP server",
				slog.String("err", err.Error()),
			)
			return err
		}
		logger.Info("HTTP server shut down successfully")
		return nil
	})
	sutdownCtx, shutdownCancel := context.WithCancel(ctx)
	defer shutdownCancel()
	go func() {
		ctx, cancel := context.WithCancel(sutdownCtx)
		defer cancel()
		gracefulCloser.Run(ctx, logger)
		os.Exit(1)
	}()

	logger.Info("application started", slog.String("port", apiPort))

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error("failed to start server", slog.String("err", err.Error()))
	}

	logger.Info("graceful shutdown complete", slog.String("port", apiPort))
}
