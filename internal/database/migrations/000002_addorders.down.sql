-- Drop indexes on order_items
DROP INDEX IF EXISTS idx_order_items_order_id;
DROP INDEX IF EXISTS idx_order_items_not_deleted;

-- Drop table order_items
DROP TABLE IF EXISTS order_items;

-- Drop indexes on orders
DROP INDEX IF EXISTS idx_orders_order_date;
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_customer_id;
DROP INDEX IF EXISTS idx_orders_not_deleted;

-- Drop table orders
DROP TABLE IF EXISTS orders;

-- Drop indexes on customers
DROP INDEX IF EXISTS idx_customers_email;
DROP INDEX IF EXISTS idx_customers_not_deleted;

-- Drop table customers
DROP TABLE IF EXISTS customers;
