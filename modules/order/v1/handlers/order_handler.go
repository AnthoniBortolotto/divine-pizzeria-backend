package order_handlers

import (
	"net/http"

	"divine-pizzeria-backend/constants"
	auth_repositories "divine-pizzeria-backend/modules/auth/v1/repositories"
	order_helpers "divine-pizzeria-backend/modules/order/v1/helpers"
	order_models "divine-pizzeria-backend/modules/order/v1/models"
	order_repositories "divine-pizzeria-backend/modules/order/v1/repositories"
	pizza_repositories "divine-pizzeria-backend/modules/pizza/v1/repositories"
	utils_validator "divine-pizzeria-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OrderHandler struct {
	orderRepo       order_repositories.OrderRepository
	pizzaSizeRepo   pizza_repositories.PizzaSizesRepository
	pizzaFlavorRepo pizza_repositories.PizzaFlavorsRepository
	authRepo        auth_repositories.AuthRepository
	helpers         *order_helpers.OrderValidator
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	orderRepo := order_repositories.NewOrderRepository(db)
	pizzaSizeRepo := pizza_repositories.NewPizzaSizeRepository(db)
	pizzaFlavorRepo := pizza_repositories.NewPizzaFlavoursRepository(db)
	authRepo := auth_repositories.NewAuthRepository(db)
	helpers := order_helpers.NewOrderValidator(*authRepo, *pizzaSizeRepo, *pizzaFlavorRepo)
	return &OrderHandler{
		orderRepo:       *orderRepo,
		pizzaSizeRepo:   *pizzaSizeRepo,
		pizzaFlavorRepo: *pizzaFlavorRepo,
		authRepo:        *authRepo,
		helpers:         helpers,
	}
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	_, hasPermission := h.helpers.ValidateUserPermission(c, constants.ADMIN_ROLE_ID)

	if !hasPermission {
		return
	}

	orders, err := h.orderRepo.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	// TODO: Implement order listing with user ID
	c.JSON(200, gin.H{
		"message": "List of orders",
		"orders":  orders,
	})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	// check if user exists
	user, hasPermission := h.helpers.ValidateUserPermission(c, constants.CUSTOMER_ROLE_ID)

	if !hasPermission {
		return
	}

	var reqBody order_models.CreateOrderRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validate.RegisterValidation("isdateafternow", utils_validator.IsDateAfterNow)

	err := validate.Struct(reqBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	// Calculate total amount and prepare order items
	orderItems, totalAmount, err := h.helpers.CalculateTotalPrice(reqBody, c)
	if err != nil {
		return
	}

	// Create order
	order := order_models.Order{
		UserID:       user.ID,
		Status:       order_models.OrderStatusPending,
		TotalPrice:   totalAmount,
		DeliveryDate: reqBody.DeliveryDate,
		Items:        orderItems,
	}

	newOrder, err := h.orderRepo.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   newOrder,
	})
}
