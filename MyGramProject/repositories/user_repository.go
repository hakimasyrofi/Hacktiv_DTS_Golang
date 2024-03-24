package repositories

import (
	"MyGramProject/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
	GetByUsername(username string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Update(userID uint, updateUser *entities.User) error
	Delete(userID uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
}

func (repo *userRepository) Create(user *entities.User) error {
	return repo.DB.Create(user).Error
}

func (repo *userRepository) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := repo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Update(userID uint, updateUser *entities.User) error {
	// Find the user by ID
	var user entities.User
	if err := repo.DB.First(&user, userID).Error; err != nil {
		return err // User not found
	}

	// Update user fields
	user.Username = updateUser.Username
	user.Email = updateUser.Email

	// Save the updated user
	if err := repo.DB.Save(&user).Error; err != nil {
		return err // Error updating user
	}

	return nil
}

func (repo *userRepository) Delete(userID uint) error {
	// Find the user by ID
	var user entities.User
	if err := repo.DB.First(&user, userID).Error; err != nil {
		return err // User not found
	}

	// Delete the user
	if err := repo.DB.Delete(&user).Error; err != nil {
		return err // Error deleting user
	}

	return nil
}
