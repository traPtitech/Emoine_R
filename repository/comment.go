package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

// Comments are sorted by created_at ASC (oldest first)
func (r *Repository) SelectEventComments(ctx context.Context, eventID uuid.UUID) ([]dbmodel.Comment, error) {
	var comments []dbmodel.Comment
	err := r.DB.NewSelect().
		Model(&comments).
		Where("? = ?", bun.Ident("event_id"), eventID).
		OrderExpr("? ASC", bun.Ident("created_at")).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
