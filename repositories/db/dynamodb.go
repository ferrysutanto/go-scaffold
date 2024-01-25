package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type ddbModel struct {
	c *dynamodb.Client
}

func (m *ddbModel) Ping(ctx context.Context) error {
	// 1. check if context is provided
	if ctx == nil {
		return errors.New("[repositories/db][ddbModel:Ping] context is nil")
	}

	// 2. start span and defer span end
	_, span := otel.Tracer("").Start(ctx, "[repositories/db][ddbModel:Ping]")
	defer span.End()

	return nil
}

func newDdbModel(ctx context.Context, conf *Config) (*ddbModel, error) {
	if ctx == nil {
		return nil, errors.New("[repositories/db][newDdbModel] context is nil")
	}

	// 2. start span and defer span end
	ctx, span := otel.Tracer("").Start(ctx, "[repositories/db][newDdbModel]")
	defer span.End()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           conf.AwsEndpointURL,
				Source:        aws.EndpointSourceCustom,
				SigningRegion: conf.AwsRegion,
			}, nil
		})),

		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     conf.AwsAccessKeyID,
				SecretAccessKey: conf.AwsSecretAccessKey,
				SessionToken:    conf.AwsSessionToken,
				Source:          conf.AwsSource,
			},
		}),
	)
	if err != nil {
		err = errors.Wrap(err, "[repositories/db][newDdbModel] failed to load default config")
		// record error in span
		span.RecordError(err)

		return nil, err
	}

	c := dynamodb.NewFromConfig(cfg)

	return &ddbModel{
		c,
	}, nil
}
