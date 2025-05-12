package pizza_handlers

import (
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"
	pizza_repositories "divine-pizzeria-backend/modules/pizza/v1/repositories"
	"errors"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PizzaHandler struct {
	pizzaSizeRepo   pizza_repositories.PizzaSizesRepository
	pizzaFlavorRepo pizza_repositories.PizzaFlavorsRepository
}

func NewPizzaHandler(db *gorm.DB) *PizzaHandler {
	return &PizzaHandler{
		pizzaSizeRepo:   *pizza_repositories.NewPizzaSizeRepository(db),
		pizzaFlavorRepo: *pizza_repositories.NewPizzaFlavoursRepository(db),
	}
}

func (h *PizzaHandler) ListPizzaSizes(c *gin.Context) {
	pizzaSizeList, err := h.pizzaSizeRepo.GetAllPizzaSizes()
	if err != nil {
		println("Error fetching pizza sizes:", err.Error())
		c.JSON(500, gin.H{
			"error": "Failed to fetch pizza sizes",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "List of pizza sizes",
		"data":    pizzaSizeList,
	})
}

func (h *PizzaHandler) AddPizzaSize(c *gin.Context) {
	var pizzaSizeBody pizza_models.PizzaSizeCreate

	if err := c.ShouldBindJSON(&pizzaSizeBody); err != nil {
		println("Error validating request body", err)
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}
	validate := validator.New()

	err := validate.Struct(pizzaSizeBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	type RoutineResult struct {
		Result bool
		Error  error
	}

	ch := make(chan RoutineResult, 2)

	// Check name
	go func() {
		result, err := h.pizzaSizeRepo.GetPizzaSizeByName(pizzaSizeBody.Name)

		if err != nil && !strings.Contains(err.Error(), "record not found") {
			println("Error checking name:", err.Error())
			ch <- RoutineResult{Result: false, Error: errors.New("Failed to check pizza size name")}
		} else if err != nil && strings.Contains(err.Error(), "record not found") {
			ch <- RoutineResult{Result: false, Error: nil}
			return
		}
		fmt.Printf("result GetPizzaSizeByName %+v", result)

		ch <- RoutineResult{Result: true, Error: nil}
	}()

	// Check display name
	go func() {
		result, _ := h.pizzaSizeRepo.GetPizzaSizeByDisplayName(pizzaSizeBody.DisplayName)
		if err != nil && !strings.Contains(err.Error(), "record not found") {
			println("Error checking display name:", err.Error())
			ch <- RoutineResult{Result: false, Error: errors.New("Failed to check pizza size display name")}
		} else if err != nil && strings.Contains(err.Error(), "record not found") {
			ch <- RoutineResult{Result: false, Error: nil}
			return
		}

		fmt.Printf("result GetPizzaSizeByDisplayName %+v", result)
		ch <- RoutineResult{Result: true, Error: nil}
	}()

	// Wait for both checks to complete
	for i := 0; i < 2; i++ {
		result := <-ch
		if result.Error != nil {
			println("Error from goroutine:", result.Error.Error())
			c.JSON(500, gin.H{
				"error": "Failed to check pizza size",
			})
			return
		}
		if result.Result {
			c.JSON(400, gin.H{
				"error":   "Pizza size already exists",
				"message": fmt.Sprintf("Pizza size with name %s or display name %s already exists", pizzaSizeBody.Name, pizzaSizeBody.DisplayName),
			})
			return
		}
	}

	result, err := h.pizzaSizeRepo.CreatePizzaSize(pizzaSizeBody)
	if err != nil {
		println("Error creating pizza size:", err.Error())
		c.JSON(500, gin.H{
			"error": "Failed to create pizza size",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Pizza size added",
		"data":    result,
	})
}

func (h *PizzaHandler) ListPizzaFlavors(c *gin.Context) {
	pizzaFlavorList, err := h.pizzaFlavorRepo.GetAllPizzaFlavors()

	if err != nil {
		println("Error fetching pizza flavors:", err.Error())
		c.JSON(500, gin.H{
			"error": "Failed to fetch pizza flavors",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "List of pizza flavors",
		"data":    pizzaFlavorList,
	})
}

func (h *PizzaHandler) AddPizzaFlavor(c *gin.Context) {
	var pizzaFlavorBody pizza_models.PizzaFlavorCreate

	if err := c.ShouldBindJSON(&pizzaFlavorBody); err != nil {
		println("Error validating request body", err)
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()

	err := validate.Struct(pizzaFlavorBody)

	if err != nil {
		println("Error from request body validator", err)
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	result, err := h.pizzaFlavorRepo.CreatePizzaFlavor(pizzaFlavorBody)

	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to create pizza size",
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Pizza flavor added",
		"data":    result,
	})
}
