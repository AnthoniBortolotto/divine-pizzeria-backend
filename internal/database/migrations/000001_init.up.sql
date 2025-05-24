-- Create user roles table
CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create indexes for user_roles
CREATE UNIQUE INDEX idx_user_roles_name ON user_roles(name);

-- Insert default roles
INSERT INTO user_roles (name, description) VALUES
    ('customer', 'Regular customer with basic access'),
    ('admin', 'Administrator with full system access');

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

-- Create indexes for users table
CREATE INDEX idx_users_not_deleted ON users(deleted_at)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_users_email ON users(email)
WHERE deleted_at IS NULL;

-- Insert default admin user (password: admin123)
INSERT INTO users (email, password, first_name, last_name, phone_number, address, role_id)
VALUES (
    'admin@divinepizzeria.com',
    '$2a$10$3Z0mehl7UuM9BvozJ.YCqe7WuyN1WvlJKpZ9dEFRay2ewm0MmfJg2', -- This is a placeholder, replace with actual hashed password
    'Admin',
    'User',
    '1234567890',
    'Admin Address',
    (SELECT id FROM user_roles WHERE name = 'admin')
);

-- Create pizza sizes table
CREATE TABLE pizza_sizes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    display_name VARCHAR(50) NOT NULL,
    price NUMERIC(10, 2) DEFAULT 0.00 NOT NULL,
    discount NUMERIC(10, 2) DEFAULT 0.00 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_pizza_sizes_not_deleted ON pizza_sizes(deleted_at)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_pizza_sizes_unique_name_not_deleted
ON pizza_sizes(name)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_pizza_sizes_unique_display_name_not_deleted
ON pizza_sizes(display_name)
WHERE deleted_at IS NULL;

-- Create pizza flavors table
CREATE TABLE pizza_flavors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    additional_price NUMERIC(10, 2) DEFAULT 0.00 NOT NULL,
    description VARCHAR(255) NOT NULL,
    ingredients VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_pizza_flavors_not_deleted ON pizza_flavors(deleted_at)
WHERE deleted_at IS NULL;

-- Create orders table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,
    total_price NUMERIC(10, 2) DEFAULT 0.00 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_orders_not_deleted ON orders(deleted_at)
WHERE deleted_at IS NULL;

CREATE INDEX idx_orders_user_id ON orders(user_id)
WHERE deleted_at IS NULL;

CREATE INDEX idx_orders_status ON orders(status)
WHERE deleted_at IS NULL;

CREATE INDEX idx_orders_order_date ON orders(order_date);

-- Create order items table
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