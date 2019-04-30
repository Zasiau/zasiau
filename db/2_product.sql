-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS products (
  id              serial primary key,
  name            VARCHAR(255) NOT NULL,
  price           Integer NOT NULL,
  created         timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated         timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE products;
