
-- +migrate Up
ALTER TABLE items ADD CONSTRAINT idx_i_jan_code_and_name UNIQUE(name, jan_code);

-- +migrate Down

ALTER TABLE items DROP INDEX idx_i_jan_code_and_name;
