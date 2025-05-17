CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_customers_not_deleted ON customers(deleted_at)
WHERE deleted_at IS NULL;


CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,
    total_amount NUMERIC(10, 2) DEFAULT 0.00 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE INDEX idx_orders_not_deleted ON orders(deleted_at)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_customers_email ON customers(email)
WHERE deleted_at IS NULL;

CREATE INDEX idx_orders_customer_id ON orders(customer_id)
WHERE deleted_at IS NULL;

CREATE INDEX idx_orders_status ON orders(status)
WHERE deleted_at IS NULL;

CREATE INDEX idx_orders_order_date ON orders(order_date);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    pizza_size_id INT NOT NULL,
    pizza_flavor_id INT NOT NULL,
    quantity INT DEFAULT 1 NOT NULL,
    unit_price NUMERIC(10, 2) DEFAULT 0.00 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (pizza_size_id) REFERENCES pizza_sizes(id),
    FOREIGN KEY (pizza_flavor_id) REFERENCES pizza_flavors(id)
);

CREATE INDEX idx_order_items_not_deleted ON order_items(deleted_at)
WHERE deleted_at IS NULL;

CREATE INDEX idx_order_items_order_id ON order_items(order_id)
WHERE deleted_at IS NULL;
