package impl

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/ferrysutanto/go-scaffold/services"
	"github.com/ferrysutanto/go-scaffold/utils"
)

func (this *srvImpl) RecordSignup(ctx context.Context, param *services.ParamRecordSignup) (*services.RespRecordSignup, error) {
	// 1. generate username
	username, err := utils.GenerateUsername(&utils.ParamGenerateUsername{
		Email: param.Email,
	})
	if err != nil {
		return nil, err
	}

	// 2. store account data to db
	acc, err := this.db.Account(ctx).CreateAccount(ctx, &db.ParamCreateAccount{
		Email:    param.Email,
		Phone:    param.Phone,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	// 3. return response
	return &services.RespRecordSignup{
		Data: &services.Signup{
			ID:           acc.ID,
			Email:        acc.Email,
			Phone:        acc.Phone,
			Username:     acc.Username,
			AuthProvider: param.AuthProvider,
		},
	}, nil
}
