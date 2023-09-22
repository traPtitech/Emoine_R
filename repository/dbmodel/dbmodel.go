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

	// event has many comments
	Comments []*Comment `bun:"rel:has-many,join:id=event_id"`
	// event has many reactions
	Reactions []*Reaction `bun:"rel:has-many,join:id=event_id"`
	// event has many tokens
	Tokens []*Token `bun:"rel:has-many,join:id=event_id"`
}

type Comment struct {
	bun.BaseModel `bun:"table:comments"`

	ID          uuid.UUID      `bun:",pk,notnull,type:uuid,default:UUID()"`
	UserID      string         `bun:",notnull"` // TODO: usernameにする？
	EventID     uuid.UUID      `bun:",notnull,type:uuid"`
	Text        string         `bun:",notnull"`
	IsAnonymous bool           `bun:",notnull"`
	Color       sql.NullString `bun:""`
	CreatedAt   time.Time      `bun:",notnull,default:CURRENT_TIMESTAMP"`
}

type Reaction struct {
	bun.BaseModel `bun:"table:reactions"`

	ID        uuid.UUID `bun:",pk,notnull,type:uuid,default:UUID()"`
	UserID    string    `bun:",notnull"`
	EventID   uuid.UUID `bun:",notnull,type:uuid"`
	StampID   uuid.UUID `bun:",notnull,type:uuid"`
	CreatedAt time.Time `bun:",notnull,default:CURRENT_TIMESTAMP"`
}

type Token struct {
	bun.BaseModel `bun:"table:token"`

	ID          uuid.UUID      `bun:",pk,notnull,type:uuid,default:UUID()"`
	Value       string         `bun:",notnull"`
	UserID      string         `bun:",notnull"`
	CreatorID   string         `bun:",notnull"`
	EventID     uuid.UUID      `bun:",notnull,type:uuid"`
	Description sql.NullString `bun:""`
	CreatedAt   time.Time      `bun:",notnull,default:CURRENT_TIMESTAMP"`
	ExpireAt    sql.NullTime   `bun:""`
}
