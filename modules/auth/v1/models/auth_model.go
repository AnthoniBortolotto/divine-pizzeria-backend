package auth_models

import (
	"time"
)

type UserRole struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Email       string     `json:"email" gorm:"unique"`
	Password    string     `json:"-"` // Password will not be exposed in JSON
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	RoleID      uint       `json:"role_id" gorm:"not null"`
	Role        UserRole   `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type RegisterRequest struct {
	Name         string `json:"name" binding:"required" validate:"required,min=3,max=50,isName"`
	Email        string `json:"email" binding:"required,email" validate:"required,email"`
	Password     string `json:"password" binding:"required" validate:"required,min=6"`
	PhoneNumber  string `json:"phone_number" binding:"required" validate:"required,min=10,max=15"`
	AddressName  string `json:"address_name" binding:"required" validate:"required,min=3,max=100"`
	Cep          string `json:"cep" binding:"required" validate:"required,isCep"`
	City         string `json:"city" binding:"required" validate:"required,min=3,max=100"`
	State        string `json:"state" binding:"required" validate:"required,min=2,max=2"`
	Complement   string `json:"complement" binding:"required" validate:"required,min=0,max=100"`
	Neighborhood string `json:"neighborhood" binding:"required" validate:"required,min=3,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
