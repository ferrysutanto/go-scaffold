package ddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/ferrysutanto/go-errors"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"go.opentelemetry.io/otel"
)

type ddbModel struct {
	c *dynamodb.Client
}

type Config struct {
	// You can supply either an sql.DB object or config of the database connection, but not both (mutually exclusive)
	Client *dynamodb.Client

	AwsAccessKeyID     *string
	AwsSecretAccessKey *string
	AwsSessionToken    *string
	AwsSource          *string
	AwsRegion          *string
	AwsEndpointURL     *string
}

func New(ctx context.Context, cfg *Config) (db.IDB, error) {
	if ctx == nil {
		return nil, errors.New("[repositories/db][newDdbModel] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][newDdbModel]")
	defer span.End()

	if cfg.Client != nil {
		return &ddbModel{
			cfg.Client,
		}, nil
	}

	accessKeyID := ""
	if cfg.AwsAccessKeyID != nil {
		accessKeyID = *cfg.AwsAccessKeyID
	}
	secretAccessKey := ""
	if cfg.AwsSecretAccessKey != nil {
		secretAccessKey = *cfg.AwsSecretAccessKey
	}
	sessionToken := ""
	if cfg.AwsSessionToken != nil {
		sessionToken = *cfg.AwsSessionToken
	}
	awsRegion := ""
	if cfg.AwsRegion != nil {
		awsRegion = *cfg.AwsRegion
	}
	awsSource := ""
	if cfg.AwsSource != nil {
		awsSource = *cfg.AwsSource
	}

	options := make([]func(*config.LoadOptions) error, 0)
	if cfg.AwsEndpointURL != nil {
		options = append(options, config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           *cfg.AwsEndpointURL,
				Source:        aws.EndpointSourceCustom,
				SigningRegion: awsRegion,
			}, nil
		})))
	}

	if cfg.AwsAccessKeyID != nil && cfg.AwsSecretAccessKey != nil {
		options = append(options, config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: secretAccessKey,
				SessionToken:    sessionToken,
				Source:          awsSource,
			},
		}))
	}

	ddbCfg, err := config.LoadDefaultConfig(ctx, options...)
	if err != nil {
		err = errors.WrapWithCode(err, "failed to load default config", 500)
		// record error in span
		span.RecordError(err)

		return nil, err
	}

	c := dynamodb.NewFromConfig(ddbCfg)

	return &ddbModel{
		c,
	}, nil
}

func (m *ddbModel) Ping(ctx context.Context) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.NewWithCode("context is required", 400)
	}

	// 2. start span and defer span end
	_, span := otel.Tracer("").Start(ctx, "[repositories/db][ddbModel:Ping]")
	defer span.End()

	return nil
}
