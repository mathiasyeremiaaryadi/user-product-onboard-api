package repositories

import (
	"errors"
	"user-product-service/entities"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db: db,
	}
}

func (userRepo userRepository) GetUsersRepository(userList []entities.UserList) ([]entities.UserList, error) {
	if err := userRepo.db.Preload("Role").Model(&entities.User{}).Find(&userList).Error; err != nil {
		return userList, err
	}

	return userList, nil
}

func (userRepo userRepository) GetUserRepository(userDetail entities.UserDetail, id int) (entities.UserDetail, error) {
	if err := userRepo.db.Preload("Role").Model(&entities.User{}).First(&userDetail, id).Error; err != nil {
		return userDetail, err
	}

	return userDetail, nil
}

func (userRepo userRepository) CreateUserRepository(user entities.User) (entities.User, error) {
	if err := userRepo.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userRepo userRepository) UpdateUserRepository(user entities.User, id int) (int, error) {
	if result := userRepo.db.Where("id = ?", id).Updates(&user); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (userRepo userRepository) DeleteUserRepository(user entities.User, id int) error {
	if err := userRepo.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := userRepo.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}
