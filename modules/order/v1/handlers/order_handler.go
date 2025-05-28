package order_handlers

import (
	"net/http"

	"divine-pizzeria-backend/constants"
	auth_repositories "divine-pizzeria-backend/modules/auth/v1/repositories"
	order_models "divine-pizzeria-backend/modules/order/v1/models"
	order_repositories "divine-pizzeria-backend/modules/order/v1/repositories"
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"
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
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	orderRepo := order_repositories.NewOrderRepository(db)
	pizzaSizeRepo := pizza_repositories.NewPizzaSizeRepository(db)
	pizzaFlavorRepo := pizza_repositories.NewPizzaFlavoursRepository(db)
	authRepo := auth_repositories.NewAuthRepository(db)
	return &OrderHandler{
		orderRepo:       *orderRepo,
		pizzaSizeRepo:   *pizzaSizeRepo,
		pizzaFlavorRepo: *pizzaFlavorRepo,
		authRepo:        *authRepo,
	}
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// TODO: Implement order listing with user ID
	c.JSON(200, gin.H{
		"message": "List of orders",
		"user_id": userID,
	})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {

	user_id := c.GetUint("user_id")

	if user_id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
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

	// check if user exists

	user, err := h.authRepo.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	if user.RoleID != constants.CUSTOMER_ROLE_ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only customers can create orders"})
		return
	}

	// Calculate total amount and prepare order items
	var totalAmount float64
	var orderItems []order_models.OrderItem

	for _, item := range reqBody.Items {
		// Get pizza size price
		pizzaSize, err := h.pizzaSizeRepo.GetPizzaSizeByID(item.PizzaSizeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pizza size"})
			return
		}

		unitPrice := pizzaSize.Price
		pizzaFlavorList := []pizza_models.PizzaFlavor{}
		// Get pizza flavor additional price
		for _, flavorID := range item.FlavorIDs {
			pizzaFlavor, err := h.pizzaFlavorRepo.GetPizzaFlavorByID(flavorID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pizza flavor"})
				return
			}
			pizzaFlavorList = append(pizzaFlavorList, pizzaFlavor)
			unitPrice += pizzaFlavor.AdditionalPrice
		}
		itemTotal := unitPrice * float64(item.Quantity)
		totalAmount += itemTotal

		orderItems = append(orderItems, order_models.OrderItem{
			PizzaSizeID: pizzaSize.ID,
			Flavors:     pizzaFlavorList,
			Quantity:    item.Quantity,
			UnitPrice:   unitPrice,
		})
	}

	// Create order
	order := order_models.Order{
		UserID:       user_id,
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
