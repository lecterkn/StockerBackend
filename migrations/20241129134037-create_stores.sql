
-- +migrate Up
CREATE TABLE stores(
    id BINARY(16) PRIMARY KEY COMMENT "店舗ID",
    user_id BINARY(16) NOT NULL COMMENT "ユーザーID",
    name VARCHAR(255) NOT NULL COMMENT "店舗名",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時"
) COMMENT = "店舗";

-- +migrate Down
DROP TABLE stores;