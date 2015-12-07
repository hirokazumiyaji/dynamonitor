package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewClient(config *Config) *dynamodb.DynamoDB {
	awsConfig := aws.NewConfig()
	if config.Aws.Region != "" {
		awsConfig.WithRegion(config.Aws.Region)
	}

	if config.Aws.AccessKeyId != "" && config.Aws.SecretAccessKey != "" {
		awsConfig.WithCredentials(
			credentials.NewStaticCredentials(
				config.Aws.AccessKeyId,
				config.Aws.SecretAccessKey,
				"",
			),
		)
	}

	if config.Aws.EndpointUrl != "" {
		awsConfig.WithEndpoint(config.Aws.EndpointUrl)
	}

	return dynamodb.New(session.New(), awsConfig)
}
