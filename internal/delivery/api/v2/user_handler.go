package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/infrastructure/auth"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
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

// Register registers the user routes
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

// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body loginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Server error"
// @Router /users/auth/login [post]
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

// @Summary Register new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body registerRequest true "User registration data"
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Server error"
// @Router /users/auth/register [post]
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
