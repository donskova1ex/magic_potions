-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS witches (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(36) unique,
    name VARCHAR(64) unique );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS witches;
-- +goose StatementEnd


