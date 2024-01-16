package cache

import "context"

// Cache is an interface that defines cache operations
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

// Config is a struct that defines cache config
type Config struct {
	DriverName string `json:"driver_name" yaml:"driver_name"`
	Host       string `json:"host" yaml:"host"`
	Port       int    `json:"port" yaml:"port"`
	Username   string `json:"username" yaml:"username"`
	Password   string `json:"password" yaml:"password"`
	DB         int    `json:"db" yaml:"db"`
}
