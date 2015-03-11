
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE INDEX index_diva_on_name ON diva USING gin (name gin_trgm_ops);
CREATE INDEX index_character_on_name ON character USING gin (name gin_trgm_ops);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX index_diva_on_name
DROP INDEX index_character_on_name
