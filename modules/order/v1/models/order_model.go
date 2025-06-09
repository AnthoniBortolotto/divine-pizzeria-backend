package order_models

import (
	auth_models "divine-pizzeria-backend/modules/auth/v1/models"
	customer_models "divine-pizzeria-backend/modules/customer/v1/models"
	pizza_models "divine-pizzeria-backend/modules/pizza/v1/models"
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID           int                      `json:"id"`
	UserID       uint                     `json:"user_id"`
	User         auth_models.User         `json:"-" gorm:"foreignKey:UserID"`
	Customer     customer_models.Customer `json:"customer" gorm:"-"` // Virtual field for JSON response
	DeliveryDate *time.Time               `json:"delivery_date"`
	Status       OrderStatus              `json:"status"`
	TotalPrice   float64                  `json:"total_price"`
	Items        []OrderItem              `json:"items"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	DeletedAt    *time.Time               `json:"deleted_at,omitempty"`
}

type OrderItem struct {
	ID          uint                       `json:"id"`
	OrderID     uint                       `json:"order_id"`
	PizzaSizeID uint                       `json:"pizza_size_id" gorm:"column:pizza_size_id"`
	PizzaSize   pizza_models.PizzaSize     `json:"pizza_size" gorm:"foreignKey:PizzaSizeID"`
	Flavors     []pizza_models.PizzaFlavor `json:"flavors" gorm:"many2many:order_item_flavors;foreignKey:ID;joinForeignKey:OrderItemID;References:ID;joinReferences:PizzaFlavorID"`
	Quantity    uint                       `json:"quantity"`
	UnitPrice   float64                    `json:"unit_price"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	DeletedAt   *time.Time                 `json:"deleted_at,omitempty"`
}

type CreateOrderRequest struct {
	DeliveryDate *time.Time        `json:"delivery_date" validate:"omitempty,isdateafternow"`
	Items        []OrderItemCreate `json:"items" binding:"required" validate:"required,min=1,dive"`
}

type OrderItemCreate struct {
	PizzaSizeID uint   `json:"pizza_size_id" binding:"required" validate:"required,min=1"`
	FlavorIDs   []uint `json:"flavor_ids" binding:"required" validate:"required,min=1,dive,required"`
	Quantity    uint   `json:"quantity" binding:"required" validate:"required,min=1"`
}

type OrderUpdate struct {
	Status OrderStatus `json:"status" binding:"required" validate:"required,oneof=pending preparing ready delivered cancelled"`
}

// CalculateTotalPrice calculates the total price of the order including all items
func (o *Order) CalculateTotalPrice() {
	var total float64
	for _, item := range o.Items {
		total += item.UnitPrice * float64(item.Quantity)
	}
	o.TotalPrice = total
}
