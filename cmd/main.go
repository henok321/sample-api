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

	server := &http.Server{
		Handler:     router,
		Addr:        ":8080",
		ReadTimeout: time.Second * 10,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			os.Exit(1)
		}
	}()
	slog.Info("Server started")
	<-done
	slog.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error", "error", err)
		exitCode = 1
		return
	}

}
