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
