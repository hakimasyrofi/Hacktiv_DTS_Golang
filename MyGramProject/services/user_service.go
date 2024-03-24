package services

import (
	"MyGramProject/entities"
	"MyGramProject/helpers"
	"MyGramProject/repositories"
	"errors"
)

type UserService interface {
	RegisterUser(user *entities.User) error
	LoginUser(email, password string) (string, error)
	UpdateUser(userID uint, updateUser *entities.User) error
	DeleteUser(userID uint) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (service *userService) RegisterUser(user *entities.User) error {
	// Check if username already exists
	existingUser, err := service.repository.GetByUsername(user.Username)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return helpers.ErrUsernameExists
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Call repository to create user
	return service.repository.Create(user)
}

func (service *userService) LoginUser(email, password string) (string, error) {
	// Find user by email
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return "", err // Handle repository error
	}
	if user == nil {
		return "", errors.New("user not found") // Handle user not found error
	}

	// Check password
	if err := helpers.VerifyPassword(user.Password, password); err != nil {
		return "", err // Handle password verification error
	}

	// Generate JWT token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", err // Handle token generation error
	}

	return token, nil
}

func (service *userService) UpdateUser(userID uint, updateUser *entities.User) error {
	// Call repository to update user
	return service.repository.Update(userID, updateUser)
}

func (service *userService) DeleteUser(userID uint) error {
	// Call repository to delete user
	return service.repository.Delete(userID)
}
