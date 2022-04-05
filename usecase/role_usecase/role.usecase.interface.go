package role_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/repository/role_repository"
)

type RoleUseCase interface {
	RoleList() dto.Result
}

type roleUseCase struct {
	roleRepo role_repository.RoleRepository
}

func GetRoleUseCase(roleRepository role_repository.RoleRepository) RoleUseCase {
	return &roleUseCase{
		roleRepo: roleRepository,
	}
}
