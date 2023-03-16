export AWS_ACCESS_KEY_ID=foo
export AWS_SECRET_ACCESS_KEY=foo

aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name ORDERS-EXECUTED-TO-WALLET

aws --endpoint-url=http://localhost:4566 sns create-topic --name LOCAL_ORDERS_EXECUTED

aws --endpoint-url http://localhost:4566  sns subscribe \
    --topic-arn arn:aws:sns:us-east-1:000000000000:LOCAL_ORDERS_EXECUTED \
    --protocol sqs \
    --notification-endpoint arn:aws:sqs:us-east-1:000000000000:ORDERS-EXECUTED-TO-WALLET

aws dynamodb create-table --cli-input-json '{
                                              "TableName": "users-wallet",
                                              "AttributeDefinitions": [
                                                {
                                                  "AttributeName": "user_id",
                                                  "AttributeType": "S"
                                                },
                                                {
                                                  "AttributeName": "product_type",
                                                  "AttributeType": "S"
                                                }
                                              ],
                                              "KeySchema": [
                                                {
                                                  "AttributeName": "user_id",
                                                  "KeyType": "HASH"
                                                },
                                                {
                                                  "AttributeName": "product_type",
                                                  "KeyType": "RANGE"
                                                }
                                              ],
                                              "ProvisionedThroughput": {
                                                "ReadCapacityUnits": 5,
                                                "WriteCapacityUnits": 5
                                              }
                                            }' --endpoint-url=http://localhost:4566 --region=us-east-1
aws dynamodb update-time-to-live --table-name users-wallet --time-to-live-specification "Enabled=false, AttributeName=ttl" --endpoint-url=http://localhost:4566 --region=us-east-1
aws dynamodb --endpoint-url=http://localhost:4566  batch-write-item --region us-east-1 --request-items '{
                                                                                                          "users-wallet": [
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "1" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "2469" },
                                                                                                                  "product_quantity": { "N": "1163" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939347Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "2" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "3340" },
                                                                                                                  "product_quantity": { "N": "8226" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939400Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "3" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "3127" },
                                                                                                                  "product_quantity": { "N": "1963" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939412Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "4" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "2490" },
                                                                                                                  "product_quantity": { "N": "2799" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939420Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "5" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "4446" },
                                                                                                                  "product_quantity": { "N": "7620" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939427Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "6" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "2447" },
                                                                                                                  "product_quantity": { "N": "5725" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939436Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "7" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "2896" },
                                                                                                                  "product_quantity": { "N": "8896" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939444Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "8" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "4462" },
                                                                                                                  "product_quantity": { "N": "5472" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939458Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "9" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "4687" },
                                                                                                                  "product_quantity": { "N": "4280" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939468Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            },
                                                                                                            {
                                                                                                              "PutRequest": {
                                                                                                                "Item": {
                                                                                                                  "user_id": { "S": "10" },
                                                                                                                  "product_type": { "S": "VIBRANIUM" },
                                                                                                                  "balance": { "N": "2515" },
                                                                                                                  "product_quantity": { "N": "8439" },
                                                                                                                  "created_at": { "S": "2023-03-15T14:44:41.939478Z" }
                                                                                                                }
                                                                                                              }
                                                                                                            }
                                                                                                          ]
                                                                                                        }
'

