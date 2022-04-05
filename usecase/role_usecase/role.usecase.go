package role_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/utils"
	"gorm.io/gorm"
)

func (role *roleUseCase) RoleList() dto.Result {
	roleList, errRepo := role.roleRepo.RoleList()
	result := []dto.Role{}
	for _, role := range roleList {
		role := dto.Role{ID: role.ID, Title: role.Title}
		result = append(result, role)
	}
	if errRepo != nil && (gorm.ErrRecordNotFound == errRepo) {
		return utils.ResponseError("Data not found", errRepo, 404)
	} else if errRepo != nil {
		return utils.ResponseError("Internal server error", errRepo, 500)
	}

	return utils.ResponseSuccess("ok", nil, result, 200)
}
