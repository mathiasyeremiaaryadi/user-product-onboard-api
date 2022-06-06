package repositories

import (
	"user-product-service/entities/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return RoleRepository{
		db: db,
	}
}

func (roleRepository RoleRepository) GetRolesRepository(roles []models.Role) ([]models.Role, error) {
	if err := roleRepository.db.Omit("Active", "CreatedAt", "UpdatedAt", "DeletedAt").Find(&roles).Error; err != nil {
		return roles, err
	}

	return roles, nil
}
