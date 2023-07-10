package models

import (
	"context"

	"github.com/pkg/errors"
)

func (m *pgModel) Ping(ctx context.Context) error {
	if err := m.db.PingContext(ctx); err != nil {
		return errors.Wrap(err, "[models][pgModel][Ping] failed to ping db")
	}

	if err := m.replDb.PingContext(ctx); err != nil {
		return errors.Wrap(err, "[models][pgModel][Ping] failed to ping replication db")
	}

	return nil
}
