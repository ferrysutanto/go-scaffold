package db

import (
	"context"
	"time"
)

type IAccountRepository interface {
	GetAccounts(ctx context.Context, param *ParamGetAccounts) (*Accounts, error)
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
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     *string   `db:"email"`
	Phone     *string   `db:"phone"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ParamGetAccounts struct {
}

type ParamCreateAccount struct {
	ID       *string
	Username string
	Email    *string
	Phone    *string
	Status   *string
}

type ParamUpdateAccount struct {
	ID       string
	Username string
	Email    *string
	Phone    *string
	Status   string
}

type ParamPatchAccount struct {
	ID       string
	Username *string
	Email    *string
	Phone    *string
	Status   *string
}
