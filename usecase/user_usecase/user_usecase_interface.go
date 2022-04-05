package user_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/repository/user_repository"
)

type UserUseCase interface {
	UserList() dto.Result
	UserDetailById(string) dto.Result
	InsertNewUser(dto.User) dto.Result
	UpdateUserData(dto.User, string) dto.Result
	DeleteUserById(string) dto.Result
	UserLogin(dto.UserLogin) dto.Result
}

type userUseCase struct {
	userRepo user_repository.UserRepository
	jwtAuth  JwtUseCase
}

func InsertUseCase(userRepository user_repository.UserRepository, jwtAuthentication JwtUseCase) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
		jwtAuth:  jwtAuthentication,
	}
}
