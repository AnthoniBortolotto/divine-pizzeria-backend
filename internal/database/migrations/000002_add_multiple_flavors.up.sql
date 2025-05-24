-- Create order_item_flavors junction table
CREATE TABLE order_item_flavors (
    id SERIAL PRIMARY KEY,
    order_item_id INT NOT NULL,
    pizza_flavor_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (order_item_id) REFERENCES order_items(id),
    FOREIGN KEY (pizza_flavor_id) REFERENCES pizza_flavors(id)
);

-- Create indexes for order_item_flavors
CREATE INDEX idx_order_item_flavors_not_deleted ON order_item_flavors(deleted_at)
WHERE deleted_at IS NULL;

CREATE INDEX idx_order_item_flavors_order_item_id ON order_item_flavors(order_item_id)
WHERE deleted_at IS NULL;

-- Migrate existing data
INSERT INTO order_item_flavors (order_item_id, pizza_flavor_id)
SELECT id, pizza_flavor_id FROM order_items;

-- Remove the pizza_flavor_id column from order_items
ALTER TABLE order_items DROP COLUMN pizza_flavor_id; 