package dbmodel

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel `bun:"table:events"`

	ID          uuid.UUID      `bun:",pk,notnull,type:uuid,default:UUID()"`
	VideoID     string         `bun:",notnull"`
	Title       string         `bun:",notnull"`
	Thumbnail   string         `bun:",notnull"`
	Description sql.NullString `bun:""`
	StartedAt   time.Time      `bun:",notnull,default:CURRENT_TIMESTAMP"`
	EndedAt     sql.NullTime   `bun:""`
}
