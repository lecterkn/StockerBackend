
-- +migrate Up
ALTER TABLE item_stocks DROP COLUMN place;

-- +migrate Down
