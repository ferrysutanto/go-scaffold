package pg

import (
	"context"
	"fmt"

	"github.com/ferrysutanto/go-errors"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"

	_ "github.com/lib/pq"
)

func instantiateDbFromCfg(ctx context.Context, conf *Config) (main, replica *sqlx.DB, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][New]")
	defer span.End()

	sslMode := "disable"
	if conf.SslMode != nil {
		sslMode = *conf.SslMode
	}
	ds := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", *conf.Host, *conf.Port, *conf.Username, *conf.Password, *conf.Name, sslMode)
	main, err = sqlx.ConnectContext(ctx, "postgres", ds)
	if err != nil {
		err = errors.WrapWithCode(err, "failed to open database connection", 500)
		span.RecordError(err)

		return nil, nil, err
	}

	replica = main // by default, replica is the same as main

	// if replica is provided, open a dedicated replica connection
	if conf.ReplicaHost != nil {
		sslMode := "disable"
		if conf.ReplicaSslMode != nil {
			sslMode = *conf.ReplicaSslMode
		}
		replicaDS := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", *conf.ReplicaHost, *conf.ReplicaPort, *conf.ReplicaUsername, *conf.ReplicaPassword, *conf.ReplicaName, sslMode)

		replica, err = sqlx.ConnectContext(ctx, "postgres", replicaDS)
		if err != nil {
			err = errors.WrapWithCode(err, "failed to open replica database connection", 500)
			span.RecordError(err)

			return nil, nil, err
		}
	}

	return main, replica, nil
}
