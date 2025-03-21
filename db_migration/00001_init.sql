-- +goose Up

CREATE TABLE messages
(
    id               SERIAL PRIMARY KEY,
    name             VARCHAR(255)             NOT NULL
);
