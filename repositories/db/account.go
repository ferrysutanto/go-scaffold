package db

import (
	"context"
	"time"
)

type IAccountDB interface {
	GetAccounts(ctx context.Context, limit, offset int) (*Accounts, error)
	FindAccountByID(ctx context.Context, id string) (*Account, error)
	CreateAccount(ctx context.Context, account *ParamCreateAccount) (*Account, error)
	UpdateAccount(ctx context.Context, account *ParamUpdateAccount) (*Account, error)
	PatchAccount(ctx context.Context, account *ParamPatchAccount) (*Account, error)
	DeleteAccountByID(ctx context.Context, id string) error

	BeginTx(ctx context.Context) (IAccountTx, error)
}

type IAccountTx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	CreateAccount(ctx context.Context, account *ParamCreateAccount) (*Account, error)
	UpdateAccount(ctx context.Context, account *ParamUpdateAccount) (*Account, error)
	PatchAccount(ctx context.Context, account *ParamPatchAccount) (*Account, error)
	DeleteAccountByID(ctx context.Context, id string) error
}

type Accounts struct {
	Data     []*Account
	Metadata *Metadata
}

type Account struct {
	ID         string
	FirstName  string
	LastName   string
	Email      string
	Sex        *string
	Birthdate  time.Time
	PictureURL *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ParamCreateAccount struct {
	ID         *string
	FirstName  string
	LastName   string
	Email      string
	Sex        *string
	Birthdate  time.Time
	PictureURL *string
}

type ParamUpdateAccount struct {
	ID         string
	FirstName  string
	LastName   string
	Email      string
	Sex        *string
	Birthdate  time.Time
	PictureURL *string
}

type ParamPatchAccount struct {
	ID         string
	FirstName  *string
	LastName   *string
	Email      *string
	Sex        *string
	Birthdate  *time.Time
	PictureURL *string
}
