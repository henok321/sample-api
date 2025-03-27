package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sample-api/api/routes"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"time"
)

func main() {
	exitCode := 0

	defer func() {
		os.Exit(exitCode)
	}()

	databaseURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		slog.Error("Database connection error", "error", err)
		exitCode = 1
		return
	}

	err = db.Ping()

	if err != nil {
		slog.Error("Database ping error", "error", err)
		exitCode = 1
		return
	}

	router := routes.SetupRouter(db)

	apiServer := &http.Server{
		Handler:     router,
		Addr:        ":8080",
		ReadTimeout: time.Second * 10,
	}

	metricsRouter := http.NewServeMux()
	metricsRouter.Handle("GET /metrics", promhttp.Handler())

	metricsServer := &http.Server{
		Handler:     metricsRouter,
		Addr:        ":9090",
		ReadTimeout: time.Second * 10,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		slog.Info("Starting metrics server", "port", "9090")
		if err := metricsServer.ListenAndServe(); err != nil {
			os.Exit(1)
		}
	}()

	go func() {
		slog.Info("Starting api server", "port", "8080")
		if err := apiServer.ListenAndServe(); err != nil {
			os.Exit(1)
		}
	}()

	<-done
	slog.Info("Shutdown signal received, shutting down servers")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := metricsServer.Shutdown(ctx); err != nil {
		slog.Error("Metrics server shutdown error", "error", err)
		exitCode = 1
		return
	}

	if err := apiServer.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error", "error", err)
		exitCode = 1
		return
	}

}
