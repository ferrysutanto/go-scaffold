package db

import (
	"context"
	"database/sql"
)

type DB interface {
	Ping(context.Context) error
}

type Config struct {
	// DriverName is the name of the database driver and it's mandatory
	DriverName string `json:"driver_name" yaml:"driver_name" env:"DB_DRIVER" env-default:"postgres" validate:"required"`

	// You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	DB         *sql.DB `validate:"required_without=DbHost"`
	DbHost     string  `json:"db_host" yaml:"db_host" env:"DB_HOST" env-default:"localhost" validate:"hostname|ip,required_without=DB,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	DbPort     int     `json:"db_port" yaml:"db_port" env:"DB_PORT" env-default:"5432" validate:"numeric,required_without=DB,required_with=DbHost,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	DbName     string  `json:"db_name" yaml:"db_name" env:"DB_NAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	DbUsername string  `json:"db_username" yaml:"db_username" env:"DB_USERNAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbPassword,required_with=DbSSLMode"`
	DbPassword string  `json:"db_password" yaml:"db_password" env:"DB_PASSWORD" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbSSLMode"`
	DbSslMode  string  `json:"db_ssl_mode" yaml:"db_ssl_mode" env:"DB_SSL_MODE" env-default:"disable" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword"`

	// ReplicaDB is the database connection for replication, it's optional
	ReplicaDB         *sql.DB `validate:"required_without=ReplicaDbHost"`
	ReplicaDbHost     string  `json:"replica_db_host" yaml:"replica_db_host" env:"REPLICATION_DB_HOST" env-default:"localhost" validate:"hostname|ip,required_without=ReplicaDB,required_with=ReplicaDbPort,required_with=ReplicaDbName,required_with=ReplicaDbUsername,required_with=ReplicaDbPassword,required_with=ReplicaDbSSLMode"`
	ReplicaDbPort     int     `json:"replica_db_port" yaml:"replica_db_port" env:"REPLICATION_DB_PORT" env-default:"5432" validate:"numeric,required_without=ReplicaDB,required_with=ReplicaDbHost,required_with=ReplicaDbName,required_with=ReplicaDbUsername,required_with=ReplicaDbPassword,required_with=ReplicaDbSSLMode"`
	ReplicaDbName     string  `json:"replica_db_name" yaml:"replica_db_name" env:"REPLICATION_DB_NAME" env-default:"postgres" validate:"required_without=ReplicaDB,required_with=ReplicaDbHost,required_with=ReplicaDbPort,required_with=ReplicaDbUsername,required_with=ReplicaDbPassword,required_with=ReplicaDbSSLMode"`
	ReplicaDbUsername string  `json:"replica_db_username" yaml:"replica_db_username" env:"REPLICATION_DB_USERNAME" env-default:"postgres" validate:"required_without=ReplicaDB,required_with=ReplicaDbHost,required_with=ReplicaDbPort,required_with=ReplicaDbName,required_with=ReplicaDbPassword,required_with=ReplicaDbSSLMode"`
	ReplicaDbPassword string  `json:"replica_db_password" yaml:"replica_db_password" env:"REPLICATION_DB_PASSWORD" env-default:"postgres" validate:"required_without=ReplicaDB,required_with=ReplicaDbHost,required_with=ReplicaDbPort,required_with=ReplicaDbName,required_with=ReplicaDbUsername,required_with=ReplicaDbSSLMode"`
	ReplicaDbSslMode  string  `json:"replica_db_ssl_mode" yaml:"replica_db_ssl_mode" env:"REPLICATION_DB_SSL_MODE" env-default:"disable" validate:"required_without=ReplicaDB,required_with=ReplicaDbHost,required_with=ReplicaDbPort,required_with=ReplicaDbName,required_with=ReplicaDbUsername,required_with=ReplicaDbPassword"`
}
