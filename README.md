# Emoine_R-server

.env ファイルを作り、DB_PASSWORD と DB_NAME と ADMIN_NAMES を入力しておくこと。例：

```
DB_PASSWORD = "example"
DB_NAME = "emoine"
ADMIN_NAMES = "aaa,bbb,ccc"
```

## 起動

```bash
docker compose up -d
go run main.go
```

