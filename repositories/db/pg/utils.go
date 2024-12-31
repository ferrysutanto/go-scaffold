package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	PrimaryDB *sql.DB

	Host     *string
	Port     *uint
	Username *string
	Password *string
	Database *string
	SslMode  *string

	MaxIdleConns *int
	MaxOpenConns *int

	ReplicaDB *sql.DB

	ReplicaHost     *string
	ReplicaPort     *uint
	ReplicaUsername *string
	ReplicaPassword *string
	ReplicaDatabase *string
	ReplicaSslMode  *string

	ReplicaMaxIdleConns *int
	ReplicaMaxOpenConns *int
}

func validateConfig(config *Config) error {
	if config.PrimaryDB == nil && (config.Host == nil || config.Port == nil || config.Username == nil || config.Password == nil || config.Database == nil) {
		return errors.New("main database configuration is missing required fields")
	}

	if config.ReplicaHost != nil || config.ReplicaPort != nil || config.ReplicaUsername != nil || config.ReplicaPassword != nil || config.ReplicaDatabase != nil {
		missingFields := []string{}
		if config.ReplicaHost == nil {
			missingFields = append(missingFields, "ReplicaHost")
		}
		if config.ReplicaPort == nil {
			missingFields = append(missingFields, "ReplicaPort")
		}
		if config.ReplicaUsername == nil {
			missingFields = append(missingFields, "ReplicaUsername")
		}
		if config.ReplicaPassword == nil {
			missingFields = append(missingFields, "ReplicaPassword")
		}
		if config.ReplicaDatabase == nil {
			missingFields = append(missingFields, "ReplicaDatabase")
		}

		if len(missingFields) > 0 && len(missingFields) < 5 {
			return errors.New("replica database configuration is incomplete, missing fields: " + strings.Join(missingFields, ", "))
		}
	}

	return nil
}

func initConnection(config *Config) (write *sqlx.DB, read *sqlx.DB, err error) {
	// if err := validateConfig(config); err != nil {
	// 	return nil, nil, err
	// }

	if config.PrimaryDB != nil {
		write = sqlx.NewDb(config.PrimaryDB, "postgres")
	} else {
		// init write connection
		write, err = sqlx.Connect("postgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
				*config.Host, *config.Port, *config.Username, *config.Password, *config.Database, *config.SslMode))
		if err != nil {
			return nil, nil, err
		}

		if config.MaxIdleConns != nil {
			write.SetMaxIdleConns(*config.MaxIdleConns)
		}

		if config.MaxOpenConns != nil {
			write.SetMaxOpenConns(*config.MaxOpenConns)
		}
	}

	// by default, fallback replica to primary
	read = write

	// now see if we have replica configuration supplied
	// if so, override the read connection
	if config.ReplicaDB != nil {
		read = sqlx.NewDb(config.ReplicaDB, "postgres")
	} else if config.ReplicaHost != nil {
		// or create a new connection
		read, err = sqlx.Connect("postgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
				*config.ReplicaHost, *config.ReplicaPort, *config.ReplicaUsername, *config.ReplicaPassword, *config.ReplicaDatabase, *config.ReplicaSslMode))
		if err != nil {
			return nil, nil, err
		}

		if config.ReplicaMaxIdleConns != nil {
			read.SetMaxIdleConns(*config.ReplicaMaxIdleConns)
		}

		if config.ReplicaMaxOpenConns != nil {
			read.SetMaxOpenConns(*config.ReplicaMaxOpenConns)
		}
	}

	return write, read, nil
}
