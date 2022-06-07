package repositories

import "user-product-service/entities"

type RoleRepository interface {
	GetRolesRepository(roles []entities.Role) ([]entities.Role, error)
}
