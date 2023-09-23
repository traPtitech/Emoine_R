package dbmodel

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var AllModels = []any{
	(*Event)(nil),
	(*Comment)(nil),
	(*Reaction)(nil),
	(*Token)(nil),
}

type Event struct {
	bun.BaseModel `bun:"table:events"`

	ID          uuid.UUID      `bun:",pk,notnull,type:uuid,default:UUID()"`
	VideoID     string         `bun:",notnull,type:char(11)"`
	Title       string         `bun:",notnull,nullzero,type:varchar(64)"`
	Thumbnail   string         `bun:",notnull,nullzero"`
	Description sql.NullString `bun:",type:text"`
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
	UserID      string         `bun:",notnull,nullzero,type:varchar(32)"` // TODO: usernameにする？
	EventID     uuid.UUID      `bun:",notnull,type:uuid"`
	Text        string         `bun:",notnull,nullzero,type:text"`
	IsAnonymous bool           `bun:",notnull"`
	Color       sql.NullString `bun:",notnull,nullzero,type:varchar(7)"`
	CreatedAt   time.Time      `bun:",notnull,default:CURRENT_TIMESTAMP"`
}

type Reaction struct {
	bun.BaseModel `bun:"table:reactions"`

	ID        uuid.UUID `bun:",pk,notnull,type:uuid,default:UUID()"`
	UserID    string    `bun:",notnull,nullzero,type:varchar(32)"`
	EventID   uuid.UUID `bun:",notnull,type:uuid"`
	StampID   uuid.UUID `bun:",notnull,type:uuid"`
	CreatedAt time.Time `bun:",notnull,default:CURRENT_TIMESTAMP"`
}

type Token struct {
	bun.BaseModel `bun:"table:token"`

	ID          uuid.UUID      `bun:",pk,notnull,type:uuid,default:UUID()"`
	Value       string         `bun:",notnull,type:char(44)"`
	UserID      string         `bun:",notnull,nullzero,type:varchar(32)"`
	CreatorID   string         `bun:",notnull,nullzero,type:varchar(32)"`
	EventID     uuid.UUID      `bun:",notnull,type:uuid"`
	Description sql.NullString `bun:",type:text"`
	CreatedAt   time.Time      `bun:",notnull,default:CURRENT_TIMESTAMP"`
	ExpireAt    sql.NullTime   `bun:""`
}
