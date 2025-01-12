package db

import (
	"context"
	"time"
)

type IProfileRepository interface {
	GetProfiles(ctx context.Context, param *ParamGetProfiles) (*Profiles, error)
	FindProfileByID(ctx context.Context, id string) (*Profile, error)
	CreateProfile(ctx context.Context, param *ParamCreateProfile) (*Profile, error)
	UpdateProfile(ctx context.Context, param *ParamUpdateProfile) (*Profile, error)
	PatchProfile(ctx context.Context, param *ParamPatchProfile) error
	DeleteProfileByID(ctx context.Context, id string) error

	BeginTx(ctx context.Context) (IProfileTx, error)
}

type IProfileTx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	CreateProfile(ctx context.Context, param *ParamCreateProfile) (*Profile, error)
	UpdateProfile(ctx context.Context, param *ParamUpdateProfile) (*Profile, error)
	PatchProfile(ctx context.Context, param *ParamPatchProfile) error
	DeleteProfileByID(ctx context.Context, id string) error
}

type Profiles struct {
	Data     []*Profile
	Metadata *Metadata
}

type Profile struct {
	ID         string
	AccountID  string
	FirstName  string
	LastName   string
	Birthdate  *time.Time
	Sex        *string
	PictureURL *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ParamGetProfiles struct {
}

type ParamCreateProfile struct {
	ID         string
	AccountID  string
	FirstName  string
	LastName   string
	Birthdate  *time.Time
	Sex        *string
	PictureURL *string
}

type ParamUpdateProfile struct {
	ID         string
	AccountID  string
	FirstName  string
	LastName   string
	Birthdate  *time.Time
	Sex        *string
	PictureURL *string
}

type ParamPatchProfile struct {
	ID         string
	AccountID  *string
	FirstName  *string
	LastName   *string
	Birthdate  *time.Time
	Sex        *string
	PictureURL *string
}
