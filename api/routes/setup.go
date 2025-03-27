package routes

import (
	"database/sql"
	"log/slog"
	"net/http"
	"sample-api/api/handlers"
	"sample-api/api/middleware"
	"sample-api/pkg/message"
)

type RouteSetup struct {
	db     *sql.DB
	router *http.ServeMux
}

func SetupRouter(db *sql.DB) *http.ServeMux {
	instance := RouteSetup{
		db:     db,
		router: http.NewServeMux(),
	}

	instance.init()

	return instance.router
}

func (s *RouteSetup) init() {
	s.router.Handle("GET /health", middleware.Metrics(middleware.RequestLogging(slog.LevelDebug, http.HandlerFunc(handlers.HealthCheck))))

	messageHandler := handlers.NewMessageHandler(message.InitalizeMessageModule(s.db))

	s.router.Handle("GET /messages/{id}", middleware.Metrics(middleware.RequestLogging(slog.LevelInfo, http.HandlerFunc(messageHandler.FindByID))))
	s.router.Handle("GET /messages", middleware.Metrics(middleware.RequestLogging(slog.LevelInfo, http.HandlerFunc(messageHandler.FindAll))))
	s.router.Handle("POST /messages", middleware.Metrics(middleware.RequestLogging(slog.LevelInfo, http.HandlerFunc(messageHandler.Create))))
	s.router.Handle("PUT /messages/{id}", middleware.Metrics(middleware.RequestLogging(slog.LevelInfo, http.HandlerFunc(messageHandler.Update))))
	s.router.Handle("DELETE /messages/{id}", middleware.Metrics(middleware.RequestLogging(slog.LevelInfo, http.HandlerFunc(messageHandler.Delete))))

}
