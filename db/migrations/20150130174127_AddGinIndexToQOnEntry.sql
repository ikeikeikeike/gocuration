
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE INDEX index_entry_on_q ON entry USING gin (q gin_trgm_ops);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX index_entry_on_q;
