package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/nguyenminhhoang/JapaneseCourses/internal/delivery/api/v1"
	v2 "github.com/nguyenminhhoang/JapaneseCourses/internal/delivery/api/v2"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/infrastructure/auth"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Router represents the API router
type Router struct {
	echo *echo.Echo
}

// NewRouter creates a new router instance
func NewRouter() *Router {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return &Router{
		echo: e,
	}
}

// RegisterRoutes registers all API routes
func (r *Router) RegisterRoutes(userUseCase domain.UserUseCase, vocabularyUseCase domain.VocabularyUseCase) {
	// Swagger documentation
	r.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// API v1
	v1Group := r.echo.Group("/api/v1")
	{
		userHandlerV1 := v1.NewUserHandler(userUseCase)
		userHandlerV1.Register(v1Group)
	}

	// API v2
	v2Group := r.echo.Group("/api/v2")
	{
		// Initialize JWT service
		jwtService := auth.NewJWTService("your-secret-key") // In production, use environment variable

		// Initialize handlers
		userHandlerV2 := v2.NewUserHandler(userUseCase)
		vocabularyHandlerV2 := v2.NewVocabularyHandler(vocabularyUseCase)
		vocabularyProgressHandlerV2 := v2.NewVocabularyProgressHandler(vocabularyUseCase, jwtService)

		// Register routes
		userHandlerV2.Register(v2Group)
		vocabularyHandlerV2.Register(v2Group)
		vocabularyProgressHandlerV2.Register(v2Group)
	}
}

// Start starts the server
func (r *Router) Start(port string) error {
	return r.echo.Start(port)
}
