package dao

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joomcode/errorx"
	"gitlab.com/projetos/orderbook/orders-management-service/internal/model"
	"log"
	"time"
)

type OrdersDao struct {
	writePool *pgxpool.Pool
	readPool  *pgxpool.Pool
}

type IOrdersDao interface {
	InsertOrder(ctx context.Context, order *model.OrdersDataRow) error
	SelectBuyOrdersToExecute(ctx context.Context, order *model.OrdersDataRow) ([]model.OrdersDataRow, error)
	SelectSellOrdersToExecute(ctx context.Context, order *model.OrdersDataRow) ([]model.OrdersDataRow, error)
	DeleteOrderById(ctx context.Context, orderId string) error
	UpdateOrderQuantityById(ctx context.Context, orderId string, quantity int) error
	SelectBuyOrdersByOrderType(ctx context.Context) ([]model.OrdersDataRow, error)
	SelectSellOrdersByOrderType(ctx context.Context) ([]model.OrdersDataRow, error)
}

func NewOrdersDao(writePool *pgxpool.Pool, readPool *pgxpool.Pool) *OrdersDao {
	return &OrdersDao{
		writePool: writePool,
		readPool:  readPool,
	}
}

func (o *OrdersDao) InsertOrder(ctx context.Context, order *model.OrdersDataRow) error {
	query := o.insertOrderQuery()
	err := o.execAction(
		ctx, query, order.OrderId, order.OrderType, order.OrderPrice, order.OrderQuantity, order.OrderProductType, order.OrderOwnerId, time.Now().UTC(),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func (o *OrdersDao) insertOrderQuery() string {
	return `
		INSERT INTO orders_store.orders(
			id,
			order_type,
			price,
			quantity,
			product_type,
			owner_id,
			order_timestamp
  		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
}

func (o *OrdersDao) SelectBuyOrdersToExecute(ctx context.Context, order *model.OrdersDataRow) ([]model.OrdersDataRow, error) {
	query := o.selectBuyOrdersToExecuteQuery()
	orders := make([]model.OrdersDataRow, 0)
	err := o.selectAction(
		ctx, &orders, query, order.OrderPrice, order.OrderOwnerId,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "Failed when run select orders query")
	}
	return orders, nil
}

func (o *OrdersDao) selectBuyOrdersToExecuteQuery() string {
	return `
		SELECT *
		FROM orders_store.Orders
		WHERE order_type = 'BUY' AND price >= $1 AND owner_id != $2
		ORDER BY price DESC, order_timestamp ASC
	`
}

func (o *OrdersDao) SelectSellOrdersToExecute(ctx context.Context, order *model.OrdersDataRow) ([]model.OrdersDataRow, error) {
	query := o.selectSellOrdersToExecuteQuery()
	orders := make([]model.OrdersDataRow, 0)
	err := o.selectAction(
		ctx, &orders, query, order.OrderPrice, order.OrderOwnerId,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "Failed when run select orders query")
	}
	return orders, nil
}

func (o *OrdersDao) selectSellOrdersToExecuteQuery() string {
	return `
		SELECT *
		FROM orders_store.Orders
		WHERE order_type = 'SELL' AND price <= $1 AND owner_id != $2
		ORDER BY price ASC, order_timestamp ASC
	`
}

func (o *OrdersDao) DeleteOrderById(ctx context.Context, orderId string) error {
	query := o.deleteOrderByIdQuery()

	err := o.execAction(
		ctx, query, orderId,
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil

}

func (o *OrdersDao) deleteOrderByIdQuery() string {
	return `
		DELETE 
		FROM orders_store.Orders 
		WHERE id = $1;
	`
}

func (o *OrdersDao) UpdateOrderQuantityById(ctx context.Context, orderId string, quantity int) error {
	query := o.updateOrderQuantityByIdQuery()

	err := o.execAction(
		ctx, query, quantity, orderId,
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil

}

func (o *OrdersDao) updateOrderQuantityByIdQuery() string {
	return `
		UPDATE orders_store.Orders 
        SET quantity = $1
		WHERE id = $2;
	`
}

func (o *OrdersDao) SelectBuyOrdersByOrderType(ctx context.Context) ([]model.OrdersDataRow, error) {
	query := o.selectBuyOrdersByOrderTypeQuery()

	orders := make([]model.OrdersDataRow, 0)
	err := o.selectAction(
		ctx, &orders, query,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "Failed when run select orders query")
	}
	return orders, nil

}

func (o *OrdersDao) selectBuyOrdersByOrderTypeQuery() string {
	return `
		SELECT *  
		FROM orders_store.Orders
		WHERE order_type = 'BUY'
		ORDER BY price DESC, order_timestamp ASC
	`
}

func (o *OrdersDao) SelectSellOrdersByOrderType(ctx context.Context) ([]model.OrdersDataRow, error) {
	query := o.selectSellOrdersByOrderTypeQuery()

	orders := make([]model.OrdersDataRow, 0)
	err := o.selectAction(
		ctx, &orders, query,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "Failed when run select orders query")
	}
	return orders, nil

}

func (o *OrdersDao) selectSellOrdersByOrderTypeQuery() string {
	return `
		SELECT *  
		FROM orders_store.Orders
		WHERE order_type = 'SELL'
		ORDER BY price ASC, order_timestamp ASC
	`
}

func (o *OrdersDao) selectAction(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	err := pgxscan.Select(
		ctx, o.readPool, dst, query, args...,
	)
	return err
}

func (o *OrdersDao) execAction(ctx context.Context, query string, args ...interface{}) error {
	_, err := o.writePool.Exec(ctx, query, args...)
	return err
}

func (o *OrdersDao) getAction(ctx context.Context, db pgxscan.Querier, dst interface{}, query string, args ...interface{}) error {
	err := pgxscan.Get(
		ctx, db, dst, query, args...,
	)
	return err
}
