package services

import (
	"database/sql"
	"fmt"
	
	"github.com/nithish-95/Local-Vibes/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	_, err = s.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (int, error) {
	var hashedPassword string
	var userID int
	err := s.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&userID, &hashedPassword)
	if err != nil {
		return 0, fmt.Errorf("invalid username or password: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("invalid username or password: %w", err)
	}

	return userID, nil
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := s.DB.QueryRow("SELECT id, username FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
