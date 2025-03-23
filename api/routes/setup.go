package routes

import (
	"database/sql"
	"net/http"
	"sample-api/api/handler"
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
	messageHandler := handler.NewMessageHandler(message.InitalizeMessageModule(s.db))

	s.router.Handle("GET /messages/{id}", http.HandlerFunc(messageHandler.FindByID))
	s.router.Handle("GET /messages", http.HandlerFunc(messageHandler.FindAll))
	s.router.Handle("POST /messages", http.HandlerFunc(messageHandler.Create))
	s.router.Handle("PUT /messages/{id}", http.HandlerFunc(messageHandler.Update))
	s.router.Handle("DELETE /messages/{id}", http.HandlerFunc(messageHandler.Delete))

}
