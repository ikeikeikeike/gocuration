
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE INDEX index_diva_on_kana ON diva USING gin (kana gin_trgm_ops);
CREATE INDEX index_character_on_kana ON character USING gin (kana gin_trgm_ops);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX index_diva_on_kana
DROP INDEX index_character_on_kana
