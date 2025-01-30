package impl

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/services"
)

func (this *srvImpl) DeleteAccount(ctx context.Context, req *services.ParamRequestAccountDeletion) error {
	if err := this.db.Account(ctx).DeleteAccountByID(ctx, req.ID); err != nil {
		return err
	}

	return nil
}
