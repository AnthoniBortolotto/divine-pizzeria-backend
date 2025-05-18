-- Drop indexes on order_items
DROP INDEX IF EXISTS idx_order_items_order_id;
DROP INDEX IF EXISTS idx_order_items_not_deleted;

-- Drop table order_items
DROP TABLE IF EXISTS order_items;

-- Drop indexes on orders
DROP INDEX IF EXISTS idx_orders_order_date;
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_user_id;
DROP INDEX IF EXISTS idx_orders_not_deleted;

-- Drop table orders
DROP TABLE IF EXISTS orders;

-- Drop indexes on pizza_flavors
DROP INDEX IF EXISTS idx_pizza_flavors_not_deleted;

-- Drop table pizza_flavors
DROP TABLE IF EXISTS pizza_flavors;

-- Drop indexes on pizza_sizes
DROP INDEX IF EXISTS idx_pizza_sizes_unique_display_name_not_deleted;
DROP INDEX IF EXISTS idx_pizza_sizes_unique_name_not_deleted;
DROP INDEX IF EXISTS idx_pizza_sizes_not_deleted;

-- Drop table pizza_sizes
DROP TABLE IF EXISTS pizza_sizes;

-- Drop indexes on users
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_not_deleted;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop indexes on user_roles
DROP INDEX IF EXISTS idx_user_roles_name;

-- Drop user_roles table
DROP TABLE IF EXISTS user_roles;
