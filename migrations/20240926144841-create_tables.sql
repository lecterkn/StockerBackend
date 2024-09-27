
-- +migrate Up
CREATE TABLE items (
    id BINARY(16) PRIMARY KEY COMMENT "商品ID",
    name VARCHAR(255) NOT NULL COMMENT "商品名",
    jan_code VARCHAR(255) NOT NULL COMMENT "JANコード",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時"
) COMMENT = "商品";

CREATE TABLE item_stocks (
    item_id BINARY(16) NOT NULL COMMENT "商品ID",
    place VARCHAR(255) COMMENT "場所",
    stock INTEGER NOT NULL COMMENT "在庫数",
    stock_min INTEGER NOT NULL COMMENT "在庫の最低値",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    FOREIGN KEY (item_id) REFERENCES items(id),
    INDEX idx_is_item_id (`item_id`)
) COMMENT = "商品詳細";

CREATE TABLE stock_ins (
    id BINARY(16) PRIMARY KEY COMMENT "入荷ID",
    item_id BINARY(16) NOT NULL COMMENT "商品ID",
    stocks INTEGER NOT NULL COMMENT "入荷数",
    price INTEGER NOT NULL COMMENT "入荷価格",
    created_at DATETIME NOT NULL COMMENT "入荷日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    FOREIGN KEY (item_id) REFERENCES items(id),
    INDEX idx_si_item_id (`item_id`)
) COMMENT = "入荷履歴";

CREATE TABLE stock_outs (
    id BINARY(16) PRIMARY KEY COMMENT "販売ID",
    item_id BINARY(16) NOT NULL COMMENT "商品ID",
    stocks INTEGER NOT NULL COMMENT "販売数",
    price INTEGER NOT NULL COMMENT "販売価格",
    created_at DATETIME NOT NULL COMMENT "販売日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    FOREIGN KEY(item_id) REFERENCES items(id),
    INDEX idx_so_item_id (`item_id`)
) COMMENT = "販売履歴";

CREATE TABLE returns (
    id BINARY(16) PRIMARY KEY COMMENT "返品ID",
    item_id BINARY(16) NOT NULL COMMENT "商品ID",
    stocks INTEGER NOT NULL COMMENT "返品数",
    price INTEGER NOT NULL COMMENT "返金額",
    created_at DATETIME NOT NULL COMMENT "返品日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    FOREIGN KEY(item_id) REFERENCES items(id),
    INDEX idx_r_item_id (`item_id`)
) COMMENT = "返金履歴";

-- +migrate Down
DROP TABLE item_stocks;
DROP TABLE stock_ins;
DROP TABLE stock_outs;
DROP TABLE returns;
DROP TABLE items;