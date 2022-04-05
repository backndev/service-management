package role_repository

import (
	"backend-onboarding/model/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	RoleList() ([]entity.Role, error)
}

type roleRepository struct {
	DB *gorm.DB
}

func GetRoleRepository(DB *gorm.DB) RoleRepository {
	return &roleRepository{
		DB: DB,
	}
}
