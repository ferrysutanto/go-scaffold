package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ferrysutanto/go-scaffold/services"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Go scaffold is a simple scaffold for Go project",
	Long: `Go scaffold is a simple scaffold for Go project.
This project is intended to be used as a starting point for a new Go project.
It provides a basic structure and a few basic functionalities.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Healthcheck",
	Long:  `Healthcheck the database connection`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := services.Healthcheck(context.Background()); err != nil {
			fmt.Printf("[go-scaffold][healthcheckCmd][Run] failed to healthcheck: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("[go-scaffold][healthcheckCmd][Run] healthcheck success\n")
	},
}

func init() {
	godotenv.Load()

	// Init services default
	if err := services.Init(context.Background(), &services.Config{
		Type: services.BASIC_SERVICE,
		DB: &services.DbObject{
			DriverName:            os.Getenv("DB_DRIVER"),
			DbHost:                os.Getenv("DB_HOST"),
			DbPort:                os.Getenv("DB_PORT"),
			DbName:                os.Getenv("DB_NAME"),
			DbUsername:            os.Getenv("DB_USERNAME"),
			DbPassword:            os.Getenv("DB_PASSWORD"),
			DbSSLMode:             os.Getenv("DB_SSL_MODE"),
			ReplicationDbHost:     os.Getenv("DB_REPLICATION_HOST"),
			ReplicationDbPort:     os.Getenv("DB_REPLICATION_PORT"),
			ReplicationDbName:     os.Getenv("DB_REPLICATION_NAME"),
			ReplicationDbUsername: os.Getenv("DB_REPLICATION_USERNAME"),
			ReplicationDbPassword: os.Getenv("DB_REPLICATION_PASSWORD"),
			ReplicationDbSSLMode:  os.Getenv("DB_REPLICATION_SSL_MODE"),
		},
	}); err != nil {
		fmt.Printf("[go-scaffold][init] failed to init services: %v\n", err)
	}

	// Add healthcheck command
	rootCmd.AddCommand(healthcheckCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
