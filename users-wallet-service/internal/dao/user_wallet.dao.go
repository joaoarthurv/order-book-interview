package dao

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/model"
)

type UserWalletDao struct {
	dynamoClient *dynamodb.Client
}

type IUserWalletDao interface {
}

func NewIUserWalletDao(dynamoClient *dynamodb.Client) *UserWalletDao {
	return &UserWalletDao{
		dynamoClient: dynamoClient,
	}
}

const UsersWallet = "users-wallet"

func (d *UserWalletDao) SaveUserWallet(userWallet *model.UserWallet) error {
	av, err := attributevalue.MarshalMap(userWallet)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(UsersWallet),
	}
	_, err = d.dynamoClient.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}

func (d *UserWalletDao) GetUserWalletByUserIdAndProductType(userId string, productType string) (*model.UserWallet, error) {
	userIdValue, _ := attributevalue.Marshal(userId)
	productTypeValue, _ := attributevalue.Marshal(productType)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(UsersWallet),
		Key:       map[string]types.AttributeValue{"user_id": userIdValue, "product_type": productTypeValue},
	}

	result, err := d.dynamoClient.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("could not find user")
	}

	var userWallet *model.UserWallet

	err = attributevalue.UnmarshalMap(result.Item, &userWallet)
	if err != nil {
		return nil, err
	}

	return userWallet, nil
}
