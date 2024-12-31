package utils

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *EnvironmentVariables
		wantErr bool
	}{
		{
			name: "Test GetEnv",
			args: args{
				ctx: context.Background(),
			},
			want: &EnvironmentVariables{
				DB: &dbConfig{
					Driver: "postgres",
					PG: &pgDBConfig{
						Host:            "maybesomewhere",
						Port:            9859,
						User:            "mustbe",
						Password:        "secret",
						Database:        "123456",
						SslMode:         "disable",
						ReplicaHost:     aws.String("replicahost"),
						ReplicaPort:     aws.Uint(8999),
						ReplicaUser:     aws.String("fakeuser"),
						ReplicaPassword: aws.String("fakepass"),
						ReplicaDatabase: aws.String("fakedb"),
						ReplicaSslMode:  aws.String("disable"),
					},
				},
				Cache: &cacheConfig{
					Driver: "redis",
					Redis: &redisConfig{
						Host:     "externalhost",
						Port:     7777,
						Username: "pr7930fnzzz3r",
						Password: "fa80uf09wi-0",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadEnv(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.DB.Driver, got.DB.Driver)
			assert.EqualValues(t, *tt.want.DB.PG, *got.DB.PG)

			assert.Equal(t, tt.want.Cache.Driver, got.Cache.Driver)
			assert.EqualValues(t, *tt.want.Cache.Redis, *got.Cache.Redis)
		})
	}
}
