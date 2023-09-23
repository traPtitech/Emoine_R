package migrations

import (
	"context"
	"log/slog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations(
	migrate.WithMigrationsDirectory("migrations"),
)

func Migrate(ctx context.Context, db *bun.DB) error {
	m := migrate.NewMigrator(db, Migrations)
	if err := m.Init(ctx); err != nil {
		return err
	}

	if err := m.Lock(ctx); err != nil {
		return err
	}
	defer m.Unlock(ctx) //nolint:errcheck

	g, err := m.Migrate(ctx)
	if err != nil {
		return err
	}

	if g == nil || g.IsZero() {
		slog.Info("DB already up to date")

		return nil
	}

	slog.Info("DB migrated successfully", slog.Int64("id", g.ID))

	return nil
}
