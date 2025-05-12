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