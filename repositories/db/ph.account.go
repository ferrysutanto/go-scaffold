package db

import "context"

type phAccDB struct{}

func (*phAccDB) GetAccounts(ctx context.Context, param *ParamGetAccounts) (*Accounts, error) {
	return nil, errNotImplemented
}

func (*phAccDB) FindAccountByID(ctx context.Context, id string) (*Account, error) {
	return nil, errNotImplemented
}

func (*phAccDB) CreateAccount(ctx context.Context, acc *ParamCreateAccount) (*Account, error) {
	return nil, errNotImplemented
}

func (*phAccDB) UpdateAccount(ctx context.Context, acc *ParamUpdateAccount) (*Account, error) {
	return nil, errNotImplemented
}

func (*phAccDB) PatchAccount(ctx context.Context, acc *ParamPatchAccount) error {
	return errNotImplemented
}

func (*phAccDB) DeleteAccountByID(ctx context.Context, id string) error {
	return errNotImplemented
}

func (*phAccDB) BeginTx(ctx context.Context) (IAccountTx, error) {
	return &phAccTx{}, nil
}

type phAccTx struct{}

func (*phAccTx) Commit(ctx context.Context) error {
	return errNotImplemented
}

func (*phAccTx) Rollback(ctx context.Context) error {
	return errNotImplemented
}

func (*phAccTx) CreateAccount(ctx context.Context, acc *ParamCreateAccount) (*Account, error) {
	return nil, errNotImplemented
}

func (*phAccTx) UpdateAccount(ctx context.Context, acc *ParamUpdateAccount) (*Account, error) {
	return nil, errNotImplemented
}

func (*phAccTx) PatchAccount(ctx context.Context, acc *ParamPatchAccount) error {
	return errNotImplemented
}

func (*phAccTx) DeleteAccountByID(ctx context.Context, id string) error {
	return errNotImplemented
}
