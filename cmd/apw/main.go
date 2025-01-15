package main

import (
	"apw/internal/config"
	handler "apw/internal/handlers"
	repo "apw/internal/repository"
	"apw/internal/server"
	"apw/internal/service"
	"apw/internal/storage"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Logger.Env)
	log.Info(
		"start logger in",
		slog.String("env", cfg.Logger.Env),
	)

	log.Debug("storage", cfg.Storage)
	log.Debug("server", cfg.Server)

	storage, err := storage.New(cfg.Storage.URL)
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repos := repo.NewRepository(storage)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("starting server", slog.String("address", cfg.Server.Address))

	srv := server.New(cfg.Server, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil {
			log.Error("failed to stop server", slog.String("error", err.Error()))
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", slog.String("error", err.Error()))
		return
	}
}
