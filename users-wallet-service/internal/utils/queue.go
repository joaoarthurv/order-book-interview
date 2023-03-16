package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/labstack/gommon/log"
)

type Queue struct {
	Client *sqs.Client
	ctx    context.Context
}

type QueueClient interface {
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(options *sqs.Options)) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(options *sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

type IQueueClient interface {
	ReceiveMessage(api QueueClient, queueConfig sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(api QueueClient, queueConfig sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
	GetClient() *sqs.Client
}

func NewQueue(ctx context.Context, region string) *Queue {
	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithEndpointResolverWithOptions(customResolverSQS()))
	if err != nil {
		log.Debug("Failed when recovering AWS credentials: ", err)
	}

	cfg.Region = region

	client := sqs.NewFromConfig(cfg)

	return &Queue{
		client,
		ctx,
	}
}

func (q *Queue) ReceiveMessage(api QueueClient, queueConfig sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(q.ctx, &queueConfig)
}

func (q *Queue) DeleteMessage(api QueueClient, queueConfig sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return api.DeleteMessage(q.ctx, &queueConfig)
}

func (q *Queue) GetClient() *sqs.Client {
	return q.Client
}

func customResolverSQS() aws.EndpointResolverWithOptionsFunc {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           "http://local-stack-order-book:4566",
			SigningRegion: "us-east-1",
		}, nil
	})
	return customResolver
}
