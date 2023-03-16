package utils

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type topicMock struct {
	mock.Mock
}

func (t *topicMock) Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	args := t.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*sns.PublishOutput), args.Error(1)
}

func Test_PublishMessageOnQueue(t *testing.T) {

	t.Run("GetClient method", func(t *testing.T) {
		t.Run("Should throw error and return nil client if something fails", func(t *testing.T) {
			topicClient := NewTopicClient(context.Background(), "sa-east-1")
			assert.NotNil(t, topicClient.GetClient())
		})
	})

	t.Run("Should throw error and return nil client if something fails", func(t *testing.T) {
		mockTopic := new(topicMock)
		c := NewTopicClient(context.TODO(), "sa-east-1")

		mockTopic.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to establish a connection with aws"))
		result, err := c.Publish(mockTopic, &sns.PublishInput{
			TopicArn: aws.String("https://topic.com"),
			Message:  nil,
		})
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("Should return sns output if nothing bad happens when publishing the message", func(t *testing.T) {
		mockTopic := new(topicMock)
		c := NewTopicClient(context.TODO(), "sa-east-1")

		mockTopic.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(&sns.PublishOutput{
			MessageId: aws.String("fake-id"),
		}, nil)

		result, err := c.Publish(mockTopic, &sns.PublishInput{
			TopicArn: aws.String("https://topic.com"),
			Message:  aws.String("fake-id"),
		})
		assert.NotNil(t, result.MessageId)
		assert.Nil(t, err)
	})
}
