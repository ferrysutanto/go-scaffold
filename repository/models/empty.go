package models

import (
	"context"

	"github.com/pkg/errors"
)

type emptyModel struct{}

func (*emptyModel) Ping(context.Context) error {
	return errors.New("[models][emptyModel][Ping] function is not implemented")
}
