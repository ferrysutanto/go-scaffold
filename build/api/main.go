package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/build/commons/servers"
	"go.opentelemetry.io/otel"

	log "github.com/sirupsen/logrus"
)

var (
	// server config
	host    = "localhost"
	port    = 8080
	appName = "go-scaffold"
)

func main() {
	// 1. init tracer and start a span and defer its closure
	ctx, span := otel.Tracer("").Start(context.Background(), "[api][main]")
	defer span.End()

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
		err = errors.WrapWithCode(err, fmt.Sprintf("[%s] failed to create server", appName), 500)
		span.RecordError(err)
		log.Fatal(err)
	}

	// Run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			err = errors.WrapWithCode(err, fmt.Sprintf("[%s] failed to listen and serve", appName), 500)
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
