package repositories

import (
	"errors"
	"user-product-service/entities/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (userRepository UserRepository) GetUsersRepository(userList []models.UserList) ([]models.UserList, error) {
	if err := userRepository.db.Preload("Role").Model(&models.User{}).Find(&userList).Error; err != nil {
		return userList, err
	}

	return userList, nil
}

func (userRepository UserRepository) GetUserRepository(userDetail models.UserDetail, id int) (models.UserDetail, error) {
	if err := userRepository.db.Preload("Role").Model(&models.User{}).First(&userDetail, id).Error; err != nil {
		return userDetail, err
	}

	return userDetail, nil
}

func (userRepository UserRepository) CreateUserRepository(user models.User) (models.User, error) {
	if err := userRepository.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userRepository UserRepository) UpdateUserRepository(user models.User, id int) (int, error) {
	if result := userRepository.db.Where("id = ?", id).Updates(&user); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (userRepository UserRepository) DeleteUserRepository(user models.User, id int) error {
	if err := userRepository.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := userRepository.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}
