package message

import (
	"database/sql"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
	return nil, nil
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
