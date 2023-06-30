package models

import "github.com/pkg/errors"

type emptyModel struct{}

func (*emptyModel) Hello() error {
	return errors.New("[models] empty model: Hello is not implemented")
}
