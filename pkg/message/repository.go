package message

import (
	"context"
	"database/sql"
	"time"
)

type Message struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageRepository interface {
	Create(message *Message) (*Message, error)
	FindAll() ([]*Message, error)
	FindByID(id int) (*Message, error)
	Update(message *Message) error
	Delete(id int) error
}

type messageRepository struct {
	db *sql.DB
}

func newMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *Message) (*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, "INSERT INTO messages (content) VALUES ($1) RETURNING id, content, created_at, updated_at", message.Content)

	createdMessage := Message{}

	if err := row.Scan(&createdMessage.ID, &createdMessage.Content, &createdMessage.CreatedAt, &createdMessage.UpdatedAt); err != nil {
		return &Message{}, err
	}

	return &createdMessage, nil
}

func (r *messageRepository) FindAll() ([]*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := r.db.QueryContext(ctx, "select id, content, created_at, updated_at from messages")

	if err != nil {
		return []*Message{}, err
	}

	allMessages := []*Message{}

	defer rows.Close()

	for rows.Next() {
		message := Message{}

		if err := rows.Scan(&message.ID, &message.Content, &message.CreatedAt, &message.UpdatedAt); err != nil {
			return []*Message{}, err
		}

		allMessages = append(allMessages, &message)

	}

	return allMessages, nil
}

func (r *messageRepository) FindByID(id int) (*Message, error) {
	return &Message{}, nil
}

func (r *messageRepository) Update(message *Message) error {
	return nil
}

func (r *messageRepository) Delete(id int) error {
	return nil
}
