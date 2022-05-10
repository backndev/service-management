package test_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/model/entity"
	"backend-onboarding/repository/test_repository"
	"backend-onboarding/usecase/user_usecase"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var userRepository = test_repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUc = user_usecase.UserUsecaseTest{UserRepo: &userRepository}

func TestGetAllUsersSuccess(t *testing.T) {
	expected := []entity.User{}
	userRepository.Mock.On("GetAllUsers").Return(expected, nil)
	result, err := userUc.UserRepo.GetAllUsers()

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestGetUserByIdNotFound(t *testing.T) {
	userRepository.Mock.On("GetUserById", "1").Return(nil, errors.New("Data not found"))

	result, err := userUc.UserRepo.GetUserById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestGetUserByIdSuccess(t *testing.T) {
	expected := &entity.User{}

	userRepository.Mock.On("GetUserById", "2").Return(expected, nil)

	result, err := userUc.UserRepo.GetUserById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestUpdateUserNotFound(t *testing.T) {
	userRepository.Mock.On("GetUserById", "1").Return(nil, errors.New("Data not found"))

	result, err := userUc.UserRepo.GetUserById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestUpdateUserSuccess(t *testing.T) {
	expected := &entity.User{}

	userRepository.Mock.On("UpdateUserData", "2").Return(expected, nil)

	result, err := userUc.UserRepo.UpdateUserData("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewUserSuccess(t *testing.T) {
	request := dto.User{
		Name:           "New User",
		PersonalNumber: "123456789",
	}

	expected := &entity.User{}

	userRepository.Mock.On("CreateNewUser").Return(expected, nil)

	result, err := userUc.UserRepo.CreateNewUser(request)

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewUserFailed(t *testing.T) {
	request := dto.User{
		Name:           "New User",
		PersonalNumber: "123456789",
	}

	expected := &entity.User{}

	userRepository.Mock.On("CreateNewUser").Return(nil, errors.New("Failed create User"))

	result, err := userUc.UserRepo.CreateNewUser(request)

	assert.NotEqual(t, expected, result)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestDeleteUserFailed(t *testing.T) {
	userRepository.Mock.On("DeleteUserById", "1").Return(nil, errors.New("Data not found"))

	result, err := userUc.UserRepo.DeleteUserById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestDeleteUserSuccess(t *testing.T) {
	expected := &entity.User{}

	userRepository.Mock.On("DeleteUserById", "2").Return(expected, nil)

	result, err := userUc.UserRepo.DeleteUserById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}
