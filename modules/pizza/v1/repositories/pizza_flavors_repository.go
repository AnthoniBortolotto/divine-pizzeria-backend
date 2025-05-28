package pizza_repositories

import (
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"

	"gorm.io/gorm"
)

type PizzaFlavorsRepository struct {
	db *gorm.DB
}

func NewPizzaFlavoursRepository(db *gorm.DB) *PizzaFlavorsRepository {
	return &PizzaFlavorsRepository{
		db: db,
	}
}

func (r *PizzaFlavorsRepository) GetAllPizzaFlavors() ([]pizza_models.PizzaFlavor, error) {
	var pizzaFlavours []pizza_models.PizzaFlavor
	if err := r.db.Find(&pizzaFlavours).Error; err != nil {
		println("Error fetching pizza flavours:", err)
		return nil, err
	}
	return pizzaFlavours, nil
}

func (r *PizzaFlavorsRepository) CreatePizzaFlavor(pizzaFlavor pizza_models.PizzaFlavorCreate) (pizza_models.PizzaFlavor, error) {
	var AdditionalPrice float64 = 0
	if pizzaFlavor.AdditionalPrice != nil {
		AdditionalPrice = *pizzaFlavor.AdditionalPrice
	}
	newPizzaFlavor := pizza_models.PizzaFlavor{
		Name:            pizzaFlavor.Name,
		AdditionalPrice: AdditionalPrice,
		Description:     pizzaFlavor.Description,
		Ingredients:     pizzaFlavor.Ingredients,
	}
	if err := r.db.Create(&newPizzaFlavor).Error; err != nil {
		println("Error creating pizza flavour:", err)
		return newPizzaFlavor, err
	}
	return newPizzaFlavor, nil
}

func (r *PizzaFlavorsRepository) GetPizzaFlavorByName(name string) (pizza_models.PizzaFlavor, error) {
	var pizzaFlavor pizza_models.PizzaFlavor
	if err := r.db.Where("name = ?", name).Where("deleted_at IS NULL").First(&pizzaFlavor).Error; err != nil {
		println("Error fetching pizza flavour by name:", err)
		return pizza_models.PizzaFlavor{}, err
	}
	return pizzaFlavor, nil
}

func (r *PizzaFlavorsRepository) GetPizzaFlavorByID(id uint) (pizza_models.PizzaFlavor, error) {
	var pizzaFlavor pizza_models.PizzaFlavor
	if err := r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&pizzaFlavor).Error; err != nil {
		println("Error fetching pizza flavor by ID:", err)
		return pizza_models.PizzaFlavor{}, err
	}
	return pizzaFlavor, nil
}
