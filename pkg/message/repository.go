package message

import (
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
	Create(message *Message) error
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

func (r *messageRepository) Create(message *Message) error {
	return nil
}

func (r *messageRepository) FindAll() ([]*Message, error) {
	rows, err := r.db.Query("select id, content, created_at, updated_at from messages")

	if err != nil {
		return nil, err
	}

	allMessages := []*Message{}

	defer rows.Close()

	for rows.Next() {
		message := Message{}

		if err := rows.Scan(&message.ID, &message.Content, &message.CreatedAt, &message.UpdatedAt); err != nil {
			return nil, err
		}

		allMessages = append(allMessages, &message)

	}

	return allMessages, nil
}

func (r *messageRepository) FindByID(id int) (*Message, error) {
	return nil, nil
}

func (r *messageRepository) Update(message *Message) error {
	return nil
}

func (r *messageRepository) Delete(id int) error {
	return nil
}
