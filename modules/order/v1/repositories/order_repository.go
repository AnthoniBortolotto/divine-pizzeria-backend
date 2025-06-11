package order_repositories

import (
	customer_models "divine-pizzeria-backend/modules/customer/v1/models"
	order_dtos "divine-pizzeria-backend/modules/order/v1/dtos"
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

func (r *OrderRepository) GetAllOrders() ([]order_models.Order, error) {
	var orders []order_models.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderList(filter order_dtos.OrderFilter) ([]order_models.Order, error) {
	var orders []order_models.Order
	query := r.db.Model(&order_models.Order{}).
		Joins("JOIN users ON orders.user_id = users.id").
		Preload("User").
		Preload("Items").
		Preload("Items.PizzaSize").
		Preload("Items.Flavors")

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.StartDate != "" {
		query = query.Where("orders.created_at >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		query = query.Where("orders.created_at <= ?", filter.EndDate)
	}

	if filter.UserName != "" {
		query = query.Where("LOWER(CONCAT(users.first_name, ' ', users.last_name)) LIKE LOWER(?)", "%"+filter.UserName+"%")
	}
	if filter.Email != "" {
		query = query.Where("users.email LIKE ?", "%"+filter.Email+"%")
	}

	if filter.Sort != "" && (filter.Sort == "ASC" || filter.Sort == "DESC") {
		query = query.Order("orders.created_at " + filter.Sort)
	}
	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}

	// Map User data to Customer field
	for i := range orders {
		orders[i].Customer = customer_models.Customer{
			FirstName:   orders[i].User.FirstName,
			LastName:    orders[i].User.LastName,
			Email:       orders[i].User.Email,
			PhoneNumber: orders[i].User.PhoneNumber,
			Address:     orders[i].User.Address,
		}
	}

	return orders, nil
}

func (r *OrderRepository) GetOrdersByUserID(userID uint) ([]order_models.Order, error) {
	var orders []order_models.Order
	if err := r.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
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

func (r *OrderRepository) UpdateOrder(order *order_models.Order) (*order_models.Order, error) {
	if err := r.db.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
