//go:generate go run github.com/xo/xo@latest schema mariadb://root:password@127.0.0.1:3306/emoine --go-pkg dbschema -o .
package dbschema

import "github.com/google/uuid"

// NOTE: xoがMySQLのUUID型をサポートしていないため、ここで定義する
type UUID = uuid.UUID
