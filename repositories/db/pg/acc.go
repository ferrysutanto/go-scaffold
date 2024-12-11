package pg

import (
	"context"

	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
)

type accountDB struct {
	mainDB    *sqlx.DB
	replicaDB *sqlx.DB

	fetchAccountByIdStmt  *sqlx.Stmt
	createAccountStmt     *sqlx.Stmt
	updateAccountStmt     *sqlx.Stmt
	patchAccountStmt      *sqlx.Stmt
	deleteAccountByIdStmt *sqlx.Stmt
}

func NewAccountDB(ctx context.Context, cfg *Config) (db.IAccountDB, error) {
	if err := validateConfig(ctx, cfg); err != nil {
		return nil, errors.WrapWithCode(err, "bad config", 400)
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][NewAccountDB]")
	defer span.End()

	var (
		err       error
		mainDB    *sqlx.DB
		replicaDB *sqlx.DB
	)

	// 4.a. check if main db object is provided
	if cfg.DB != nil {
		mainDB = cfg.DB
		replicaDB = cfg.DB

		if cfg.ReplicaDB != nil {
			replicaDB = cfg.ReplicaDB
		}
	} else { // 4.b. if main db object is not provided
		mainDB, replicaDB, err = instantiateDbFromCfg(ctx, cfg)
		if err != nil {
			return nil, err
		}
	}

	// 7. declare new pgModel object
	resp := accountDB{
		mainDB:    mainDB,
		replicaDB: replicaDB,
	}

	// 8. return pgModel object and nil error
	return &resp, nil
}
