package services

import (
	"context"
)

type svcPH struct{}

func placeholder() IService {
	return &svcPH{}
}

// Healthcheck empty placeholder
func (*svcPH) Healthcheck(ctx context.Context) error {
	return ErrNotImplemented
}

// RecordSignup empty placeholder
func (*svcPH) RecordSignup(ctx context.Context, param *ParamRecordSignup) (*RespRecordSignup, error) {
	return nil, ErrNotImplemented
}

// DeleteAccount empty placeholder
func (*svcPH) DeleteAccount(ctx context.Context, param *ParamRequestAccountDeletion) error {
	return ErrNotImplemented
}
