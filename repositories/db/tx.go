package db

import "context"

type txImpl struct{}

func (*txImpl) Commit(ctx context.Context) error {
	return errNotImplemented
}

func (*txImpl) Rollback(ctx context.Context) error {
	return errNotImplemented
}

func (*txImpl) CreateAccount(ctx context.Context, account *ParamCreateAccount) (*Account, error) {
	return nil, errNotImplemented
}
func (*txImpl) UpdateAccount(ctx context.Context, account *ParamUpdateAccount) (*Account, error) {
	return nil, errNotImplemented
}
func (*txImpl) PatchAccount(ctx context.Context, account *ParamPatchAccount) (*Account, error) {
	return nil, errNotImplemented
}
func (*txImpl) DeleteAccountByID(ctx context.Context, id string) error {
	return errNotImplemented
}

func (*txImpl) CreateProfile(ctx context.Context, account *ParamCreateProfile) (*Profile, error) {
	return nil, errNotImplemented
}
func (*txImpl) UpdateProfile(ctx context.Context, account *ParamUpdateProfile) (*Profile, error) {
	return nil, errNotImplemented
}
func (*txImpl) PatchProfile(ctx context.Context, account *ParamPatchProfile) (*Profile, error) {
	return nil, errNotImplemented
}
func (*txImpl) DeleteProfileByID(ctx context.Context, id string) error {
	return errNotImplemented
}
