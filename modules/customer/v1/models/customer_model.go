package customer_models

import (
	"fmt"
	"strings"
	"time"
)

type Customer struct {
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type CustomerCreate struct {
	Name         string `json:"name" binding:"required" validate:"required,min=3,max=50,isName"`
	Email        string `json:"email" binding:"required,email" validate:"required,email"`
	PhoneNumber  string `json:"phone_number" binding:"required" validate:"required,min=10,max=15"`
	AddressName  string `json:"address_name" binding:"required" validate:"required,min=3,max=100"`
	Cep          string `json:"cep" binding:"required" validate:"required,isCep"`
	City         string `json:"city" binding:"required" validate:"required,min=3,max=100"`
	State        string `json:"state" binding:"required" validate:"required,min=2,max=2"`
	Complement   string `json:"complement" binding:"required" validate:"required,min=0,max=100"`
	Neighborhood string `json:"neighborhood" binding:"required" validate:"required,min=3,max=100"`
}

func (c *CustomerCreate) ToCustomer() Customer {
	nameParts := strings.SplitN(c.Name, " ", 2)

	return Customer{
		FirstName:   nameParts[0],
		LastName:    nameParts[1],
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
		Address:     fmt.Sprintf("%s, %s, %s - %s, %s", c.AddressName, c.Neighborhood, c.City, c.State, c.Cep),
	}
}
