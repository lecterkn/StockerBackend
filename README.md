# Stockerバックエンド

## 開発環境

開発言語：`go 1.23.1`
データベース：`MySQL`
フレームワーク：`fiber`

## 前提条件

- go言語がインストールされていること
- dockerがインストールされていること

## 実行手順

1. 依存関係を取得する

```shell
# 依存関係ダウンロード
go mod tidy

# swag インストール
go install github.com/swaggo/swag/cmd/swag@latest

#sql-migrate インストール

go install github.com/rubenv/sql-migrate/...@latest
```

2. MySQLのサーバーを起動する

```shell
docker compose up -d
```

3. データベースのマイグレーションを行う

```shell
sql-migrate up
```

4. .envを用意する

.env.exampleを.envとしてフォルダ内にコピーする

```shell
cp .env.example .env
```

5. アプリケーションを実行

```shell
go run cmd/stocker/main.go
```

## APIドキュメント

アプリケーションを開いた状態で[Swaggerのページ](http://127.0.0.1:8810/swagger/index.html)を開く

```http
http://127.0.0.1:8810/swagger/index.html
```

生成方法

```shell
swag init -g cmd/stocker/main.go 
```

## 依存性注入

- wireを使用して依存性注入を行う
- `internal\stocker\wire.go`に依存関係を記載する
- 下記のコマンドで`wire_gen.go`を生成する

```shell
wire gen ./internal/stocker/
```

## MySQL接続方法

直接データベースを確認したい場合の確認方法

```shell
docker-compose exec stockergo-mysql mysql -u root -p
```
