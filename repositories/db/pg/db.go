package pg

import (
	"context"
	"strings"

	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
)

type pgDB struct {
	mainDB    *sqlx.DB
	replicaDB *sqlx.DB

	/* ======================
	put sql statements here
	=========================
	createStmt *sqlx.Stmt
	fetchByIdStmt *sqlx.Stmt
	========================= */
}

type Config struct {
	DB       *sqlx.DB // You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	Host     *string
	Port     *uint
	Name     *string
	Username *string
	Password *string
	SslMode  *bool // SSLMode is the ssl mode of the database connection, default is disable

	ReplicaDB       *sqlx.DB // ReplicaDB is the database connection for replication, it's optional, if it's not provided, then it will use the same connection as the main database
	ReplicaHost     *string
	ReplicaPort     *uint
	ReplicaName     *string
	ReplicaUsername *string
	ReplicaPassword *string
	ReplicaSslMode  *bool
}

func validateConfig(ctx context.Context, cfg *Config) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	// 2. check if config is provided
	if cfg == nil {
		return errors.NewWithCode("config is required", 400)
	}

	// 3. validate config
	if cfg.DB != nil && cfg.Host != nil {
		return errors.NewWithCode("db instance and host cannot be provided together", 400)
	}

	if cfg.DB == nil && cfg.Host == nil {
		return errors.NewWithCode("db instance or host is required", 400)
	}

	var errs = make([]string, 0)

	if cfg.DB == nil {
		if cfg.Host == nil {
			errs = append(errs, "host is required")
		}

		if cfg.Port == nil {
			errs = append(errs, "port is required")
		}

		if cfg.Username == nil {
			errs = append(errs, "user is required")
		}

		if cfg.Password == nil {
			errs = append(errs, "password is required")
		}

		if cfg.Name == nil {
			errs = append(errs, "database is required")
		}
	}

	if len(errs) > 0 {
		return errors.NewWithCode(strings.Join(errs, ", "), 400)
	}

	return nil
}

func New(ctx context.Context, cfg *Config) (db.IDB, error) {
	if err := validateConfig(ctx, cfg); err != nil {
		return nil, errors.WrapWithCode(err, "bad config", 400)
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][New]")
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
	resp := pgDB{
		mainDB,
		replicaDB,
	}

	// 8. return pgModel object and nil error
	return &resp, nil
}

func (m *pgDB) Ping(ctx context.Context) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/db][pgModel:Ping] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][pgModel:Ping]")
	defer span.End()

	// 3. ping main db
	if err := m.mainDB.PingContext(ctx); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][pgModel:Ping] failed to ping db")
		// record error in span
		span.RecordError(err)

		return err
	}

	// 4. ping replica db, same as above but for replica
	if err := m.replicaDB.PingContext(ctx); err != nil {
		// wrap error with additional info
		err = errors.Wrap(err, "[repositories/db][pgModel:Ping] failed to ping db replica")
		// record error in span
		span.RecordError(err)

		return err
	}

	return nil
}
