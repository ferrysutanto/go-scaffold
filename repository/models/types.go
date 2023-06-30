package models

import "database/sql"

type Models interface {
	Hello() error
}

type Config struct {
	// DriverName is the name of the database driver and it's mandatory
	DriverName string `json:"driver_name" yaml:"driver_name" env:"DB_DRIVER_NAME" env-default:"postgres" validate:"required"`

	// You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	DB *sql.DB `validate:"required_without=DbHost"`

	DbHost     string `json:"db_host" yaml:"db_host" env:"DB_HOST" env-default:"localhost" validate:"hostname|ip,required_without=DB,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	DbPort     string `json:"db_port" yaml:"db_port" env:"DB_PORT" env-default:"5432" validate:"numeric,required_without=DB,required_with=DbHost,required_with=DbName,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	DbName     string `json:"db_name" yaml:"db_name" env:"DB_NAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbUsername,required_with=DbPassword,required_with=DbSSLMode"`
	DbUsername string `json:"db_username" yaml:"db_username" env:"DB_USERNAME" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbPassword,required_with=DbSSLMode"`
	DbPassword string `json:"db_password" yaml:"db_password" env:"DB_PASSWORD" env-default:"postgres" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbSSLMode"`
	DbSSLMode  string `json:"db_ssl_mode" yaml:"db_ssl_mode" env:"DB_SSL_MODE" env-default:"disable" validate:"required_without=DB,required_with=DbHost,required_with=DbPort,required_with=DbName,required_with=DbUsername,required_with=DbPassword"`

	// ReplicationDB is the database connection for replication, it's optional
	ReplicationDB         *sql.DB `validate:"required_without=ReplicationDbHost"`
	ReplicationDbHost     string  `json:"replication_db_host" yaml:"replication_db_host" env:"REPLICATION_DB_HOST" env-default:"localhost" validate:"hostname|ip,required_without=ReplicationDB,required_with=ReplicationDbPort,required_with=ReplicationDbName,required_with=ReplicationDbUsername,required_with=ReplicationDbPassword,required_with=ReplicationDbSSLMode"`
	ReplicationDbPort     string  `json:"replication_db_port" yaml:"replication_db_port" env:"REPLICATION_DB_PORT" env-default:"5432" validate:"numeric,required_without=ReplicationDB,required_with=ReplicationDbHost,required_with=ReplicationDbName,required_with=ReplicationDbUsername,required_with=ReplicationDbPassword,required_with=ReplicationDbSSLMode"`
	ReplicationDbName     string  `json:"replication_db_name" yaml:"replication_db_name" env:"REPLICATION_DB_NAME" env-default:"postgres" validate:"required_without=ReplicationDB,required_with=ReplicationDbHost,required_with=ReplicationDbPort,required_with=ReplicationDbUsername,required_with=ReplicationDbPassword,required_with=ReplicationDbSSLMode"`
	ReplicationDbUsername string  `json:"replication_db_username" yaml:"replication_db_username" env:"REPLICATION_DB_USERNAME" env-default:"postgres" validate:"required_without=ReplicationDB,required_with=ReplicationDbHost,required_with=ReplicationDbPort,required_with=ReplicationDbName,required_with=ReplicationDbPassword,required_with=ReplicationDbSSLMode"`
	ReplicationDbPassword string  `json:"replication_db_password" yaml:"replication_db_password" env:"REPLICATION_DB_PASSWORD" env-default:"postgres" validate:"required_without=ReplicationDB,required_with=ReplicationDbHost,required_with=ReplicationDbPort,required_with=ReplicationDbName,required_with=ReplicationDbUsername,required_with=ReplicationDbSSLMode"`
	ReplicationDbSSLMode  string  `json:"replication_db_ssl_mode" yaml:"replication_db_ssl_mode" env:"REPLICATION_DB_SSL_MODE" env-default:"disable" validate:"required_without=ReplicationDB,required_with=ReplicationDbHost,required_with=ReplicationDbPort,required_with=ReplicationDbName,required_with=ReplicationDbUsername,required_with=ReplicationDbPassword"`
}
