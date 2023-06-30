package service

import "github.com/pkg/errors"

type emptyService struct{}

func (*emptyService) Hello() error {
	return errors.New("[service] empty service: Hello is not implemented")
}
