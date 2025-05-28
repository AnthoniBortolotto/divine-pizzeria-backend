package pizza_repositories

import (
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"
	"fmt"

	"gorm.io/gorm"
)

type PizzaSizesRepository struct {
	db *gorm.DB
}

func NewPizzaSizeRepository(db *gorm.DB) *PizzaSizesRepository {
	return &PizzaSizesRepository{
		db: db,
	}
}

func (r *PizzaSizesRepository) GetAllPizzaSizes() ([]pizza_models.PizzaSize, error) {
	var pizzaSizes []pizza_models.PizzaSize
	if err := r.db.Find(&pizzaSizes).Error; err != nil {
		println("Error fetching pizza sizes:", err)
		return nil, err
	}
	return pizzaSizes, nil
}

func (r *PizzaSizesRepository) GetPizzaSizeByName(name string) (pizza_models.PizzaSize, error) {
	var pizzaSize pizza_models.PizzaSize
	if err := r.db.Where("name = ?", name).Where("deleted_at IS NULL").First(&pizzaSize).Error; err != nil {
		println("Error fetching pizza size by name:", err)
		return pizza_models.PizzaSize{}, err
	}
	return pizzaSize, nil
}

func (r *PizzaSizesRepository) GetPizzaSizeByDisplayName(displayName string) (pizza_models.PizzaSize, error) {
	var pizzaSize pizza_models.PizzaSize
	if err := r.db.Where("display_name = ?", displayName).Where("deleted_at IS NULL").First(&pizzaSize).Error; err != nil {
		println("Error fetching pizza size by display name:", err)
		return pizza_models.PizzaSize{}, err
	}
	return pizzaSize, nil
}

func (r *PizzaSizesRepository) GetPizzaSizeByID(id uint) (pizza_models.PizzaSize, error) {
	var pizzaSize pizza_models.PizzaSize
	if err := r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&pizzaSize).Error; err != nil {
		println("Error fetching pizza size by ID:", err)
		return pizza_models.PizzaSize{}, err
	}
	return pizzaSize, nil
}

func (r *PizzaSizesRepository) CreatePizzaSize(pizzaSize pizza_models.PizzaSizeCreate) (pizza_models.PizzaSize, error) {
	newPizzaSize := pizza_models.PizzaSize{
		Name:        pizzaSize.Name,
		DisplayName: pizzaSize.DisplayName,
		Price:       pizzaSize.Price,
		Discount:    pizzaSize.Discount,
	}
	fmt.Printf("Creating new pizza size: %+v\n", newPizzaSize)
	if err := r.db.Create(&newPizzaSize).Error; err != nil {
		println("Error creating pizza size:", err)
		return newPizzaSize, err
	}
	return newPizzaSize, nil
}
