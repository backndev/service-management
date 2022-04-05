package role_repository

import (
	"backend-onboarding/model/entity"
	"gorm.io/gorm"
)

func (roleRepo *roleRepository) RoleList() ([]entity.Role, error) {
	var roleList []entity.Role
	err := roleRepo.DB.Model(&entity.Role{}).Select("id, title").Scan(&roleList).Error
	if err != nil {
		return nil, err
	}
	if len(roleList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return roleList, nil
}
