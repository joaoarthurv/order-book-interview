package service

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/labstack/gommon/log"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/model"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/utils"
)

type TopicService struct {
	iTopicClient utils.ITopicClient
	topicArn     string
}

type ITopicService interface {
	SendEvents(orderExecuted model.OrderExecutedBody) error
}

func NewTopicService(iTopicClient utils.ITopicClient, topicArn string) *TopicService {
	return &TopicService{
		iTopicClient: iTopicClient,
		topicArn:     topicArn,
	}
}

func (t *TopicService) SendEvents(orderExecuted *model.OrderExecutedBody) error {
	_, err := t.iTopicClient.Publish(t.iTopicClient.GetClient(), &sns.PublishInput{
		TopicArn: aws.String(t.topicArn),
		Message:  aws.String(utils.StructToString(orderExecuted)),
	})

	if err != nil {
		log.Errorf("Failed send executed order with [orderBuyId: %v] | [orderSellId: %v] | [orderQuantity: %v] | [orderPrice %v] to topic",
			orderExecuted.OrderBuyId, orderExecuted.OrderSellId, orderExecuted.OrderQuantity, orderExecuted.Price)
		return err
	}

	return nil
}
