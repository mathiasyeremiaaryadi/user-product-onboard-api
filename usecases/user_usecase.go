package usecases

import (
	"strconv"
	"user-product-service/entities/dto"
	"user-product-service/entities/models"
	"user-product-service/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
	}
}

func (userUseCase UserUseCase) GetUsersUseCase() ([]models.UserList, error) {
	var userList []models.UserList
	users, err := userUseCase.userRepository.GetUsersRepository(userList)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (userUseCase UserUseCase) GetUserUseCase(id string) (models.UserDetail, error) {
	var userDetail models.UserDetail

	userID, err := strconv.Atoi(id)
	if err != nil {
		return userDetail, err
	}

	user, err := userUseCase.userRepository.GetUserRepository(userDetail, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (userUseCase UserUseCase) CreateUserUseCase(userPostRequestBody dto.UserPostRequestBody) (models.User, error) {
	userData := models.User{
		PersonalNumber: userPostRequestBody.PersonalNumber,
		Password:       userPostRequestBody.Password,
		Email:          userPostRequestBody.Email,
		Name:           userPostRequestBody.Name,
		RoleID:         5,
	}

	createdUser, err := userUseCase.userRepository.CreateUserRepository(userData)
	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}

func (userUseCase UserUseCase) UpdateUserUseCase(userPutRequestBody dto.UserPutRequestBody, id string) (int, error) {
	userData := models.User{
		PersonalNumber: userPutRequestBody.PersonalNumber,
		Password:       userPutRequestBody.Password,
		Email:          userPutRequestBody.Email,
		Name:           userPutRequestBody.Name,
		Active:         userPutRequestBody.Active,
		RoleID:         userPutRequestBody.Role["id"],
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return userID, err
	}

	updatedUserID, err := userUseCase.userRepository.UpdateUserRepository(userData, userID)
	if err != nil {
		return updatedUserID, err
	}

	return updatedUserID, nil
}

func (userUseCase UserUseCase) DeleteUserUseCase(id string) error {
	var deletedUser models.User

	userID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = userUseCase.userRepository.DeleteUserRepository(deletedUser, userID)
	if err != nil {
		return err
	}

	return nil
}
