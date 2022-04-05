package user_repository

import (
	"backend-onboarding/model/entity"
	"backend-onboarding/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (userRepo *userRepository) UserList() ([]entity.UserGetAll, error) {
	var userList []entity.UserGetAll
	err := userRepo.Connection.Model(&entity.User{}).Select("users.name, users.active, users.id, roles.title, users.roleId").Joins("left join roles on roles.id = users.roleId").Scan(&userList).Error
	//check error execute query
	if err != nil {
		return nil, err
	}
	if len(userList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return userList, nil
}

func (userRepo *userRepository) UserDetailById(id uuid.UUID) (*entity.UserGetById, error) {
	userdetail := entity.UserGetById{}
	err := userRepo.Connection.Model(&entity.User{}).Where("users.id = ?", id).Select("users.name, users.active, users.email, users.personalNumber, users.id, roles.title, users.roleId").Joins("left join roles on roles.id = users.roleId").First(&userdetail).Error
	//check user exist or not
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &userdetail, nil
}

func (userRepo *userRepository) InsertNewUser(user entity.User) (*entity.User, *entity.Role, error) {
	role := entity.Role{}

	result := userRepo.Connection.Where("email = ? OR PersonalNumber = ?", user.Email, user.PersonalNumber).Find(&user)
	if result.RowsAffected > 0 {
		return nil, nil, gorm.ErrRegistered
	}

	user.ID = uuid.New()
	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash
	if err := userRepo.Connection.Where("title = ?", "viewer").Find(&role).Error; err != nil {
		return nil, nil, err
	}

	user.RoleId = role.ID
	if err := userRepo.Connection.Create(&user).Error; err != nil {
		return nil, nil, err
	}

	return &user, &role, nil
}

func (userRepo *userRepository) UpdateUserData(user entity.User, id string) (*entity.User, error) {
	hash, _ := utils.HashPassword(user.Password)

	result := userRepo.Connection.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"name":           user.Name,
		"password":       hash,
		"roleId":         user.RoleId,
		"active":         user.Active,
		"email":          user.Email,
		"personalNumber": user.PersonalNumber,
	})

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (userRepo *userRepository) DeleteUserById(id uuid.UUID) error {
	user := entity.User{}
	result := userRepo.Connection.Where("id = ?", id).Find(&user)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if err := userRepo.Connection.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (userRepo *userRepository) GetRoleByRoleId(id uuid.UUID) (*entity.Role, error) {
	role := entity.Role{}
	result := userRepo.Connection.Where("id = ?", id).Find(&role)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &role, nil
}

func (userRepo *userRepository) GetUserByPersonalNumber(pn string) (*entity.User, error) {
	user := entity.User{}
	result := userRepo.Connection.Where("PersonalNumber = ?", pn).Find(&user)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}
