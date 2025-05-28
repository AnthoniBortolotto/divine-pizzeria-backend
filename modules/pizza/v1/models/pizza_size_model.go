package pizza_models

import "time"

type PizzaSize struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	DisplayName string     `json:"display_name"`
	Price       float64    `json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Discount    float64    `json:"discount"`
}

type PizzaSizeCreate struct {
	Name        string  `json:"name" binding:"required" validate:"required,min=3,max=20"`
	DisplayName string  `json:"display_name" binding:"required" validate:"required,min=1,max=3"`
	Price       float64 `json:"price" binding:"required" validate:"required,min=0"`
	Discount    float64 `json:"discount" optional:"true" validate:"min=0"`
}
