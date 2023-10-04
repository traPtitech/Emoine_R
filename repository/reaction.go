package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

func (r *Repository) SelectEventReactions(ctx context.Context, eid uuid.UUID) ([]dbmodel.Reaction, error) {
	var reactions []dbmodel.Reaction
	err := r.DB.NewSelect().
		Model(&reactions).
		Where("? = ?", bun.Ident("event_id"), eid).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}
