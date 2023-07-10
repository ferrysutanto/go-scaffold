package models

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type pgModel struct {
	db     *sqlx.DB
	replDb *sqlx.DB
	v      *validator.Validate
}

func newPgModel(ctx context.Context, conf *Config) (*pgModel, error) {
	var db *sqlx.DB
	if conf.DB != nil {
		db = sqlx.NewDb(conf.DB, "postgres")
		if err := db.Ping(); err != nil {
			return nil, errors.Wrap(err, "[models][newPgModel] failed to ping db")
		}
	} else {
		var err error
		db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%v", conf.DbHost, conf.DbPort, conf.DbUsername, conf.DbPassword, conf.DbName, conf.DbSSLMode))
		if err != nil {
			log.Errorf("failed to open database: %v", err)
			return nil, errors.Wrap(err, "[models][newPgModel] failed to open db")
		}
	}

	var replDb *sqlx.DB
	if conf.ReplicationDB != nil {
		replDb = sqlx.NewDb(conf.ReplicationDB, "postgres")
		if err := replDb.Ping(); err != nil {
			return nil, errors.Wrap(err, "[models][newPgModel] failed to ping replication db")
		}
	} else {
		var err error
		replDb, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%v", conf.ReplicationDbHost, conf.ReplicationDbPort, conf.ReplicationDbUsername, conf.ReplicationDbPassword, conf.ReplicationDbName, conf.ReplicationDbSSLMode))
		if err != nil {
			return nil, errors.Wrap(err, "[models][newPgModel] failed to open replication db")
		}
	}

	resp := &pgModel{
		db:     db,
		replDb: replDb,
		v:      validator.New(),
	}

	return resp, nil
}
