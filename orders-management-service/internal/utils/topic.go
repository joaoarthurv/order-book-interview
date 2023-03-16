package utils

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Topic struct {
	client *sns.Client
	ctx    context.Context
}

type TopicClient interface {
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type ITopicClient interface {
	Publish(api TopicClient, params *sns.PublishInput) (*sns.PublishOutput, error)
	GetClient() *sns.Client
}

func NewTopicClient(ctx context.Context, region string) *Topic {
	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithEndpointResolverWithOptions(customResolver()))
	if err != nil {
		fmt.Println("Failed when recovering AWS credentials: ", err)
	}

	cfg.Region = region

	client := sns.NewFromConfig(cfg)

	return &Topic{
		client,
		ctx,
	}
}

func (t *Topic) Publish(api TopicClient, params *sns.PublishInput) (*sns.PublishOutput, error) {
	output, err := api.Publish(t.ctx, params)
	if err != nil {
		fmt.Println(">>> Failed when publishing message to the topic: ", err)
	}
	return output, err
}

func (t *Topic) GetClient() *sns.Client {
	return t.client
}

func customResolver() aws.EndpointResolverWithOptionsFunc {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           "http://local-stack-order-book:4566",
			SigningRegion: "us-east-1",
		}, nil
	})
	return customResolver
}
