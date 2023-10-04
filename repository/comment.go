package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

func (r *Repository) SelectEventComments(ctx context.Context, eventID uuid.UUID) ([]dbmodel.Comment, error) {
	var comments []dbmodel.Comment
	err := r.DB.NewSelect().
		Model(&comments).
		Where("? = ?", bun.Ident("event_id"), eventID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
