package customer_handlers

import (
	customer_models "divine-pizzeria-backend/modules/customer/v1/models"
	customer_repositories "divine-pizzeria-backend/modules/customer/v1/repositories"
	utils_validator "divine-pizzeria-backend/utils"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	customerRepo customer_repositories.CustomerRepository
}

func NewCustomerHandler(db *gorm.DB) *CustomerHandler {
	return &CustomerHandler{
		customerRepo: *customer_repositories.NewCustomerRepository(db),
	}
}

func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	customers, err := h.customerRepo.GetAllCustomers()
	if err != nil {
		fmt.Printf("Error on getting customers %s", err.Error())
		c.JSON(500, gin.H{
			"error": "Failed to fetch customers",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "List of customers",
		"data":    customers,
	})
}

func (h *CustomerHandler) AddCustomer(c *gin.Context) {
	var customerBody customer_models.CustomerCreate
	if err := c.ShouldBindJSON(&customerBody); err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	validate.RegisterValidation("isName", utils_validator.IsName)
	validate.RegisterValidation("isCep", utils_validator.IsCep)

	err := validate.Struct(customerBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	// Convert CustomerCreate to Customer using the mapper
	customer := customerBody.ToCustomer()

	// Create customer in database
	createdCustomer, err := h.customerRepo.CreateCustomer(customer)
	if err != nil {
		fmt.Printf("Error on creating customer %s", err.Error())
		c.JSON(500, gin.H{
			"error": "Failed to create customer",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Customer added successfully",
		"data":    createdCustomer,
	})
}
