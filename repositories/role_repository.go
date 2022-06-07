package repositories

import (
	"user-product-service/entities"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		db: db,
	}
}

func (roleRepo roleRepository) GetRolesRepository(roles []entities.Role) ([]entities.Role, error) {
	if err := roleRepo.db.Omit("Active", "CreatedAt", "UpdatedAt", "DeletedAt").Find(&roles).Error; err != nil {
		return roles, err
	}

	return roles, nil
}
