-- Add back the pizza_flavor_id column to order_items
ALTER TABLE order_items ADD COLUMN pizza_flavor_id INT;

-- Migrate data back (taking the first flavor for each order item)
UPDATE order_items oi
SET pizza_flavor_id = (
    SELECT pizza_flavor_id 
    FROM order_item_flavors oif 
    WHERE oif.order_item_id = oi.id 
    ORDER BY oif.id 
    LIMIT 1
);

-- Add back the foreign key constraint
ALTER TABLE order_items 
ADD CONSTRAINT fk_order_items_pizza_flavor 
FOREIGN KEY (pizza_flavor_id) REFERENCES pizza_flavors(id);

-- Drop indexes from order_item_flavors
DROP INDEX IF EXISTS idx_order_item_flavors_order_item_id;
DROP INDEX IF EXISTS idx_order_item_flavors_not_deleted;

-- Drop the order_item_flavors table
DROP TABLE IF EXISTS order_item_flavors; 