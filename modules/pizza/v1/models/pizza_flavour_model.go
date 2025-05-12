package pizza_models

import "time"

type PizzaFlavor struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	AdditionalPrice float64    `json:"additional_price"`
	Description     string     `json:"description"`
	Ingredients     string     `json:"ingredients"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type PizzaFlavorCreate struct {
	Name            string   `json:"name" binding:"required" validate:"required,min=3,max=20"`
	Description     string   `json:"description" binding:"required" validate:"required,min=3,max=100"`
	AdditionalPrice *float64 `json:"additional_price,omitempty" validate:"min=0"`
	Ingredients     string   `json:"ingredients" binding:"required" validate:"required,min=3,max=100"`
}
