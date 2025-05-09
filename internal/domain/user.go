package domain

import (
	"time"

	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
)

// User represents the user entity
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int64) error
}

// UserUseCase defines the interface for user business logic
type UserUseCase interface {
	Login(username, password string) (*models.User, error)
	Register(user *models.User) error
	GetUserByID(id int64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int64) error
}
