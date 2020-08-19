-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO products(name) VALUES ('IF-2291');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM products WHERE name = 'IF-2291';