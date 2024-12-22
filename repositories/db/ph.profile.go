package db

import "context"

type phProfDB struct{}

func (*phProfDB) GetProfiles(ctx context.Context, param *ParamGetProfiles) (*Profiles, error) {
	return nil, errNotImplemented
}

func (*phProfDB) FindProfileByID(ctx context.Context, id string) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfDB) CreateProfile(ctx context.Context, acc *ParamCreateProfile) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfDB) UpdateProfile(ctx context.Context, acc *ParamUpdateProfile) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfDB) PatchProfile(ctx context.Context, acc *ParamPatchProfile) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfDB) DeleteProfileByID(ctx context.Context, id string) error {
	return errNotImplemented
}

func (*phProfDB) BeginTx(ctx context.Context) (IProfileTx, error) {
	return &phProfTx{}, nil
}

type phProfTx struct{}

func (*phProfTx) Commit(ctx context.Context) error {
	return errNotImplemented
}

func (*phProfTx) Rollback(ctx context.Context) error {
	return errNotImplemented
}

func (*phProfTx) CreateProfile(ctx context.Context, acc *ParamCreateProfile) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfTx) UpdateProfile(ctx context.Context, acc *ParamUpdateProfile) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfTx) PatchProfile(ctx context.Context, acc *ParamPatchProfile) (*Profile, error) {
	return nil, errNotImplemented
}

func (*phProfTx) DeleteProfileByID(ctx context.Context, id string) error {
	return errNotImplemented
}
