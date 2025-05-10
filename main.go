package main

import (
	"context"
	"log"

	_ "github.com/nguyenminhhoang/JapaneseCourses/docs"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/delivery/api/router"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/infrastructure/database"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/repository"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/usecase"
)

// @title Japanese Courses API
// @version 2.0
// @description A RESTful API service for a Japanese language learning platform
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v2
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Initialize database connection
	dbConfig := &database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "7602119",
		DBName:   "japanese_courses",
	}

	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	// Initialize router
	r := router.NewRouter()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	vocabularyRepo := repository.NewVocabularyRepository(db)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)
	vocabularyUseCase := usecase.NewVocabularyUseCase(vocabularyRepo)

	// Register routes
	r.RegisterRoutes(userUseCase, vocabularyUseCase)

	// Start server
	log.Fatal(r.Start(":8080"))
}
