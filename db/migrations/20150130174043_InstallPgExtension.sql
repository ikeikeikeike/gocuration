
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS pg_trgm;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP EXTENSION IF EXISTS pg_trgm;
