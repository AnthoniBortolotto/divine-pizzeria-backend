package customer_repositories

import (
	customer_models "divine-pizzeria-backend/modules/customer/v1/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetAllCustomers() ([]customer_models.Customer, error) {
	var customers []customer_models.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		println("Error fetching customers:", err)
		return nil, err
	}
	return customers, nil
}
func (r *CustomerRepository) CreateCustomer(customer customer_models.Customer) (customer_models.Customer, error) {

	newCustomer := customer_models.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
	}
	if err := r.db.Create(&newCustomer).Error; err != nil {
		println("Error creating customer:", err)
		return newCustomer, err
	}
	return newCustomer, nil
}
