
-- +migrate Up
ALTER TABLE stock_ins DROP COLUMN place;

-- +migrate Down
