-- +goose Up
CREATE TABLE subs(
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    url TEXT NOT NULL unique
);

-- +goose Down

DROP TABLE IF EXISTS subs;