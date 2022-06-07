package usecases

import "user-product-service/entities"

type RoleUseCase interface {
	GetRolesUseCase() ([]entities.Role, error)
}
