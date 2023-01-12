# Emoine_R-server

.env ファイルを作り、以下の項目を入力しておくこと。例：

```
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