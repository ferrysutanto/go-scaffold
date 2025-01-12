package pg

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/jmoiron/sqlx"
)

type PgTx struct {
	tx *sqlx.Tx

	accountTx db.IAccountTx
	profileTx db.IProfileTx
}

func (this *PgTx) Commit(ctx context.Context) error {
	return this.tx.Commit()
}

func (this *PgTx) Rollback(ctx context.Context) error {
	return this.tx.Rollback()
}

func (this *PgTx) CreateAccount(ctx context.Context, account *db.ParamCreateAccount) (*db.Account, error) {
	return this.accountTx.CreateAccount(ctx, account)
}
func (this *PgTx) UpdateAccount(ctx context.Context, account *db.ParamUpdateAccount) (*db.Account, error) {
	return this.accountTx.UpdateAccount(ctx, account)
}
func (this *PgTx) PatchAccount(ctx context.Context, account *db.ParamPatchAccount) error {
	return this.accountTx.PatchAccount(ctx, account)
}
func (this *PgTx) DeleteAccountByID(ctx context.Context, id string) error {
	return this.accountTx.DeleteAccountByID(ctx, id)
}

func (this *PgTx) CreateProfile(ctx context.Context, account *db.ParamCreateProfile) (*db.Profile, error) {
	return this.profileTx.CreateProfile(ctx, account)
}
func (this *PgTx) UpdateProfile(ctx context.Context, account *db.ParamUpdateProfile) (*db.Profile, error) {
	return this.profileTx.UpdateProfile(ctx, account)
}
func (this *PgTx) PatchProfile(ctx context.Context, account *db.ParamPatchProfile) error {
	return this.profileTx.PatchProfile(ctx, account)
}
func (this *PgTx) DeleteProfileByID(ctx context.Context, id string) error {
	return this.profileTx.DeleteProfileByID(ctx, id)
}
