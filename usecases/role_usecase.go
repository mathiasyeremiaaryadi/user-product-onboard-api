package usecases

import (
	"user-product-service/entities"
	"user-product-service/repositories"
)

type roleUseCase struct {
	roleRepository repositories.RoleRepository
}

func NewRoleUseCase(roleRepository repositories.RoleRepository) RoleUseCase {
	return roleUseCase{
		roleRepository: roleRepository,
	}
}

func (roleUC roleUseCase) GetRolesUseCase() ([]entities.Role, error) {
	var roles []entities.Role
	roles, err := roleUC.roleRepository.GetRolesRepository(roles)

	if err != nil {
		return roles, err
	}

	return roles, nil
}
