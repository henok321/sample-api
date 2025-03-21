package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	exitCode := 0

	defer func() {
		os.Exit(exitCode)
	}()

	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Health check")
		_, err := writer.Write([]byte("{'status': 'ok'}"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /foo", func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Foo request")
		_, err := writer.Write([]byte("{'message': 'foo'}"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /bar", func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Foo request")
		_, err := writer.Write([]byte("{'message': 'foo'}"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)
	})

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
