# Stockerバックエンド

## 開発環境

開発言語：`go言語`
データベース：`MySQL`
仮想環境：`docker`

## MySQL接続

```shell
docker-compose exec mysql mysql -u root -p
```

## 仕様ライブラリ

- github.com/google/uuid
- github.com/rubenv/sql-migrate/...
- gorm.io/gorm

## 実行方法

```shell
go mod tidy
go run cmd/stocker/main.go
```

## 作業ログ

UUID

```shell
go get github.com/google/uuid
```

sql-migrate

```shell
go get -v github.com/rubenv/sql-migrate/...
go install github.com/rubenv/sql-migrate/...@latest
```

gorm

```shell
go get -u gorm.io/gorm
```