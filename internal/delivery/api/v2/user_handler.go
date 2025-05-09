package v2

import (
	"net/http"

	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/infrastructure/auth"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
	jwtService  *auth.JWTService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
		jwtService:  auth.NewJWTService("your-secret-key"), // In production, use environment variable
	}
}

// RegisterRoutes registers the user routes
func (h *UserHandler) Register(g *echo.Group) {
	users := g.Group("/users")
	{
		users.POST("/auth/login", h.Login)           // Changed from /login to /auth/login
		users.POST("/auth/register", h.RegisterUser) // Changed from /register to /auth/register
	}
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	DeviceID string `json:"device_id"` // New field in v2
}

func (h *UserHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	user, err := h.userUseCase.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	// Generate JWT token
	token, err := h.jwtService.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate token",
		})
	}

	// Enhanced response in v2
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":       user,
		"device_id":  req.DeviceID,
		"version":    "v2",
		"token":      token,
		"token_type": "Bearer",
	})
}

type registerRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name"`
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req registerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := h.userUseCase.Register(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Enhanced response in v2
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user":      user,
		"full_name": req.FullName,
		"version":   "v2",
	})
}
