-- +goose Up
-- +goose StatementBegin
--SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS recipes_to_ingredients(
  recipe_id SERIAL REFERENCES recipes(id),
  ingredient_id SERIAL REFERENCES ingredients(id),
  quantity INTEGER,
  PRIMARY KEY (recipe_id, ingredient_id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--SELECT 'down SQL query';
DROP TABLE IF EXISTS recipes_to_ingredients;
-- +goose StatementEnd
