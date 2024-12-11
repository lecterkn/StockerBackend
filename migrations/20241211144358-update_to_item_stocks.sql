
-- +migrate Up
ALTER TABLE item_stocks ADD COLUMN price INTEGER COMMENT "定価";
ALTER TABLE item_stocks DROP COLUMN place;

-- +migrate Down
