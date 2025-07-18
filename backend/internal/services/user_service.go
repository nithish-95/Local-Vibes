package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/nithish-95/Local-Vibes/backend/internal/models"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{Username: username, Password: string(hashedPassword)}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to register user: %w", result.Error)
	}
	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (uint, error) {
	var user models.User
	result := s.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return 0, fmt.Errorf("invalid username or password: %w", result.Error)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("invalid username or password: %w", err)
	}

	return user.ID, nil
}

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := s.DB.First(&user, userID)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found: %w", result.Error)
	}
	return &user, nil
}
