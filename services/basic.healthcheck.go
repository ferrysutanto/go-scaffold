package services

import (
	"context"

	"github.com/pkg/errors"
)

func (s *basicService) Healthcheck(ctx context.Context) error {
	if err := s.m.Ping(ctx); err != nil {
		return errors.Wrap(err, "[services][basicService][Healthcheck] failed to ping model")
	}

	return nil
}
