-- +goose Up
-- +goose StatementBegin
--SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS potions (
  uuid VARCHAR(36) PRIMARY KEY,
  witch_uuid VARCHAR REFERENCES witches(uuid), 
  recipe_uuid VARCHAR REFERENCES recipes(uuid),
  status VARCHAR ,
  created_at TIMESTAMP,
  updated_at TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--SELECT 'down SQL query';
DROP TABLE IF EXISTS potions;
-- +goose StatementEnd