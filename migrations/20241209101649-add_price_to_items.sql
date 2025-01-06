
-- +migrate Up
ALTER TABLE item_stocks ADD price INTEGER COMMENT "定価";

-- +migrate Down
ALTER TABLE item_stocks DROP COLUMN price;
