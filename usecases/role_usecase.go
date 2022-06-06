package usecases

import (
	"user-product-service/entities/models"
	"user-product-service/repositories"
)

type RoleUseCase struct {
	roleRepository repositories.RoleRepository
}

func NewRoleUseCase(roleRepository repositories.RoleRepository) RoleUseCase {
	return RoleUseCase{
		roleRepository: roleRepository,
	}
}

func (roleUseCase RoleUseCase) GetRolesUseCase() ([]models.Role, error) {
	var roles []models.Role
	roles, err := roleUseCase.roleRepository.GetRolesRepository(roles)

	if err != nil {
		return roles, err
	}

	return roles, nil
}
