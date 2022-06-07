package usecases

import (
	"strconv"
	"user-product-service/dto"
	"user-product-service/entities"
	"user-product-service/repositories"
)

type userUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) UserUseCase {
	return userUseCase{
		userRepository: userRepository,
	}
}

func (userUC userUseCase) GetUsersUseCase() ([]entities.UserList, error) {
	var userList []entities.UserList
	users, err := userUC.userRepository.GetUsersRepository(userList)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (userUC userUseCase) GetUserUseCase(id string) (entities.UserDetail, error) {
	var userDetail entities.UserDetail

	userID, err := strconv.Atoi(id)
	if err != nil {
		return userDetail, err
	}

	user, err := userUC.userRepository.GetUserRepository(userDetail, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (userUC userUseCase) CreateUserUseCase(userPostRequestBody dto.UserPostRequestBody) (entities.User, error) {
	userData := entities.User{
		PersonalNumber: userPostRequestBody.PersonalNumber,
		Password:       userPostRequestBody.Password,
		Email:          userPostRequestBody.Email,
		Name:           userPostRequestBody.Name,
		RoleID:         5,
	}

	createdUser, err := userUC.userRepository.CreateUserRepository(userData)
	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}

func (userUC userUseCase) UpdateUserUseCase(userPutRequestBody dto.UserPutRequestBody, id string) (int, error) {
	userData := entities.User{
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

	updatedUserID, err := userUC.userRepository.UpdateUserRepository(userData, userID)
	if err != nil {
		return updatedUserID, err
	}

	return updatedUserID, nil
}

func (userUC userUseCase) DeleteUserUseCase(id string) error {
	var deletedUser entities.User

	userID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = userUC.userRepository.DeleteUserRepository(deletedUser, userID)
	if err != nil {
		return err
	}

	return nil
}
