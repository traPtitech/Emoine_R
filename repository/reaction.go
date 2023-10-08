package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

// Reactions are sorted by created_at ASC (oldest first)
func (r *Repository) SelectEventReactions(ctx context.Context, eid uuid.UUID) ([]dbmodel.Reaction, error) {
	var reactions []dbmodel.Reaction
	err := r.DB.NewSelect().
		Model(&reactions).
		Where("? = ?", bun.Ident("event_id"), eid).
		OrderExpr("? ASC", bun.Ident("created_at")).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}
