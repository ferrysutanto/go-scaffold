package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handler "github.com/ferrysutanto/go-scaffold/build/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type dbConfig struct {
	driver   string        `env:"DB_DRIVER" envDefault:"postgres" validate:"required"`
	host     string        `env:"DB_HOST" envDefault:"localhost" validate:"required"`
	port     string        `env:"DB_PORT" envDefault:"5432" validate:"required"`
	username string        `env:"DB_USERNAME" envDefault:"postgres" validate:"required"`
	password string        `env:"DB_PASSWORD" envDefault:"postgres" validate:"required"`
	name     string        `env:"DB_NAME" envDefault:"postgres" validate:"required"`
	timeout  time.Duration `env:"DB_TIMEOUT" envDefault:"5s" validate:"required"`
	sslMode  string        `env:"DB_SSL_MODE" envDefault:"disable" validate:"required"`
}

var (
	hdlr handler.Handler

	// db configs
	mainDbConfig    dbConfig
	replicaDbConfig dbConfig

	// shutdown timeout
	shutdownTimeout = 5 * time.Second

	// server config
	env  = "development"
	host = "localhost"
	port = ":8080"

	// server shutdown
	shutdownChan = make(chan os.Signal, 1)
	// server shutdown timeout
	shutdownTimeoutChan = make(chan bool, 1)
	// server shutdown done
	shutdownDoneChan = make(chan bool, 1)
	// server shutdown error
	shutdownErrorChan = make(chan error, 1)

	vldtr = validator.New()
)

func validate(ctx context.Context) error {
	// validate db configs
	if err := vldtr.StructCtx(ctx, mainDbConfig); err != nil {
		return fmt.Errorf("[500] failed to validate main db config: %v", err)
	}

	if err := vldtr.StructCtx(ctx, replicaDbConfig); err != nil {
		return fmt.Errorf("[500] failed to validate replica db config: %v", err)
	}

	return nil
}

func mapOsEnvToVariables() {
	// map os env to variables
	mainDbConfig.driver = os.Getenv("DB_DRIVER")
	mainDbConfig.host = os.Getenv("DB_HOST")
	mainDbConfig.port = os.Getenv("DB_PORT")
	mainDbConfig.username = os.Getenv("DB_USERNAME")
	mainDbConfig.password = os.Getenv("DB_PASSWORD")
	mainDbConfig.name = os.Getenv("DB_NAME")
	mainDbConfig.timeout = 5 * time.Second
	mainDbConfig.sslMode = os.Getenv("DB_SSL_MODE")

	replicaDbConfig.host = os.Getenv("DB_HOST")
	replicaDbConfig.port = os.Getenv("DB_PORT")
	replicaDbConfig.username = os.Getenv("DB_USERNAME")
	replicaDbConfig.password = os.Getenv("DB_PASSWORD")
	replicaDbConfig.name = os.Getenv("DB_NAME")
	replicaDbConfig.timeout = 5 * time.Second
	replicaDbConfig.sslMode = os.Getenv("DB_SSL_MODE")
}

func main() {
	godotenv.Load()

	mapOsEnvToVariables()

	if err := validate(context.Background()); err != nil {
		log.Fatal(err)
	}

	flagHost := flag.String("host", "", "host for the account api server")
	flagPort := flag.String("port", "", "port for the account api server")
	flag.Parse()

	if *flagHost != "" {
		host = *flagHost
	}

	if *flagPort != "" {
		port = *flagPort
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	handlerConfig := handler.Config{
		Type:                  handler.TypeBasic,
		DbDriverName:          mainDbConfig.driver,
		DbHost:                mainDbConfig.host,
		DbPort:                mainDbConfig.port,
		DbUsername:            mainDbConfig.username,
		DbPassword:            mainDbConfig.password,
		DbName:                mainDbConfig.name,
		DbSSLMode:             mainDbConfig.sslMode,
		ReplicationDbHost:     replicaDbConfig.host,
		ReplicationDbPort:     replicaDbConfig.port,
		ReplicationDbUsername: replicaDbConfig.username,
		ReplicationDbPassword: replicaDbConfig.password,
		ReplicationDbName:     replicaDbConfig.name,
		ReplicationDbSSLMode:  replicaDbConfig.sslMode,
	}

	if err := handler.InitDefaultHandler(ctx, handlerConfig); err != nil {
		fmt.Printf("%+v\n", err)
		log.Println("====================================")
		log.Fatal(err)
	}

	r := gin.Default()
	// r.Use(middleware.CircuitBreakerMiddleware())

	r.GET("/ping", handler.Healthcheck)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s%s", host, port),
		Handler: r,
	}

	// Run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[go-app] listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println(("[go-app] Shutdown Server..."))

	offCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(offCtx); err != nil {
		log.Fatalf("[go-app] Server Shutdown: %v", err)
	}

	log.Println("[go-app] Server exiting")

}
