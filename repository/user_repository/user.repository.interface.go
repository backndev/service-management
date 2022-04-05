package user_repository

import (
	"backend-onboarding/model/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserList() ([]entity.UserGetAll, error)
	UserDetailById(uuid.UUID) (*entity.UserGetById, error)
	InsertNewUser(entity.User) (*entity.User, *entity.Role, error)
	UpdateUserData(entity.User, string) (*entity.User, error)
	DeleteUserById(uuid.UUID) error
	GetRoleByRoleId(uuid.UUID) (*entity.Role, error)
	GetUserByPersonalNumber(string) (*entity.User, error)
}

type userRepository struct {
	Connection *gorm.DB
}

func GetUserRepository(mysqlConn *gorm.DB) UserRepository {
	return &userRepository{
		Connection: mysqlConn,
	}
}
