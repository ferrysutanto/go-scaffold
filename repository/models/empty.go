package models

import "github.com/pkg/errors"

type emptyModel struct{}

func (*emptyModel) Hello() error {
	return errors.New("[models][emptyModel][Hello] function is not implemented")
}
