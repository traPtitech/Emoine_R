package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

func (r *Repository) SelectEventTokens(ctx context.Context, eventID uuid.UUID) ([]dbmodel.Token, error) {
	var tokens []dbmodel.Token
	err := r.DB.NewSelect().
		Model(&tokens).
		Where("? = ?", bun.Ident("event_id"), eventID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
