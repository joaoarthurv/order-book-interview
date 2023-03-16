package model

import "time"

type UserWallet struct {
	UserId          string    `json:"userId,omitempty"  dynamodbav:"user_id"`
	ProductType     string    `json:"productType,omitempty"  dynamodbav:"product_type"`
	Balance         float64   `json:"balance,omitempty"  dynamodbav:"balance"`
	ProductQuantity int       `json:"productQuantity,omitempty"  dynamodbav:"product_quantity"`
	CreatedAt       time.Time `json:"created_at,omitempty"  dynamodbav:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty"  dynamodbav:"updated_at"`
}
