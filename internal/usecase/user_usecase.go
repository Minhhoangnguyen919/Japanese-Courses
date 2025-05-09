package usecase

import (
	"errors"

	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new instance of userUseCase
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) Login(username, password string) (*models.User, error) {
	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (uc *userUseCase) Register(user *models.User) error {
	// Check if username already exists
	existingUser, err := uc.userRepo.GetByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return uc.userRepo.Create(user)
}

func (uc *userUseCase) GetUserByID(id int64) (*models.User, error) {
	// Implementation will be added when needed
	return nil, nil
}

func (uc *userUseCase) UpdateUser(user *models.User) error {
	// Implementation will be added when needed
	return nil
}

func (uc *userUseCase) DeleteUser(id int64) error {
	// Implementation will be added when needed
	return nil
}
