package pg

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, config *Config) (db.IGenericDB, error) {
	write, read, err := initConnection(config)
	if err != nil {
		return nil, err
	}

	config.PrimaryDB = write.DB
	config.ReplicaDB = read.DB

	accRepo, err := newAccountRepository(ctx, config)
	if err != nil {
		return nil, err
	}

	profRepo, err := newProfileRepository(ctx, config)
	if err != nil {
		return nil, err
	}

	return &PG{
		prim:    write,
		account: accRepo,
		profile: profRepo,
	}, nil
}

type PG struct {
	prim *sqlx.DB
	repl *sqlx.DB

	account *AccountRepository
	profile *ProfileRepository
}

func (this *PG) Ping(ctx context.Context) error {
	if err := this.prim.PingContext(ctx); err != nil {
		return err
	}

	if err := this.repl.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (this *PG) BeginTx(ctx context.Context) (db.ITx, error) {
	tx, err := this.prim.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	accountTx, err := this.account.beginTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	profileTx, err := this.profile.beginTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &PgTx{
		tx:        tx,
		accountTx: accountTx,
		profileTx: profileTx,
	}, nil
}

func (this *PG) Account(ctx context.Context) db.IAccountRepository {
	return this.account
}

func (this *PG) Profile(ctx context.Context) db.IProfileRepository {
	return this.profile
}
