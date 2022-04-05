package user_delivery

import (
	"backend-onboarding/usecase/user_usecase"
	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	UserList(*gin.Context)
	UserDetailById(*gin.Context)
	InsertNewUser(*gin.Context)
	UpdateUserData(*gin.Context)
	DeleteUserById(*gin.Context)
	UserLogin(*gin.Context)
}

type userDelivery struct {
	userUseCase user_usecase.UserUseCase
}

func GetUserDelivery(userUseCase user_usecase.UserUseCase) UserDelivery {
	return &userDelivery{
		userUseCase: userUseCase,
	}
}
