package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sample-api/pkg/message"
	"strconv"
	"time"
)

type CreateMessageRequest struct {
	Content string `json:"content"`
}

type UpdateMessageRequest struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type MessageResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MessageListResponse struct {
	Messages []*MessageResponse `json:"messages"`
}

type MessageHandler struct {
	service message.MessageService
}

func NewMessageHandler(service message.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (h *MessageHandler) Create(writer http.ResponseWriter, request *http.Request) {
	messageCreateRequest := CreateMessageRequest{}

	if err := json.NewDecoder(request.Body).Decode(&messageCreateRequest); err != nil {
		slog.ErrorContext(request.Context(), "Failed to decode request body", "error", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(&message.Message{Content: messageCreateRequest.Content})

	if err != nil {
		slog.ErrorContext(request.Context(), "Failed to create message", "error", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Location", "/messages/"+strconv.Itoa(id))
}

func (h *MessageHandler) FindAll(writer http.ResponseWriter, request *http.Request) {
	allMessages, err := h.service.FindAll()

	if err != nil {
		slog.ErrorContext(request.Context(), "Failed to find all messages", "error", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	responseBody := MessageListResponse{
		Messages: []*MessageResponse{},
	}

	for _, message := range allMessages {
		responseBody.Messages = append(responseBody.Messages, &MessageResponse{
			ID:        message.ID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
		})
	}
	if err := json.NewEncoder(writer).Encode(responseBody); err != nil {
		slog.ErrorContext(request.Context(), "Failed to encode response", "error", err)
		return
	}
}

func (h *MessageHandler) FindByID(writer http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	message, err := h.service.FindByID(id)

	if err != nil {
		slog.ErrorContext(request.Context(), "Failed to find message by ID", "error", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if message == nil {
		slog.Info("Message not found", "id", id)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	responseBody := MessageResponse{
		ID:        message.ID,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,

		UpdatedAt: message.UpdatedAt,
	}

	if err := json.NewEncoder(writer).Encode(responseBody); err != nil {
		slog.ErrorContext(request.Context(), "Failed to encode response", "error", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *MessageHandler) Update(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func (h *MessageHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}
