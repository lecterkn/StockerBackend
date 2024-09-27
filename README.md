# H11 backend

## MySQL接続
```
docker-compose exec mysql mysql -u root -p
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