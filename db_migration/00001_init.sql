-- +goose Up

CREATE TABLE messages
(
    id               SERIAL PRIMARY KEY,
    content             VARCHAR(255)             NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
