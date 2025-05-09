package main

import (
	"context"
	"log"

	"github.com/nguyenminhhoang/JapaneseCourses/internal/delivery/api/router"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/infrastructure/database"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/repository"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/usecase"
)

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
