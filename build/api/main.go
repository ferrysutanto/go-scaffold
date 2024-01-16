package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ferrysutanto/go-scaffold/build/api/servers"
	"github.com/ferrysutanto/go-scaffold/services"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"

	log "github.com/sirupsen/logrus"
)

var (
	// server config
	host    = "localhost"
	port    = 8080
	appName = "go-scaffold"
)

func init() {
	ctx := context.Background()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("[%s] No .env file found...", appName)
	}

	if envAppName := os.Getenv("APP_NAME"); envAppName != "" {
		appName = envAppName
	}

	if appHost := os.Getenv("APP_HOST"); appHost != "" {
		host = appHost
	}

	if appPort := os.Getenv("APP_PORT"); appPort != "" {
		p, err := strconv.Atoi(appPort)
		if err != nil {
			log.Fatalf("[%s] failed to parse APP_PORT", appName)
		}
		port = p
	}

	if err := services.Init(ctx); err != nil {
		err = errors.Wrapf(err, "[%s] failed to init services", appName)
		log.Fatalln(err)
	}
}

func main() {
	// 1. init tracer and start a span and defer its closure
	ctx, span := otel.Tracer("").Start(context.Background(), "[api][main]")
	defer span.End()

	// 2. load .env file
	span.AddEvent("load .env file")
	if err := godotenv.Load(); err != nil {
		log.Printf("[%s] No .env file found...", appName)
	}

	// 3. parse flags
	// flagHost := ""
	// flag.StringVar(&flagHost, "host", "", "host for the account api server")
	span.AddEvent("parse flags")
	if flagHost := flag.String("host", "", "host for the account api server"); flagHost != nil && *flagHost != "" {
		host = *flagHost
	}
	if flagPort := flag.Int("port", 0, "port for the account api server"); flagPort != nil && *flagPort != 0 {
		port = *flagPort
	}
	flag.Parse()

	log.Printf("[%s] Starting Server at %s:%d", appName, host, port)
	// 4. create server
	srv, err := servers.New(ctx, &servers.Config{
		Host: host,
		Port: port,
	})
	if err != nil {
		err = errors.Wrapf(err, "[%s] failed to create server", appName)
		span.RecordError(err)
		log.Fatal(err)
	}

	// Run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			err = errors.Wrapf(err, "[%s] failed to listen and serve", appName)
			span.RecordError(err)
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("[%s] Shutting down Server...", appName)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[%s] Server Shutdown: %v", appName, err)
	}

	log.Printf("[%s] Server exiting", appName)
}
