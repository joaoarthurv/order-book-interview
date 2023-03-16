CREATE SCHEMA orders_store;

CREATE TABLE orders_store.Orders
(
    id              UUID,
    order_type      VARCHAR,
    price           DECIMAL,
    quantity        INTEGER,
    product_type    VARCHAR,
    owner_id        VARCHAR,
    order_timestamp timestamp,
    PRIMARY KEY (id)
);

CREATE INDEX idx_order_book_price ON orders_store.Orders (price);

CREATE INDEX idx_order_book_type ON orders_store.Orders (order_type);
