package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

func (r *Repository) SelectEvents(ctx context.Context, limit int, offset int) ([]dbmodel.Event, int, error) {
	var events []dbmodel.Event
	cnt, err := r.DB.NewSelect().
		Model(&events).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return events, cnt, nil
}

func (r *Repository) SelectEvent(ctx context.Context, id uuid.UUID) (*dbmodel.Event, error) {
	var event dbmodel.Event
	err := r.DB.NewSelect().
		Model(&event).
		Where("? = ?", bun.Ident("id"), id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *Repository) InsertEvent(ctx context.Context, event *dbmodel.Event) error {
	_, err := r.DB.NewInsert().
		Model(event).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
