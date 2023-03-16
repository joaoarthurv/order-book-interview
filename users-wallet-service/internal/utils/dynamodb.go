package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamoClient() *dynamodb.Client {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithEndpointResolverWithOptions(customResolverDynamoDb()),
	)
	if err != nil {
	}
	client := dynamodb.NewFromConfig(cfg)
	return client
}

func customResolverDynamoDb() aws.EndpointResolverWithOptionsFunc {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           "http://local-stack-order-book:4566",
			SigningRegion: "us-east-1",
		}, nil
	})
	return customResolver
}
