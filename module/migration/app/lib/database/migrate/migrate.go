package migrate

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"{{{ .Package }}}/app/lib/database"
	"{{{ .Package }}}/app/lib/telemetry"
	"{{{ .Package }}}/app/util"
)

func Migrate(ctx context.Context, s *database.Service, logger util.Logger) error {
	logger = logger.With("svc", "migrate")
	ctx, span, logger := telemetry.StartSpan(ctx, "database:migrate", logger)
	defer span.Complete()

	err := createMigrationTableIfNeeded(ctx, s, nil, logger)
	if err != nil {
		return errors.Wrap(err, "unable to create migration table")
	}

	tx, err := s.StartTransaction(logger)
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}
	defer func() {
		_ = tx.Rollback()
	}()

	maxIdx := maxMigrationIdx(ctx, s, tx, logger)

	if len(databaseMigrations) > maxIdx+1 {
		c := len(databaseMigrations) - maxIdx
		logger.Infof("applying [%d] database %s...", c, util.StringPluralMaybe("migration", c))
	}

	for i, file := range databaseMigrations {
		err = run(ctx, maxIdx, i, file, s, tx, logger)
		if err != nil {
			return errors.Wrapf(err, "error running database migration [%s]", file.Title)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	logger.Infof("verified [%d] database %s", maxIdx, util.StringPluralMaybe("migration", maxIdx))
	return nil
}

func run(ctx context.Context, maxIdx int, i int, file *MigrationFile, s *database.Service, tx *sqlx.Tx, logger util.Logger) error {
	idx := i + 1
	switch {
	case idx == maxIdx:
		m := getMigrationByIdx(ctx, s, maxIdx, tx, logger)
		if m == nil {
			return nil
		}
		if m.Title != file.Title {
			logger.Infof("migration [%d] name has changed from [%s] to [%s]", idx, m.Title, file.Title)
			err := removeMigrationByIdx(ctx, s, idx, tx, logger)
			if err != nil {
				return err
			}
			err = applyMigration(ctx, s, idx, file, tx, logger)
			if err != nil {
				return err
			}
			return nil
		}
		nc := file.Content
		if nc != m.Src {
			logger.Infof("migration [%d:%s] content has changed from [%dB] to [%dB]", idx, file.Title, len(nc), len(m.Src))
			err := removeMigrationByIdx(ctx, s, idx, tx, logger)
			if err != nil {
				return err
			}
			err = applyMigration(ctx, s, idx, file, tx, logger)
			if err != nil {
				return err
			}
		}
	case idx > maxIdx:
		err := applyMigration(ctx, s, idx, file, tx, logger)
		if err != nil {
			return err
		}
	default:
		// noop
	}
	return nil
}

func applyMigration(ctx context.Context, s *database.Service, idx int, file *MigrationFile, tx *sqlx.Tx, logger util.Logger) error {
	logger.Infof("applying database migration [%d]: %s", idx, file.Title)
	sql, err := exec(ctx, file, s, tx, logger)
	if err != nil {
		return err
	}
	m := &Migration{Idx: idx, Title: file.Title, Src: sql, Created: time.Now()}
	return newMigration(ctx, s, m, tx, logger)
}
