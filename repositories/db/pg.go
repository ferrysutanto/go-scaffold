package db

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type pgModel struct {
	db     *sqlx.DB
	replDb *sqlx.DB
}

func newPgModel(ctx context.Context, conf *Config) (*pgModel, error) {
	// 1. check if context is provided
	if ctx == nil {
		return nil, errors.New("[repositories/db][newPgModel] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][newPgModel]")
	defer span.End()

	// 3. define main db
	var db *sqlx.DB

	// 4.a. check if main db object is provided
	if conf.DB != nil {
		// 4.a.1. cast main db object to sqlx.DB for more convenient usage
		db = sqlx.NewDb(conf.DB, "postgres")
		// 4.a.2. ping main db to check if connection is established
		if err := db.PingContext(ctx); err != nil {
			// wrap error with additional info
			err = errors.Wrap(err, "[repositories/db][newPgModel] failed to ping main db")
			// record error in span
			span.RecordError(err)

			return nil, err
		}
	} else { // 4.b. if main db object is not provided
		var err error
		// 4.b.1. open main db connection
		db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%v", conf.DbHost, conf.DbPort, conf.DbUsername, conf.DbPassword, conf.DbName, conf.DbSslMode))
		if err != nil {
			// wrap error with additional info
			err = errors.Wrap(err, "[repositories/db][newPgModel] failed to open main db")
			// record error in span
			span.RecordError(err)

			return nil, err
		}
	}

	// 5. define replica db
	var replDb *sqlx.DB

	// 6.a. check if replica db object is provided
	if conf.ReplicaDB != nil {
		// 6.a.1. cast replica db object to sqlx.DB for more convenient usage
		replDb = sqlx.NewDb(conf.ReplicaDB, "postgres")
		// 6.a.2. ping replica db to check if connection is established
		if err := replDb.Ping(); err != nil {
			// wrap error with additional info
			err = errors.Wrap(err, "[repositories/db][newPgModel] failed to ping replica db")
			// record error in span
			span.RecordError(err)

			return nil, err
		}
	} else { // 6.b. if replica db object is not provided
		var err error
		// 6.b.1. open replica db connection
		replDb, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%v", conf.ReplicaDbHost, conf.ReplicaDbPort, conf.ReplicaDbUsername, conf.ReplicaDbPassword, conf.ReplicaDbName, conf.ReplicaDbSslMode))
		if err != nil {
			// wrap error with additional info
			err = errors.Wrap(err, "[repositories/db][newPgModel] failed to open replica db")
			// record error in span
			span.RecordError(err)

			return nil, err
		}
	}

	// 7. declare new pgModel object
	resp := &pgModel{
		db:     db,
		replDb: replDb,
	}

	// 8. return pgModel object and nil error
	return resp, nil
}
