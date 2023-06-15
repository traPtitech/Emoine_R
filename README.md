# Emoine_R-server

.env ファイルを作り、以下の項目を入力しておくこと。例：

```sh
DB_PASSWORD = "example"
DB_NAME = "emoine"
ADMIN_NAMES = "aaa,bbb,ccc"
CLIENT_ID = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
```

## 起動

```bash
docker compose up -d
go run main.go
```

## Linter

```sh
go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix ./...
```

## コード生成

- OpenAPI Schemas: [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)

```sh
go generate ./...
```
