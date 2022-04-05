package user_usecase

import (
	"backend-onboarding/repository/user_repository"
	"github.com/golang-jwt/jwt"
)

type JwtUseCase interface {
	GenerateToken(string, string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ValidateTokenAndGetUserId(string) (string, error)
	ValidateTokenAndGetRole(string) (string, error)
}

type jwtUseCase struct {
	userRepo user_repository.UserRepository
}

func GetJwtUsecase(userRepository user_repository.UserRepository) JwtUseCase {
	return &jwtUseCase{
		userRepo: userRepository,
	}
}
