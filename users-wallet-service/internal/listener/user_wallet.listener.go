package listener

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/labstack/gommon/log"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/model"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/service"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/utils"
	"sync"
)

var (
	waitGroup sync.WaitGroup
)

func UserWalletListener(userWalletService *service.UserWalletService, queue utils.IQueueClient) {
	queueUrl := "http://local-stack-order-book:4566/000000000000/ORDERS-EXECUTED-TO-WALLET"

	for {
		sqsMessage, queueErr := queue.ReceiveMessage(queue.GetClient(), sqs.ReceiveMessageInput{
			QueueUrl:            &queueUrl,
			MaxNumberOfMessages: 10,
		})

		if queueErr != nil {
			log.Error("error on receiving message. Check the status of the queue server to server.")
			continue
		}

		if sqsMessage.Messages != nil && len(sqsMessage.Messages) > 0 {
			waitGroup.Add(len(sqsMessage.Messages))
			for index := range sqsMessage.Messages {
				go func(key types.Message) {
					var messageToConvert *model.Message
					defer waitGroup.Done()

					if err := json.Unmarshal([]byte(*key.Body), &messageToConvert); err != nil {
						log.Error("failed when unmarshalling events message body")
						return
					}

					if messageToConvert != nil {
						var userWalletBody *model.OrderExecutedBody

						if err := json.Unmarshal([]byte(messageToConvert.Message), &userWalletBody); err != nil {
							log.Error("failed when unmarshalling events userWalletBody message")
							return
						}

						if userWalletBody != nil {
							if err := userWalletService.ProcessOrderExecuted(userWalletBody); err != nil {
								log.Error("Failed to ProcessOrderExecuted")
								return
							}
							if _, errorToDeleteMessage := queue.DeleteMessage(queue.GetClient(), sqs.DeleteMessageInput{
								QueueUrl:      &queueUrl,
								ReceiptHandle: key.ReceiptHandle,
							}); errorToDeleteMessage != nil {
								return
							}
						}
					}
				}(sqsMessage.Messages[index])
			}
			waitGroup.Wait()
		}
	}
}
