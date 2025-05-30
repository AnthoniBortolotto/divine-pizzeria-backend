package order_validators

import (
	"divine-pizzeria-backend/constants"
	auth_repositories "divine-pizzeria-backend/modules/auth/v1/repositories"
	order_models "divine-pizzeria-backend/modules/order/v1/models"
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"
	pizza_repositories "divine-pizzeria-backend/modules/pizza/v1/repositories"
	"sync"
)

type ValidationResult struct {
	Error error
	Data  interface{}
}

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

func (v *OrderValidator) ValidateUser(userID uint) ValidationResult {
	user, err := v.authRepo.GetUserByID(userID)
	if err != nil {
		return ValidationResult{Error: err}
	}
	if user.RoleID != constants.CUSTOMER_ROLE_ID {
		return ValidationResult{Error: err}
	}
	return ValidationResult{Data: user}
}

func (v *OrderValidator) ValidatePizzaSize(sizeID uint) ValidationResult {
	pizzaSize, err := v.pizzaSizeRepo.GetPizzaSizeByID(sizeID)
	if err != nil {
		return ValidationResult{Error: err}
	}
	return ValidationResult{Data: pizzaSize}
}

func (v *OrderValidator) ValidatePizzaFlavor(flavorID uint) ValidationResult {
	pizzaFlavor, err := v.pizzaFlavorRepo.GetPizzaFlavorByID(flavorID)
	if err != nil {
		return ValidationResult{Error: err}
	}
	return ValidationResult{Data: pizzaFlavor}
}

func (v *OrderValidator) ValidateOrderItems(items []order_models.OrderItemCreate) ([]order_models.OrderItem, float64, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var orderItems []order_models.OrderItem
	var totalAmount float64
	var validationErrors []error

	// Create channels for results
	sizeResults := make(chan ValidationResult, len(items))
	flavorResults := make(map[uint]chan ValidationResult)

	// Validate pizza sizes
	for _, item := range items {
		wg.Add(1)
		go func(sizeID uint) {
			defer wg.Done()
			sizeResults <- v.ValidatePizzaSize(sizeID)
		}(item.PizzaSizeID)
	}

	// Collect size results
	go func() {
		wg.Wait()
		close(sizeResults)
	}()

	// Process size results and validate flavors
	sizeMap := make(map[uint]*pizza_models.PizzaSize)
	for result := range sizeResults {
		if result.Error != nil {
			mu.Lock()
			validationErrors = append(validationErrors, result.Error)
			mu.Unlock()
			continue
		}
		if size, ok := result.Data.(pizza_models.PizzaSize); ok {
			sizeMap[size.ID] = &size
		}
	}

	// Validate flavors
	for _, item := range items {
		for _, flavorID := range item.FlavorIDs {
			if _, exists := flavorResults[flavorID]; !exists {
				flavorResults[flavorID] = make(chan ValidationResult, 1)
				wg.Add(1)
				go func(fID uint) {
					defer wg.Done()
					flavorResults[fID] <- v.ValidatePizzaFlavor(fID)
				}(flavorID)
			}
		}
	}

	// Collect flavor results
	go func() {
		wg.Wait()
		for _, ch := range flavorResults {
			close(ch)
		}
	}()

	// Process results and build order items
	flavorMap := make(map[uint]*pizza_models.PizzaFlavor)
	for _, ch := range flavorResults {
		for result := range ch {
			if result.Error != nil {
				mu.Lock()
				validationErrors = append(validationErrors, result.Error)
				mu.Unlock()
				continue
			}
			if flavor, ok := result.Data.(pizza_models.PizzaFlavor); ok {
				flavorMap[uint(flavor.ID)] = &flavor
			}
		}
	}

	// Build order items
	for _, item := range items {
		size, sizeExists := sizeMap[item.PizzaSizeID]
		if !sizeExists {
			continue
		}

		var flavors []pizza_models.PizzaFlavor
		unitPrice := size.Price

		for _, flavorID := range item.FlavorIDs {
			if flavor, exists := flavorMap[flavorID]; exists {
				flavors = append(flavors, *flavor)
				unitPrice += flavor.AdditionalPrice
			}
		}

		itemTotal := unitPrice * float64(item.Quantity)
		totalAmount += itemTotal

		orderItems = append(orderItems, order_models.OrderItem{
			PizzaSizeID: size.ID,
			Flavors:     flavors,
			Quantity:    item.Quantity,
			UnitPrice:   unitPrice,
		})
	}

	if len(validationErrors) > 0 {
		return nil, 0, validationErrors[0]
	}

	return orderItems, totalAmount, nil
}
