package order_repositories

import (
	order_models "divine-pizzeria-backend/modules/order/v1/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

//todo: implement the repository methods for order

func (r *OrderRepository) GetAllOrders() ([]order_models.Order, error) {
	var orders []order_models.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) CreateOrder(order order_models.Order) (*order_models.Order, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *OrderRepository) GetOrderByID(id int) (*order_models.Order, error) {
	var order order_models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
