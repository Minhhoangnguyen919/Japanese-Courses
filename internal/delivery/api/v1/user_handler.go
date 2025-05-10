package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// Register registers the user routes
func (h *UserHandler) Register(g *echo.Group) {
	users := g.Group("/users")
	{
		users.POST("/login", h.Login)
		users.POST("/register", h.RegisterUser)
	}
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
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

	return c.JSON(http.StatusOK, user)
}

type registerRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
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

	return c.JSON(http.StatusCreated, user)
}
