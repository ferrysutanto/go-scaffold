package services

import "github.com/pkg/errors"

type emptyService struct{}

func (*emptyService) Hello() error {
	return errors.New("[services][emptyService] Hello is not implemented")
}
