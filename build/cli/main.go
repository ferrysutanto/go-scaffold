package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ferrysutanto/go-scaffold/services"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
)

var (
	appName = "go-scaffold"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("[%s] No .env file found...", appName)
	}

	envAppName := os.Getenv("APP_NAME")
	if envAppName != "" {
		appName = envAppName
	}
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Go scaffold is a simple scaffold for Go project",
	Long: `Go scaffold is a simple scaffold for Go project.
This project is intended to be used as a starting point for a new Go project.
It provides a basic structure and a few basic functionalities.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. declare context
		ctx := cmd.Context()

		// 2. start a span and defer its closure
		_, span := otel.Tracer("").Start(ctx, "[cli][rootCmd]")
		defer span.End()

		// 3. execute the service
		if err := cmd.Help(); err != nil {
			err = errors.Wrapf(err, "[%s] failed to run rootCmd", appName)
			span.RecordError(err)
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Healthcheck",
	Long:  `Healthcheck do overall checking for application readiness to accept requests`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. declare context
		ctx := cmd.Context()

		// 2. start a span and defer its closure
		ctx, span := otel.Tracer("").Start(ctx, "[cli][healthcheckCmd]")
		defer span.End()

		// 3. execute the service
		if err := services.Healthcheck(ctx); err != nil {
			fmt.Printf("[%s] failed to healthcheck: %v\n", appName, err)
			os.Exit(1)
		}

		fmt.Printf("[%s] healthcheck success\n", appName)
	},
}

func init() {
	ctx := context.Background()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("[%s] No .env file found...", appName)
	}

	// Init services default
	if err := services.Init(ctx); err != nil {
		fmt.Printf("[%s] failed to init services: %v\n", appName, err)
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
