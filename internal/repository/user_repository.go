package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
)

type userRepository struct {
	db *pgx.Conn
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *pgx.Conn) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, username, password, email, created_at, updated_at 
		FROM users 
		WHERE username = $1
	`

	var user models.User
	err := r.db.QueryRow(context.Background(), query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, password, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.QueryRow(
		context.Background(),
		query,
		user.Username,
		user.Password,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}

func (r *userRepository) Update(user *models.User) error {
	query := `
		UPDATE users 
		SET username = $1, password = $2, email = $3, updated_at = $4
		WHERE id = $5
	`

	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Username,
		user.Password,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

func (r *userRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
