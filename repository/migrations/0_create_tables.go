package migrations

import (
	"context"

	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"github.com/uptrace/bun"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			if err := tx.Rollback(); err != nil {
				panic(err)
			}
		}()

		for _, model := range dbmodel.AllModels {
			if _, err := tx.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}

	down := func(ctx context.Context, db *bun.DB) error {
		if _, err := db.NewDropTable().Model(&dbmodel.Event{}).IfExists().Exec(ctx); err != nil {
			return err
		}

		return nil
	}

	Migrations.MustRegister(up, down)
}
