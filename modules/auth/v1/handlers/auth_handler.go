package auth_handlers

import (
	"divine-pizzeria-backend/config"
	auth_models "divine-pizzeria-backend/modules/auth/v1/models"
	auth_repositories "divine-pizzeria-backend/modules/auth/v1/repositories"
	utils_validator "divine-pizzeria-backend/utils"

	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authRepo *auth_repositories.AuthRepository
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authRepo: auth_repositories.NewAuthRepository(db),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerBody auth_models.RegisterRequest
	if err := c.ShouldBindJSON(&registerBody); err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	validate.RegisterValidation("isName", utils_validator.IsName)
	validate.RegisterValidation("isCep", utils_validator.IsCep)

	err := validate.Struct(registerBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	// Check if user already exists
	existingUser, _ := h.authRepo.GetUserByEmail(registerBody.Email)
	if existingUser != nil {
		c.JSON(400, gin.H{
			"error": "User with this email already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to process password",
		})
		return
	}

	// Split name into first and last name
	nameParts := strings.SplitN(registerBody.Name, " ", 2)
	lastName := ""
	if len(nameParts) > 1 {
		lastName = nameParts[1]
	}

	// Get customer role ID
	customerRole, err := h.authRepo.GetRoleByName("customer")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get customer role",
		})
		return
	}

	// Create user with customer role
	user := &auth_models.User{
		Email:       registerBody.Email,
		Password:    string(hashedPassword),
		FirstName:   nameParts[0],
		LastName:    lastName,
		PhoneNumber: registerBody.PhoneNumber,
		Address:     fmt.Sprintf("%s, %s, %s - %s, %s", registerBody.AddressName, registerBody.Neighborhood, registerBody.City, registerBody.State, registerBody.Cep),
		RoleID:      customerRole.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdUser, err := h.authRepo.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	fmt.Printf("User created: %v\n", createdUser)

	c.JSON(201, gin.H{
		"message": "Customer registered successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginBody auth_models.LoginRequest
	if err := c.ShouldBindJSON(&loginBody); err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	// Get user by email
	user, err := h.authRepo.GetUserByEmail(loginBody.Email)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Generate JWT token with role claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	jwtSecret := config.GetEnv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"data": auth_models.LoginResponse{
			AccessToken: tokenString,
		},
	})
}
