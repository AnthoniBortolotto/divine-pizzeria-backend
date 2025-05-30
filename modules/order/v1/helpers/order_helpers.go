package order_helpers

import (
	"net/http"

	"divine-pizzeria-backend/constants"
	auth_models "divine-pizzeria-backend/modules/auth/v1/models"
	auth_repositories "divine-pizzeria-backend/modules/auth/v1/repositories"
	order_models "divine-pizzeria-backend/modules/order/v1/models"
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"
	pizza_repositories "divine-pizzeria-backend/modules/pizza/v1/repositories"

	"github.com/gin-gonic/gin"
)

type OrderValidator struct {
	authRepo        auth_repositories.AuthRepository
	pizzaSizeRepo   pizza_repositories.PizzaSizesRepository
	pizzaFlavorRepo pizza_repositories.PizzaFlavorsRepository
}

func NewOrderValidator(
	authRepo auth_repositories.AuthRepository,
	pizzaSizeRepo pizza_repositories.PizzaSizesRepository,
	pizzaFlavorRepo pizza_repositories.PizzaFlavorsRepository,
) *OrderValidator {
	return &OrderValidator{
		authRepo:        authRepo,
		pizzaSizeRepo:   pizzaSizeRepo,
		pizzaFlavorRepo: pizzaFlavorRepo,
	}
}

func (v *OrderValidator) ValidateUserPermission(c *gin.Context, permissionID uint) (*auth_models.User, bool) {
	// Get user ID from context
	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, false
	}
	user, err := v.authRepo.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return nil, false
	}

	if permissionID == constants.ALL_ROLES || user.RoleID != permissionID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to perform this action"})
		return nil, false
	}
	return user, true
}

func (v *OrderValidator) CalculateTotalPrice(reqBody order_models.CreateOrderRequest, c *gin.Context) ([]order_models.OrderItem, float64, error) {
	var totalAmount float64
	var orderItems []order_models.OrderItem

	for _, item := range reqBody.Items {
		// Get pizza size price
		pizzaSize, err := v.pizzaSizeRepo.GetPizzaSizeByID(item.PizzaSizeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pizza size"})
			return nil, 0, err
		}

		unitPrice := pizzaSize.Price
		pizzaFlavorList := []pizza_models.PizzaFlavor{}
		// Get pizza flavor additional price
		for _, flavorID := range item.FlavorIDs {
			pizzaFlavor, err := v.pizzaFlavorRepo.GetPizzaFlavorByID(flavorID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pizza flavor"})
				return nil, 0, err
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

	return orderItems, totalAmount, nil
}
