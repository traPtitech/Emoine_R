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

## xoのテンプレート変更

[./templates/xo/go](./templates/xo/go)以下を変更することで、xoの出力を変更可能です。

[xo/xo](https://github.com/xo/xo)に追従するには以下のコマンドを実行してください。その後、差分を確認して、必要に応じて修正してください。

```sh
go run github.com/xo/xo@latest dump -t go templates/xo/go
```
