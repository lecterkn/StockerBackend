
-- +migrate Up
ALTER TABLE item_stocks ADD price INTEGER NOT NULL COMMENT "定価";

-- +migrate Down
