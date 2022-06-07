package repositories

import "user-product-service/entities"

type UserRepository interface {
	GetUsersRepository(userList []entities.UserList) ([]entities.UserList, error)
	GetUserRepository(userDetail entities.UserDetail, id int) (entities.UserDetail, error)
	CreateUserRepository(user entities.User) (entities.User, error)
	UpdateUserRepository(user entities.User, id int) (int, error)
	DeleteUserRepository(user entities.User, id int) error
}
