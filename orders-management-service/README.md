# order-book-management-service
> Service responsible for  orders (buy and sell) management on orders book context

## What it does?
* The service receives sell and buy order requests and manages them accordingly. 
Currently, these orders are received via an API and are validated by calling the 
"users-wallet" service. The purpose of the validation is to ensure that the user has 
sufficient resources to execute the order, thus preventing the creation of orders that cannot 
be executed. Once the validation is confirmed, the request is sent to the postgreSQL database.
If an order is executed, it will be sent to an SNS topic. Two SQS are subscribed to this topic in their
respective contexts: order history and user wallets. At this moment, only the user wallet service has been 
implemented.

## Who uses it?
* This system can be used by any client interested in a service that has the logic of orders book implemented.

## Service outage impacts
* If this service goes down, new purchase and sale orders will be discarded. An evolution of the product would be to implement a Kinesis stream and provide resilience for this service.

## Architecture
![img.png](img.png)

## Technologies and Dependencies
* Golang 1.19.2
* PostgresSQL
* SNS
* Docker

## Step-by-step guide to run the application:

1. Clone this repository on your local environment.
2. Open the script folder and access the file: `docker-compose.yml`.
   * Change the $ABSOLUTE_PATH variable to match your local environment [Line 9 and 57]. For example:
>$ABSOLUTE_PATH/orderbook/orders-management-service -> /Users/joao.vale/go/src/gitlab.com/projects/orderbook/orders-management-service
3. Go to the root of the project and run the command `$ make up` in the terminal. This command will start the script that will create the local environment and all the necessary functionalities to run the service.
4. To stop, just run the command `$ make down`.