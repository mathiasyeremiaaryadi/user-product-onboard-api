package usecases

import (
	"user-product-service/dto"
	"user-product-service/entities"
)

type UserUseCase interface {
	GetUsersUseCase() ([]entities.UserList, error)
	GetUserUseCase(id string) (entities.UserDetail, error)
	CreateUserUseCase(userPostRequestBody dto.UserPostRequestBody) (entities.User, error)
	UpdateUserUseCase(userPutRequestBody dto.UserPutRequestBody, id string) (int, error)
	DeleteUserUseCase(id string) error
}
