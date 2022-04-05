package user_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/model/entity"
	"backend-onboarding/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (user *userUseCase) UserList() dto.Result {
	userList, errRepo := user.userRepo.UserList()
	result := []dto.UserGetAll{}
	for _, user := range userList {
		role := dto.Role{ID: user.RoleId, Title: user.Title}
		resultData := dto.UserGetAll{
			ID:     user.ID,
			Name:   user.Name,
			Role:   role,
			Active: user.Active,
		}
		result = append(result, resultData)
	}
	if errRepo != nil && (gorm.ErrRecordNotFound == errRepo) {
		return utils.ResponseError("Data not found", errRepo, 404)
	} else if errRepo != nil {
		return utils.ResponseError("Internal server error", errRepo, 500)
	}

	return utils.ResponseSuccess("ok", nil, result, 200)
}

func (user *userUseCase) UserDetailById(id string) dto.Result {
	userUUID, err := uuid.Parse(id)

	if err != nil {
		return utils.ResponseError("Data not found", err, 404)
	}
	userDetail, errRepo := user.userRepo.UserDetailById(userUUID)

	if errRepo != nil && (gorm.ErrRecordNotFound == errRepo) {
		return utils.ResponseError("Data not found", errRepo, 404)
	} else if errRepo != nil {
		return utils.ResponseError("Internal server error", errRepo, 500)
	}

	role := dto.Role{
		ID:    userDetail.RoleId,
		Title: userDetail.Title,
	}

	userResponse := dto.UserGetById{
		ID:             userDetail.ID,
		Name:           userDetail.Name,
		Email:          userDetail.Email,
		Role:           role,
		PersonalNumber: userDetail.PersonalNumber,
		Active:         userDetail.Active,
	}
	return utils.ResponseSuccess("ok", nil, userResponse, 200)
}

func (user *userUseCase) InsertNewUser(newUser dto.User) dto.Result {
	userInsert := entity.User{
		ID:             newUser.ID,
		Name:           newUser.Name,
		Email:          newUser.Email,
		PersonalNumber: newUser.PersonalNumber,
		Password:       newUser.Password,
	}

	userData, _, err := user.userRepo.InsertNewUser(userInsert)

	if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}

	return utils.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 201)
}

func (user *userUseCase) UpdateUserData(userUpdate dto.User, id string) dto.Result {
	userInsert := entity.User{
		Name:           userUpdate.Name,
		Email:          userUpdate.Email,
		PersonalNumber: userUpdate.PersonalNumber,
		Active:         userUpdate.Active,
		Password:       userUpdate.Password,
		RoleId:         userUpdate.Role.ID,
	}

	_, errRepo := user.userRepo.UpdateUserData(userInsert, id)

	if errRepo != nil && (gorm.ErrRecordNotFound == errRepo) {
		return utils.ResponseError("Data not found", errRepo, 404)
	} else if errRepo != nil {
		return utils.ResponseError("Internal server error", errRepo, 500)
	}

	userUUID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseError("Data not found", err, 404)
	}

	userUpdate.ID = userUUID
	return utils.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (user *userUseCase) DeleteUserById(id string) dto.Result {
	userUUID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseError("Data not found", err, 404)
	}

	error := user.userRepo.DeleteUserById(userUUID)

	if error != nil && (gorm.ErrRecordNotFound == error) {
		return utils.ResponseError("Data not found", error, 404)
	} else if error != nil {
		return utils.ResponseError("Internal server error", error, 500)
	}
	return utils.ResponseSuccess("ok", nil, nil, 200)
}

func (user *userUseCase) UserLogin(userLogin dto.UserLogin) dto.Result {
	userData, err := user.userRepo.GetUserByPersonalNumber(userLogin.PersonalNumber)

	if err != nil {
		return utils.ResponseError("User not found", map[string]interface{}{"message": "Personal Number not found"}, 404)
	}

	errPwd := utils.CheckPasswordHash(userLogin.Password, userData.Password)

	if errPwd != nil {
		return utils.ResponseError("User not found", map[string]interface{}{"message": "Wrong Password"}, 404)
	}

	jwt, _ := user.jwtAuth.GenerateToken(userData.ID.String(), userData.RoleId.String())

	return utils.ResponseSuccess("ok", nil, map[string]interface{}{"token": jwt, "name": userData.Name}, 200)
}
